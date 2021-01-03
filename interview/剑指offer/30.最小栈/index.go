package index

type MinStack struct {
	stack []int
	minStack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minStack: []int{},
	}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	if len(this.minStack) == 0 || this.minStack[len(this.minStack)-1] >= x {
		this.minStack = append(this.minStack, x)
	}
}


func (this *MinStack) Pop()  {
	tmp := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if tmp == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}
