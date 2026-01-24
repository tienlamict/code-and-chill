package array

// maxArea returns the maximum amount of water a container can store.
//
// Two pointers:
// - Start with the widest container (left=0, right=n-1)
// - Area is limited by the shorter line, so we move the pointer at the shorter line
//   hoping to find a taller one while width decreases.
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	left, right := 0, len(height)-1
	best := 0

	for left < right {
		h := height[left]
		if height[right] < h {
			h = height[right]
		}

		area := h * (right - left)
		if area > best {
			best = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return best
}


