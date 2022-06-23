package problem

// Sliding Window Maximum
// https://leetcode.com/problems/sliding-window-maximum/

type Deque struct {
	inner []int
}

func (d *Deque) PushFront(v int) {
	d.inner = append([]int{v}, d.inner...)
}

func (d *Deque) PopFront() int {
	front := d.Front()
	d.inner = d.inner[1:]
	return front
}

func (d *Deque) Push(v int) {
	d.inner = append(d.inner, v)
}

func (d *Deque) Pop() int {
	back := d.Back()
	d.inner = d.inner[:len(d.inner)-1]
	return back
}

func (d *Deque) Front() int {
	return d.inner[0]
}

func (d *Deque) Back() int {
	return d.inner[len(d.inner)-1]
}

func (d *Deque) Len() int {
	return len(d.inner)
}

func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	deque := &Deque{}

	for i := range nums {
		for deque.Len() > 0 && deque.Front() < i-k+1 {
			deque.PopFront()
		}

		for deque.Len() > 0 && nums[deque.Back()] < nums[i] {
			deque.Pop()
		}

		deque.Push(i)

		if i >= k-1 {
			ans = append(ans, nums[deque.Front()])
		}
	}

	return ans
}
