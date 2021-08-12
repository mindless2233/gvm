package rtda

type Thread struct {
	pc    int
	stack *stack
}

func newThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame(frame *Frame) {
	self.stack.pop(frame)
}

func (self *Thread) CurrentFrame(frame *Frame) {
	self.stack.top(frame)
}
