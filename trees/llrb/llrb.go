//llrb not tolerate elem repeat
package llrb

type Elem interface {
	Compare(elem Elem) int //return < 0: less; >0: more; ==0: equal
}

type Tree struct {
	root  *node
	count int
}

func New() *Tree {
	return &Tree{}
}

type node struct {
	isBlack     bool
	data        Elem
	left, right *node
}

func (n *node) rotateLeft() *node {
	tmp := n.right
	tmp.isBlack = n.isBlack
	n.right = tmp.left
	tmp.left = n
	n.isBlack = false
	return tmp
}

func (n *node) rotateRight() *node {
	tmp := n.left
	tmp.isBlack = true //always true when llrb
	n.left = tmp.right
	tmp.right = n
	n.isBlack = false
	return tmp
}

func (n *node) flip() *node {
	n.isBlack = false
	n.left.isBlack = true
	n.right.isBlack = true
	return n
}

func (n *node) rotateRightAndFlip() *node {
	tmp := n.left
	n.left = tmp.right
	tmp.right = n
	tmp.isBlack = false
	tmp.left.isBlack = true
	n.isBlack = true
	return tmp
}

func (n *node) isRed() bool {
	if n == nil {
		return false
	}
	return !n.isBlack
}

func (n *node) fix() (*node, bool) {
	if !n.left.isRed() && n.right.isRed() {
		return n.rotateLeft(), false
	}
	if n.left.isRed() {
		if n.left.left.isRed() {
			return n.rotateRightAndFlip(), false
		}
		if n.right.isRed() {
			return n.flip(), false
		}
		if n.isRed() {
			return n, false
		}
	}
	return n, true
}

func newNode(data Elem) *node {
	return &node{data: data}
}

func (t *Tree) fix(nodes []**node) {
	abort := false
	for i := len(nodes) - 1; i >= 0; i-- {
		*nodes[i], abort = (*nodes[i]).fix()
		if abort {
			break
		}
	}
	t.root.isBlack = true
}

func (t *Tree) Insert(elem Elem) bool {
	return t.insert(elem, false)
}

func (t *Tree) insert(elem Elem, replace bool) bool {
	n := newNode(elem)
	if t.count == 0 {
		n.isBlack = true
		t.root = n
		t.count++
		return true
	}
	path := make([]**node, 0, 50)
	path = append(path, &t.root)
	tmp := t.root
	for {
		res := n.data.Compare(tmp.data)
		if res < 0 {
			if tmp.left == nil {
				tmp.left = n
				t.count++
				if tmp.isBlack {
					return true
				}
				//rotate right, then color flip, then...
				break
			}
			path = append(path, &tmp.left)
			tmp = tmp.left
			continue
		} else if res == 0 {
			if replace {
				tmp.data = elem
				return true
			}
			return false
		}
		if tmp.right == nil {
			//rotate left or flip
			tmp.right = n
			t.count++
			break
		}
		path = append(path, &tmp.right)
		tmp = tmp.right
	}
	t.fix(path)
	return true
}

func (t *Tree) Update(new Elem) bool {
	if node := t.get(new); node != nil {
		node.data = new
		return true
	}
	return false
}

func (t *Tree) InsertOrUpdate(elem Elem) {
	t.insert(elem,true)
}

func (t *Tree) Get(elem Elem) Elem {
	if node := t.get(elem); node != nil {
		return node.data
	}
	return nil
}

func (t *Tree) get(elem Elem) *node {
	tmp := t.root
	for tmp != nil {
		res := elem.Compare(tmp.data)
		if res < 0 {
			tmp = tmp.left
		} else if res > 0 {
			tmp = tmp.right
		}else {
			return tmp
		}
	}
	return nil
}

func (t *Tree) Has(elem Elem) bool {
	return t.Get(elem) != nil
}

func (t *Tree) Delete(elem Elem) bool {
	return false
}

func (t *Tree) Len() int {
	return t.count
}

func (t *Tree) Max() Elem {
	tmp := t.root
	if tmp == nil {
		return nil
	}
	for ; tmp.right != nil; tmp = tmp.right {}
	return tmp.data
}

func (t *Tree) Min() Elem {
	tmp := t.root
	if tmp == nil {
		return nil
	}
	for ; tmp.left != nil; tmp = tmp.left {}
	return tmp.data
}
