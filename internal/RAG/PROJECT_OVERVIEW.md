# Tổng Quan Project ERC-RAG

## 1. Bài toán

Project **ERC-RAG** (Emotion Recognition in Conversation – Retrieval Augmented Generation) giải quyết bài toán **Nhận diện cảm xúc trong hội thoại**: 
cho một đoạn chat nhiều lượt giữa các nhân vật, dự đoán nhãn cảm xúc của utterance cuối cùng trong 7 lớp:

```
neutral, surprise, fear, sadness, joy, disgust, anger
```

Dữ liệu sử dụng: **MELD** (Multimodal EmotionLines Dataset, bản text-only từ declare-lab/MELD trên GitHub).

## 2. Ý tưởng tiếp cận

Thay vì dùng một bộ phân loại thuần tuý, project kết hợp 3 thành phần:

1. **RAG (Retrieval Augmented Generation)** – truy xuất các ví dụ hội thoại tương tự từ tập train để làm few-shot cho LLM.
2. **Emotion Shift Score** – cải tiến so với RAG chuẩn: xếp hạng lại các ví dụ dựa trên **độ tương đồng quỹ đạo cảm xúc** của nhân vật (không chỉ tương đồng văn bản).
3. **LLM Generator** – có thể chạy 4 nhánh: API (Groq / Gemini), Local CPU (GGUF Llama 3.2), hoặc **Fine-tune LoRA** cục bộ trên GPU (Qwen2.5-3B-Instruct, LLama3-8B).

## 3. Kiến trúc mã nguồn

```
config.py                    # Config tập trung (paths, model names, hyperparameters)
src/
 ├─ data_processor.py        # Làm sạch CSV MELD -> JSON theo dialogue
 ├─ vector_store.py          # ChromaDB + SentenceTransformer (all-MiniLM-L6-v2)
 ├─ retriever.py             # EmotionAwareRetriever: Semantic + Emotion Shift
 ├─ prompt_builder.py        # Few-shot prompt theo ChatML
 ├─ generator_gemini.py      # Inference qua Gemini API
 ├─ generator_groq.py        # Inference qua Groq API (Llama-3.3-70B)
 ├─ generator_local_cpu.py   # Llama 3.2 3B GGUF qua llama-cpp-python
 ├─ generator_lora.py        # LoRA/PEFT (training + inference)
 └─ evaluate.py              # Accuracy, F1, Confusion Matrix, Learning Curve
scripts/
 ├─ 01_download_data.py      # Tải MELD (train/dev/test)
 ├─ 02_preprocess.py         # Làm sạch và nhóm theo dialogue
 ├─ 03_build_vectordb.py     # Build ChromaDB từ train split
 ├─ 04_train_lora.py         # Fine-tune LoRA với SFTTrainer
 ├─ 05_inference_{groq,gemini,local,lora}.py
 ├─ 06_evaluate.py           # Tổng hợp metric và vẽ plot
 ├─ 07_infer_eval_colab.py   # Script đánh giá trên Google Colab
 └─ pack_for_colab.py        # Đóng gói project để chạy Colab
```

## 4. Các bước giải quyết bài toán (pipeline)

### Bước 1 – Chuẩn bị dữ liệu

- **`01_download_data.py`**: tải 3 file `train_sent_emo.csv`, `dev_sent_emo.csv`, `test_sent_emo.csv` từ repo declare-lab/MELD về `data/raw/`.
- **`02_preprocess.py`** (dùng `DataProcessor` ở [src/data_processor.py](src/data_processor.py)):
  - Chuẩn hoá tên cột (`Utterance`, `Speaker`, `Emotion`, `Dialogue_ID`, `Utterance_ID`).
  - Làm sạch text (ký tự lạ, khoảng trắng).
  - **Gom theo `Dialogue_ID`** và sắp xếp theo `Utterance_ID` để giữ thứ tự thời gian.
  - Với mỗi utterance lưu thêm `emotion_history` = danh sách cảm xúc của các câu trước đó trong cùng hội thoại → đây là cơ sở cho Emotion Shift Score ở bước retrieval.
  - Xuất `train.json / dev.json / test.json` vào `data/processed/`.

### Bước 2 – Xây Vector Store

- **`03_build_vectordb.py`** dùng [src/vector_store.py](src/vector_store.py):
  - Với mỗi utterance trong tập train, tạo **document** là 3 câu gần nhất nối bằng ` [SEP] ` theo dạng `Speaker: text`.
  - Embedding bằng `all-MiniLM-L6-v2` (SentenceTransformer).
  - Lưu vào **ChromaDB** persistent, space = `cosine`.
  - Metadata kèm theo: `target_emotion`, `emotion_history` (JSON-string), `document_text`.

### Bước 3 – Retrieval cải tiến (Emotion-Aware)

Logic nằm ở [src/retriever.py:51](src/retriever.py:51):

1. Lấy top-`fetch_k = max(top_k*4, 20)` ứng viên gần nhất theo semantic (ChromaDB).
2. Map mỗi cảm xúc sang **valence 1D** (`joy=1.0`, `anger=-1.0`, …).
3. Tính **khoảng cách quỹ đạo cảm xúc** giữa query và ứng viên bằng MAE trên phần đuôi (các cảm xúc gần hiện tại nhất), chuẩn hoá về `[0,1]`.
4. Điểm tổng hợp: `final_distance = α * sem_dist + (1-α) * emo_dist` (mặc định `α = 0.7`).
5. Sắp xếp tăng dần theo `final_distance`, trả về `TOP_K = 2` ví dụ tốt nhất.

### Bước 4 – Dựng Prompt Few-shot

- [src/prompt_builder.py](src/prompt_builder.py) sinh prompt **ChatML** gồm:
  - System: yêu cầu chỉ trả về 1 từ cảm xúc trong 7 nhãn, mặc định `neutral` nếu không chắc chắn.
  - User: các `Example i` (ngữ cảnh + nhãn) + đoạn hội thoại query.
  - Assistant: để mô hình sinh nhãn.

### Bước 5 – Generator (chọn 1 trong 4 luồng)

| Luồng | Script | Model | Yêu cầu |
|---|---|---|---|
| A1 | `05_inference_groq.py` | `llama-3.3-70b-versatile` qua Groq API | `GROQ_API_KEY`, nhanh ~5 phút cho test set |
| A2 | `05_inference_gemini.py` | `gemini-2.5-flash-lite` | `GEMINI_API_KEY` |
| B  | `05_inference_local.py` | Llama-3.2-3B-Instruct GGUF Q4_K_M | CPU, llama-cpp-python |
| C  | `05_inference_lora.py`  | Qwen2.5-3B-Instruct / LLama3-8B sau LoRA | GPU NVIDIA |

### Bước 6 – Fine-tune LoRA (luồng C)

[scripts/04_train_lora.py](scripts/04_train_lora.py):

1. **Balanced Sampling**: gom toàn bộ utterance trong train, phân nhóm theo 7 nhãn, lấy mỗi nhãn `TRAIN_MAX_SAMPLES / 7` mẫu (`5000 / 7`) để chống mất cân bằng lớp.
2. Với từng mẫu, gọi `EmotionAwareRetriever` để lấy ví dụ → build prompt → nối nhãn đích vào cuối → tạo dataset SFT dạng `{"text": ...}`.
3. Huấn luyện với **TRL SFTTrainer** + **PEFT LoRA** (r=16, α=32, dropout=0.05), `paged_adamw_8bit`, `gradient_checkpointing`, batch=2, grad accum=8, lr=2e-4, `MAX_EPOCHS=5`.
4. **Early Stopping** `patience=3` trên `eval_loss` (eval mỗi 50 steps) → `load_best_model_at_end=True`.
5. Lưu log theo step ra `outputs/plots/training_logs.json` để vẽ Learning Curve ở bước sau, lưu model tại `outputs/models/best_lora_model`.
6. (Tuỳ chọn) report metrics sang **Weights & Biases** nếu có `WANDB_API_KEY`.

### Bước 7 – Đánh giá & So sánh

[scripts/06_evaluate.py](scripts/06_evaluate.py) + [src/evaluate.py](src/evaluate.py):

- Đọc các CSV `outputs/results/{gemini,lora,groq,local}_predictions.csv`.
- Dự đoán không hợp lệ (ảo giác ngoài 7 nhãn) → ép về `neutral`.
- Tính **Accuracy, Macro-F1, Weighted-F1, Micro-F1**.
- Vẽ **Confusion Matrix** cho từng mô hình.
- Vẽ **Learning Curve** (train vs val loss) + đánh dấu điểm Early Stop / best epoch.

### Bước 8 – Chạy đánh giá trên Google Colab

- `pack_for_colab.py` đóng gói project (kèm cờ `--include-model`) thành `project.zip`.
- Upload lên Colab (T4/L4/A100 GPU), giải nén, `pip install -r requirements.txt`.
- `07_infer_eval_colab.py` chạy suy luận trên toàn bộ 2610 mẫu test cho cả 2 model đã fine-tune (Qwen2.5-3B, Llama3-8B) ở 2 chế độ: **RAG few-shot** và **Zero-shot** → so sánh đóng góp của RAG.

## 5. Các tham số chính (`config.py`)

- `TOP_K = 2`, `EMOTION_ALPHA = 0.7` (trọng số semantic vs emotion shift).
- `MAX_NEW_TOKENS = 64`, `TEMPERATURE = 0.1` (ưu tiên output ổn định).
- `TRAIN_MAX_SAMPLES = 5000`, `BALANCE_DATA = True`.
- LoRA: `r=16, α=32, dropout=0.05`, `BATCH_SIZE=2`, `GRAD_ACC=8`, `LR=2e-4`, `EPOCHS=5`, `EARLY_STOPPING_PATIENCE=3`.
- Embedding: `all-MiniLM-L6-v2`; Generator mặc định: `Qwen/Qwen2.5-3B-Instruct`.

## 6. Điểm nhấn kỹ thuật

1. **Emotion-Aware Retrieval** ([src/retriever.py:22](src/retriever.py:22)): kết hợp semantic + quỹ đạo valence của nhân vật → chọn ví dụ có *hành vi chuyển đổi tâm lý* tương tự, không chỉ nội dung tương tự.
2. **Balanced Sampling + LoRA PEFT + Early Stopping** ([scripts/04_train_lora.py:57](scripts/04_train_lora.py:57)): chống mất cân bằng lớp (MELD rất lệch về `neutral`) và chống overfitting.
3. **Đa luồng inference** (API / CPU / GPU fine-tune) cho phép so sánh chi phí – hiệu năng và benchmark chéo (zero-shot vs RAG vs LoRA) trên cùng tập test.
