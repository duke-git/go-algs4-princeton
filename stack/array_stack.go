package stack

type ArrayStack struct {
	items  []string
	length int
}

func NewArrayStack(capacity int) *ArrayStack {
	return &ArrayStack{
		items:  make([]string, 0, capacity),
		length: 0,
	}
}

func (s *ArrayStack) IsEmpty() bool {
	return s.length == 0
}

func (s *ArrayStack) Push(item string) {
	capacity := len(s.items)
	if s.length == capacity {
		s.Resize(2 * s.length)
	}
	s.length++
	s.items[s.length] = item
}

func (s *ArrayStack) Pop() string {
	capacity := len(s.items)
	item := s.items[s.length]
	s.length--
	s.items[s.length] = ""
	if s.length > 0 && s.length == capacity/4 {
		s.Resize(capacity / 2)
	}
	return item
}

func (s *ArrayStack) Resize(capacity int) {
	copy := make([]string, 0, capacity)
	for i := 0; i < s.length; i++ {
		copy[i] = s.items[i]
	}
	s.items = copy
}
