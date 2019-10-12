package trees

type Loser struct {
	count int
	r     []int
	p     []Element
}

type Element interface {
	Lose(element Element) bool
}

func NewLoser(initElements []Element) *Loser {
	count := len(initElements)
	l := &Loser{
		count: count,
		r:     make([]int, count),
		p:     initElements,
	}
	for i := 0; i < count; i++ {
		l.r[i] = i
	}
	for k := count - 1; k >= 0; k-- {
		l.Update(k, initElements[k])
	}
	return l
}

func (l *Loser) Update(index int, element Element) bool {
	if index < 0 || index >= l.count {
		return false
	}
	l.p[index] = element
	pIndex := l.r[(index+l.count)/2]
	var lose = false
	for {
		if l.p[pIndex].Lose(element) {
			pIndex = l.r[pIndex/2]
			if pIndex == 0 {
				break
			}
			continue
		} else {
			l.r[pIndex] = index
			lose = true
		}
	}
	if !lose {
		l.r[0] = index
	}
	return true
}

func (l *Loser) Destroy() {

}
