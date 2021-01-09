package index

type MaxQueue struct {
	queue          []int
	queueLength    int
	maxQueue       []int
	maxQueueLength int
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue:    []int{},
		maxQueue: []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.maxQueue) == 0 {
		return -1
	}
	return this.maxQueue[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	this.queueLength++
	for this.maxQueueLength != 0 && this.maxQueue[this.maxQueueLength-1] < value {
		this.maxQueueLength--
	}
	this.maxQueue = this.maxQueue[0 : this.maxQueueLength]

	this.maxQueue = append(this.maxQueue, value)
	this.maxQueueLength++
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	tmp := this.queue[0]
	this.queue = this.queue[1:]
	this.queueLength--
	if tmp == this.maxQueue[0] {
		this.maxQueue = this.maxQueue[1:]
		this.maxQueueLength--
	}
	return tmp
}
