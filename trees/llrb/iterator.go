package llrb

import "github.com/helldealer/snowy-forest/stack"

/*
1. begin should less or equal with end
2. if begin is nil, iterate from min or to min, if end is nil, iterate to max or from max
3. iterate include begin and end elem
*/

func (t *Tree) Iterate(begin, end Elem, ascend bool, do func(elem Elem) bool) bool {
	if begin != nil && end != nil {
		if begin.Compare(end) > 0 {
			return false
		}
	}
	if ascend {
		return t.iterateAscend(begin, end, do)
	}
	return t.iterateDescend(begin, end, do)
}

func (t *Tree) iterateDescend(begin, end Elem, do func(elem Elem) bool) bool {
	s := stack.New(50)
	tmp := t.root
	res := -1
Loop:
	for tmp != nil || s.Len() > 0 {
		for tmp != nil {
			if end != nil {
				res = tmp.data.Compare(end)
			}
			if res < 0 {
				s.Push(tmp)
				tmp = tmp.right
			} else if res == 0 {
				s.Push(tmp)
				break
			} else {
				tmp = tmp.left
				continue Loop
			}
		}
		if s.Len() > 0 {
			top := s.Pop().(*node)
			if begin != nil {
				if top.data.Compare(begin) < 0 {
					return true
				}
			}
			if !do(top.data) {
				return false
			}
			tmp = top.left
		}
	}
	return true
}

func (t *Tree) iterateAscend(begin, end Elem, do func(elem Elem) bool) bool {
	s := stack.New(50)
	tmp := t.root
	res := 1
Loop:
	for tmp != nil || s.Len() > 0 {
		for tmp != nil {
			if begin != nil {
				res = tmp.data.Compare(begin)
			}
			if res > 0 {
				s.Push(tmp)
				tmp = tmp.left
			} else if res == 0 {
				s.Push(tmp)
				break
			} else {
				tmp = tmp.right
				continue Loop
			}
		}
		if s.Len() > 0 {
			top := s.Pop().(*node)
			if end != nil {
				if top.data.Compare(end) > 0 {
					return true
				}
			}
			if !do(top.data) {
				return false
			}
			tmp = top.right
		}
	}
	return true
}

func (t *Tree) IterateAll(do func(elem Elem) bool) bool {
	return t.iterateAll(t.root, do)
}

func (t *Tree) iterateAll(n *node, do func(elem Elem) bool) bool {
	if n == nil {
		return true
	}
	if n.left != nil {
		if !t.iterateAll(n.left, do) {
			return false
		}
	}
	if !do(n.data) {
		return false
	}
	if n.right != nil {
		if !t.iterateAll(n.right, do) {
			return false
		}
	}
	return true
}
