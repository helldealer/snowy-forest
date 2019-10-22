/*
loser tree is for minimum or maximum value of an given array
*/
package loser

type Elem interface {
	Lose(LoserElem Elem) bool
}

type Tree struct {
	count int
	r     []int
	p     []Elem
}

func New(initLoserElements []Elem, absoluteWinner Elem) *Tree {
	count := len(initLoserElements)
	l := &Tree{
		count: count,
		r:     make([]int, count),
		p:     append(initLoserElements, absoluteWinner),
	}
	for i := 0; i < count; i++ {
		l.r[i] = count
	}
	for k := 0; k < count; k++ {
		l.Update(k, initLoserElements[k])
	}
	return l
}

//param index must be the last time winner's index
//so the loser tree's update op is just for element which is the last time winner
func (l *Tree) Update(index int, LoserElem Elem) (int, Elem) {
	if index < 0 || index >= l.count {
		return -1, nil
	}
	l.p[index] = LoserElem
	winner := index
	father := (index + l.count) / 2
	for {
		if l.p[winner].Lose(l.p[l.r[father]]) {
			l.r[father], winner = winner, l.r[father]
		}
		father = father / 2
		if father == 0 {
			break
		}
	}
	l.r[0] = winner
	return l.r[0], l.p[l.r[0]]
}

//output winner to loser
//after iterate, the tree is full of absolute loser
func (l *Tree) Iterate(absoluteLoser Elem) []Elem {
	var res = make([]Elem, l.count)
	res[0] = l.p[l.r[0]]
	for i := 1; i < l.count; i++ {
		_, e := l.Update(l.r[0], absoluteLoser)
		res[i] = e
	}
	return res
}

func (l *Tree) Winner() (int, Elem) {
	return l.r[0], l.p[l.r[0]]
}
