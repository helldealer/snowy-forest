package stack

type Stack struct {
	pool  []interface{}
	count int
}

func New(init int) *Stack {
	return &Stack{pool: make([]interface{}, 0, init)}
}

func (s *Stack) Pop() interface{} {
	if s.count == 0 {
		return nil
	}
	//rough, make it option
	if s.count >= 20 && 5 * s.count <= cap(s.pool) {
		tmp := s.pool
		s.pool = make([]interface{}, s.count, 2*s.count)
		copy(s.pool, tmp)
	}
	s.count--
	ret := s.pool[s.count]
	s.pool = s.pool[:s.count]
	return ret
}

func (s *Stack) Push(in interface{}) {
	s.pool = append(s.pool, in)
	s.count++
}

func (s *Stack) Peek() interface{} {
	if s.count == 0 {
		return nil
	}
	return s.pool[s.count-1]
}

func (s *Stack) Len() int {
	return s.count
}
