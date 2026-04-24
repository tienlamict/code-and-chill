
## A. Câu hỏi về bài toán & dữ liệu

### A1. Bài toán ERC khác gì so với phân tích cảm xúc (Sentiment Analysis) thông thường?
- Sentiment Analysis thường gán 1 nhãn (positive/negative/neutral) cho **một câu độc lập**.
- ERC (Emotion Recognition in Conversation) phải gán nhãn cho utterance trong **bối cảnh cả đoạn hội thoại nhiều lượt**: 
cần theo dõi ai nói, lịch sử cảm xúc, và sự dịch chuyển cảm xúc (emotion shift).
- Nhãn ERC cũng mịn hơn: 7 lớp (`neutral, surprise, fear, sadness, joy, disgust, anger`) thay vì 3.

### A2. Tại sao chọn MELD?
- MELD là benchmark chuẩn cho ERC, lấy từ series *Friends*, có gán nhãn speaker + emotion + dialogue_id.
- Đủ đa dạng (≈13K utterance) để train, dev, test tách sẵn.
- Bản text-only (từ declare-lab/MELD) nhẹ, không cần xử lý audio/video, phù hợp phạm vi project.
- Có **mất cân bằng lớp nặng** (~50% là `neutral`), là thử thách thực tế để chứng minh kỹ thuật balanced sampling.

### A3. Dữ liệu được tiền xử lý thế nào?
Trong [src/data_processor.py](src/data_processor.py):
1. Đọc CSV, chuẩn hoá tên cột (`Utterance, Speaker, Emotion, Dialogue_ID, Utterance_ID`).
2. Làm sạch text: thay ký tự lạ (`\x92 → '`), gộp khoảng trắng.
3. **Nhóm theo `Dialogue_ID`**, sắp theo `Utterance_ID` để giữ đúng trật tự thời gian.
4. Với từng utterance, lưu kèm `emotion_history` là danh sách cảm xúc của các lượt nói trước → đây là dữ liệu nền cho Emotion Shift Score.
5. Xuất JSON theo cấu trúc cây dialogue.

### A4. Có bao nhiêu mẫu và phân bố lớp ra sao?
- Train ≈ 9989, Dev ≈ 1109, Test ≈ 2610 utterances.
- Phân bố lệch mạnh: `neutral` chiếm đa số, `fear, disgust` rất ít (~2–3%). Đây là lý do phải dùng **Macro-F1** và balanced sampling khi train.

---

## B. Câu hỏi về kiến trúc & lý do chọn RAG

### B1. Tại sao dùng RAG mà không phân loại trực tiếp bằng BERT/RoBERTa?
- BERT classifier cần fine-tune riêng, khó tận dụng tri thức tổng quát của LLM, và không dễ cho few-shot.
- **RAG + LLM** có lợi thế:
  - Không cần classifier head, chỉ sinh text.
  - Cho phép few-shot động: mỗi query lấy ví dụ khác nhau → thích ứng theo ngữ cảnh.
  - Có thể chạy **zero-shot** với LLM mạnh (Groq 70B, Gemini) mà không cần train.
  - Dễ cắm-vào-chơi với nhiều backend LLM khác nhau (Groq, Gemini, local, LoRA).
- Quan trọng nhất: RAG cho phép chúng tôi **inject thêm tín hiệu ngoài văn bản** (quỹ đạo cảm xúc) vào bước retrieval – đây là điểm cải tiến.

### B2. "RAG cải tiến" ở đây cụ thể là gì?
RAG chuẩn chỉ xếp hạng bằng độ tương đồng semantic (cosine giữa embedding). Chúng tôi **bổ sung Emotion Shift Score**:
- Map mỗi cảm xúc sang valence 1-D (`joy=1.0 … anger=-1.0`).
- Với mỗi ứng viên trong top-N, tính MAE giữa quỹ đạo valence của query và quỹ đạo valence trong lịch sử ứng viên → `emo_dist ∈ [0,1]`.
- `final_distance = α * sem_dist + (1-α) * emo_dist` với `α = 0.7`.
- Sắp xếp lại và chọn top-K = 2.

Ý nghĩa: ví dụ được chọn không chỉ **giống nội dung** mà còn có **pattern chuyển đổi cảm xúc tương tự** nhân vật hiện tại → few-shot có giá trị hơn cho dự đoán emotion shift.

### B3. Có thể thay MAE bằng DTW được không? Tại sao không dùng DTW?
- DTW (Dynamic Time Warping) giải quyết khi 2 chuỗi khác độ dài hoặc lệch pha.
- Trong project này chuỗi rất ngắn (vài–vài chục cảm xúc) và chúng tôi đã cắt phần đuôi cùng độ dài (`min_len`), MAE đã đủ, rẻ hơn và dễ giải thích.
- DTW là hướng mở rộng hợp lý nếu muốn so sánh toàn bộ quỹ đạo dài ngắn khác nhau – README có nhắc tới như tương lai phát triển.

### B4. Valence mapping hard-code có thô không? Có thể cải tiến ra sao?
- Có – mapping hiện tại là tay (`joy=1.0, anger=-1.0`).
- Cải tiến: dùng mô hình **VAD** (Valence–Arousal–Dominance) để có vector 3-D, hoặc học embedding của cảm xúc từ dữ liệu để khoảng cách phản ánh thống kê thực.
- Hiện tại mapping 1-D được chọn vì đơn giản, giải thích được, và đủ tốt để phân biệt 7 nhãn trên MELD.

### B5. Tại sao `α = 0.7`?
- Đây là hyperparameter cân bằng giữa "ngữ nghĩa văn bản" và "tương đồng tâm lý".
- 0.7 nghiêng về semantic vì text vẫn là tín hiệu mạnh nhất; emotion shift bổ trợ. Chúng tôi thử nghiệm và chọn giá trị tốt nhất trên dev set.
- Nếu `α = 1.0` → RAG chuẩn; `α = 0` → chỉ dựa emotion shift (thiếu nội dung).

---

## C. Câu hỏi về Vector Store & Embedding

### C1. Vì sao chọn `all-MiniLM-L6-v2` làm embedding model?
- Model 22M tham số, 384-d vector, **chạy được trên CPU**, tốc độ encode rất nhanh.
- Chất lượng semantic trên tiếng Anh tốt, phù hợp MELD (toàn tiếng Anh).
- Nhẹ → phù hợp pipeline chạy trên máy sinh viên, không ép buộc GPU cho bước build index.

### C2. Tại sao document là 3 câu gần nhất, không phải toàn bộ dialogue?
- Ngữ cảnh cảm xúc phụ thuộc mạnh vào **các lượt nói gần nhất** – câu xa thường ít liên quan.
- Cửa sổ 3 giữ được short-term context mà không bị dilute khi encoding.
- Nếu lấy cả dialogue, mean-pooling làm mất tín hiệu của câu hiện tại (câu cần dự đoán).

### C3. Sao lại nối bằng ` [SEP] `?
- Đánh dấu ranh giới câu cho mô hình nhúng không bị lẫn lộn speaker, đồng thời dễ tách ngược khi build prompt (`replace(" [SEP] ", "\n")`).

### C4. ChromaDB dùng metric gì? Distance có ý nghĩa thế nào?
- `metadata={"hnsw:space": "cosine"}` – khoảng cách cosine = `1 - cosine_similarity`.
- Distance càng nhỏ → 2 vector càng giống. Trong retriever dùng trực tiếp làm `sem_dist ∈ [0, ~1]`.

### C5. `fetch_k = max(top_k*4, 20)` để làm gì?
- Lấy **nhiều** ứng viên theo semantic trước, sau đó **re-rank** bằng score tổng hợp để chọn top-K cuối.
- Nếu chỉ lấy đúng top_k theo semantic thì không còn cơ hội để emotion shift thay đổi thứ hạng → mất ý nghĩa của RAG cải tiến.

---

## D. Câu hỏi về Generator & Fine-tune LoRA

### D1. Tại sao lại có 4 luồng generator?
Để so sánh trade-off:
| Luồng | Ưu điểm | Nhược điểm |
|---|---|---|
| Groq 70B | Mạnh, nhanh (~5 phút cả test set) | Phụ thuộc API, có rate limit |
| Gemini Flash-Lite | Ổn định, miễn phí quota | Chậm hơn, cần API key Google |
| Local CPU (Llama 3.2 3B GGUF) | Offline, không tốn tiền | Chậm, model nhỏ → độ chính xác thấp |
| LoRA (Qwen 3B / Llama3-8B) | Tối ưu riêng cho MELD, so được với few-shot | Cần GPU, phải train |

Mục tiêu sư phạm: chứng minh có thể triển khai đa cấu hình, và đo **đóng góp thật của fine-tuning** so với LLM pre-trained.

### D2. LoRA là gì? Vì sao dùng thay vì full fine-tuning?
- LoRA (Low-Rank Adaptation): đóng băng trọng số gốc, chỉ học 2 ma trận rank thấp `A ∈ R^{d×r}`, `B ∈ R^{r×d}` thêm vào các tầng linear.
- Ưu điểm:
  - Tham số huấn luyện rất nhỏ (thường 0.1–1% so với full).
  - Tiết kiệm VRAM, tốc độ train nhanh.
  - Checkpoint nhỏ → dễ chia sẻ, dễ swap.
- Full fine-tune model 3B/8B cần nhiều chục GB VRAM, không khả thi trên máy sinh viên.

### D3. Cấu hình LoRA chọn như thế nào?
Trong [src/generator_lora.py:41](src/generator_lora.py:41):
- `r = 16`: dung lượng vừa phải, cân bằng giữa biểu đạt và over-parameterization.
- `lora_alpha = 32` → scaling `α/r = 2`.
- `target_modules = [q_proj, k_proj, v_proj, o_proj, gate_proj, up_proj, down_proj]`: áp vào cả attention và MLP để học sâu hơn, quan trọng với task sinh text ngắn như nhãn.
- `dropout = 0.05`: nhẹ, chống overfit trên 5000 mẫu.
- `task_type = CAUSAL_LM` vì base là decoder-only.

### D4. QLoRA (4-bit) ở đâu, và tại sao dùng?
- `BitsAndBytesConfig(load_in_4bit=True, bnb_4bit_quant_type="nf4", double_quant=True, compute_dtype=fp16)`.
- Giảm VRAM base model xuống ~1/4, cho phép load Llama3-8B trên T4 (15GB).
- NF4 (Normal Float 4-bit) tốt hơn FP4 vì phân bố trọng số LLM xấp xỉ Gaussian.

### D5. Balanced Sampling trong train như thế nào?
[scripts/04_train_lora.py:57](scripts/04_train_lora.py:57):
- Gom tất cả utterance theo 7 nhãn.
- Lấy `TRAIN_MAX_SAMPLES / 7 ≈ 714` mẫu/lớp (nếu lớp đủ), lớp thiếu lấy hết.
- Bù phần còn thiếu bằng sampling ngẫu nhiên từ phần còn lại của pool.
- Shuffle trước khi train.
- Mục đích: tránh mô hình "học vẹt" đoán `neutral` cho mọi câu – nguyên nhân phổ biến khiến Macro-F1 thấp.

### D6. Early Stopping hoạt động ra sao?
- `EarlyStoppingCallback(patience=3)` trên `eval_loss`.
- `eval_strategy="steps"`, `eval_steps=50` → cứ 50 step lại đánh giá dev.
- `load_best_model_at_end=True` → cuối training dùng checkpoint có dev loss tốt nhất.
- Nếu 3 lần eval liên tiếp không cải thiện → dừng. Giúp tránh overfit và tiết kiệm thời gian.

### D7. Prompt format là gì, tại sao chọn ChatML?
[src/prompt_builder.py](src/prompt_builder.py): ChatML `<|im_start|>system … user … assistant|>`.
- Qwen2.5-Instruct và Phi-3 đều đã được align theo ChatML trong pre-training → tận dụng đúng format gốc cho kết quả tốt hơn.
- System prompt đặt "ràng buộc": chỉ trả về 1 từ trong 7 nhãn, mặc định `neutral` nếu không chắc → hạn chế ảo giác.

### D8. Sao không tính loss chỉ trên nhãn mà lại SFT toàn bộ prompt?
- Trong implementation hiện tại dùng SFTTrainer mặc định → tính loss toàn bộ text.
- Ưu điểm: đơn giản, vẫn hội tụ vì nhãn ngắn (1 từ) nằm cuối.
- Cải tiến có thể làm: dùng `DataCollatorForCompletionOnlyLM` hoặc mask loss trên phần prompt, chỉ tính loss trên nhãn để train hiệu quả hơn. Đây là hướng nâng cấp đã note.

### D9. Tại sao chọn Qwen2.5-3B-Instruct và Llama3-8B?
- **Qwen2.5-3B-Instruct**: nhỏ, nhanh, tiếng Anh mạnh, license thoải mái → baseline.
- **Llama3-8B**: lớn hơn, so sánh được scaling effect của mô hình.
- Cả hai đều có bản instruct nên phù hợp task phân loại bằng sinh text.

---

## E. Câu hỏi về Evaluation

### E1. Dùng metric nào và tại sao?
[src/evaluate.py](src/evaluate.py):
- **Accuracy**: chỉ tham khảo vì lớp mất cân bằng.
- **Macro-F1**: **quan trọng nhất** – trung bình F1 của 7 lớp đều nhau → đo khả năng nhận diện cả các lớp hiếm (`fear, disgust`).
- **Weighted-F1**: F1 trung bình theo tần suất lớp – cho nhìn tổng thể.
- **Micro-F1**: xấp xỉ Accuracy với multi-class → bổ sung tham chiếu.
- **Confusion Matrix**: hiển thị chi tiết nhầm lẫn giữa các lớp (thường `surprise/joy` hay lẫn, `fear/sadness` hay lẫn).

### E2. Xử lý output "ảo giác" (hallucination) của LLM thế nào?
- LLM có thể trả text không nằm trong 7 nhãn (ví dụ `"happy"`, `"the emotion is..."`).
- Trong [src/generator_lora.py:85](src/generator_lora.py:85): quét output, nếu chứa substring là 1 nhãn hợp lệ → map về nhãn đó; nếu không → `neutral`.
- Trong [src/evaluate.py:28](src/evaluate.py:28): lần nữa ép dự đoán ngoài 7 nhãn → `neutral`. Đây là fallback bảo đảm không crash metric.

### E3. Vì sao có zero-shot so với RAG trong 07_infer_eval_colab.py?
- Flag `--zero_shot` cho phép chạy **cùng model** nhưng **không dùng retrieval** → so sánh trực tiếp đóng góp của RAG.
- Kỳ vọng: RAG > Zero-shot về Macro-F1 (ví dụ được chọn gần với query giúp model định hướng tốt hơn cho các lớp hiếm).

### E4. Learning Curve dùng để làm gì?
- Vẽ train loss và val loss theo epoch ([src/evaluate.py:53](src/evaluate.py:53)).
- Mục đích sư phạm: chứng minh mô hình đang **học thật**, không overfit, và chỉ rõ điểm Early Stop.
- Nếu val loss tăng trong khi train giảm → dấu hiệu overfit → cần thêm regularization/tăng dropout.

---

## F. Câu hỏi mẹo & nâng cao

### F1. Nếu emotion_history bị sai ở bước inference (vì dùng nhãn dự đoán chứ không phải ground truth), ảnh hưởng gì?
- Trong [scripts/07_infer_eval_colab.py:105](scripts/07_infer_eval_colab.py:105) đang dùng `predicted_emo` để update `emotion_history` → **exposure bias**: lỗi tích lũy.
- Đây là đánh giá "realistic" (giống môi trường thật không có ground truth online) → phản ánh độ ổn định của hệ thống.
- Cải tiến: chạy kép, một lần teacher-forcing (dùng true history) để đo upper bound, một lần autoregressive như hiện tại.

### F2. Tại sao không dùng Cross-Encoder để re-rank?
- Cross-encoder chính xác hơn nhưng **chậm** (phải encode cặp query-candidate).
- Project ưu tiên tốc độ để chạy được trên CPU. Emotion shift là tín hiệu đặc thù của task, cross-encoder chưa chắc nắm bắt được.
- Có thể kết hợp: bi-encoder + cross-encoder + emotion shift làm 3-stage pipeline – là hướng mở rộng.

### F3. Có thể nhận vào input đa phương thức (audio, video) của MELD không?
- MELD gốc có audio + video + text. Project này chỉ dùng text.
- Lý do: giới hạn tài nguyên, pipeline RAG thuận tiện cho text.
- Mở rộng: dùng model như **HuBERT/Wav2Vec** cho audio, **CLIP/ViT** cho video, fuse feature trước khi classification. Đây là hướng multimodal ERC.

### F4. Giảng viên: "Project khác gì EmotionLLaMA / InstructERC đã công bố?"
- Khác điểm chính: **Emotion Shift Score** trong retrieval, cụ thể hóa khái niệm RAG theo pattern cảm xúc của nhân vật.
- Kết hợp **balanced sampling + LoRA + Early Stopping** trong một pipeline reproducible, có benchmark chéo 4 backend.
- Chúng tôi không tuyên bố SOTA; mục tiêu là thiết kế hệ thống có thể giải thích, đánh giá được đóng góp từng thành phần.

### F5. Overfitting xử lý thế nào khi train chỉ 5000 mẫu?
- Balanced sampling tránh model đoán nhãn đa số.
- `lora_dropout=0.05`, `gradient_checkpointing=True`.
- `EarlyStoppingCallback(patience=3)` + `load_best_model_at_end`.
- `paged_adamw_8bit` giữ lr ổn định.
- Nếu vẫn overfit: giảm `r` của LoRA, tăng dropout, hoặc early stop sớm hơn.

### F6. Groq 70B đã mạnh rồi, sao còn fine-tune model 3B?
- **Chi phí & quyền kiểm soát**: LoRA chạy local, không phụ thuộc API, không lộ dữ liệu.
- **Chứng minh năng lực kỹ thuật**: Fine-tuning là phần học thuật chính của môn.
- **So sánh**: LoRA 3B có thể xấp xỉ Groq 70B trên task hẹp → chứng minh giá trị của fine-tune trên domain cụ thể.

### F7. Vì sao `TOP_K = 2` mà không là 5 hoặc 10?
- Prompt dài → chi phí inference + VRAM tăng, dễ làm mô hình lơ câu query chính.
- Thử nghiệm cho thấy `K=2` là sweet spot: đủ ngữ cảnh để few-shot nhưng không làm loãng.
- Với model nhỏ (3B) context ngắn, K lớn giảm hiệu quả hơn là tăng.

### F8. Cải tiến trong tương lai?
1. Thay valence 1-D bằng embedding cảm xúc học được hoặc VAD 3-D.
2. Hybrid retriever: BM25 + bi-encoder + emotion shift.
3. Mask loss chỉ trên nhãn khi fine-tune SFT.
4. Mở rộng multimodal (audio/video) của MELD.
5. Đo hiệu năng với ensemble (Groq zero-shot + LoRA) bằng majority vote.
6. Thử DTW thay MAE cho chuỗi cảm xúc khác độ dài.

---

## G. Các câu hỏi nhanh

| Câu hỏi | Đáp án ngắn |
|---|---|
| Dataset? | MELD (text-only), 7 nhãn, declare-lab repo |
| Embedding? | `all-MiniLM-L6-v2`, 384-d, cosine |
| Vector DB? | ChromaDB persistent, HNSW cosine |
| Top-K retrieval? | 2 ví dụ, `fetch_k=max(8,20)=20` trước re-rank |
| Công thức mix? | `d = 0.7·sem + 0.3·emo_shift` |
| Base model fine-tune? | Qwen2.5-3B-Instruct, Llama3-8B |
| Kỹ thuật fine-tune? | QLoRA 4-bit NF4 + LoRA r=16 |
| Optimizer? | `paged_adamw_8bit`, lr=2e-4, batch=2×grad_acc=8 |
| Max epochs? | 5 (nhưng có Early Stopping patience=3) |
| Metric chính? | Macro-F1 (do lớp mất cân bằng) |
| Số mẫu test? | 2610 utterances |
| Chạy nhanh nhất? | Groq API (~5 phút) |
| Tác giả? | Source code tự viết, dataset MELD (declare-lab) |

---

## H. Note

1. **Luôn gắn câu trả lời về mã nguồn cụ thể**: trỏ tới file/hàm dòng số khi được hỏi – thể hiện bạn hiểu code mình viết.
2. **Trọng tâm cần nhấn mạnh**: Emotion Shift Score (điểm sáng tạo), Balanced Sampling + Early Stopping (chống học vẹt), 4 luồng generator (so sánh khoa học).
3. **Khi không biết**: thành thật nói "đây là hướng mở rộng chúng tôi chưa triển khai" thay vì bịa.
4. **Giải thích sai lầm**: nếu thầy chỉ ra bug (vd. loss tính trên cả prompt), thừa nhận và nêu hướng sửa – đánh giá cao tinh thần học hỏi.
5. **Demo**: nếu có thời gian, chạy 1 câu truy vấn live, show retrieval + prompt + kết quả → trực quan hơn slide.
