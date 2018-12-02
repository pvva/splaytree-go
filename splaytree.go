package splaytree

type Comparable interface {
	CompareTo(i interface{}) int
}

type splayNode struct {
	parent *splayNode
	left   *splayNode
	right  *splayNode
	value  Comparable
}

type SplayTree struct {
	root *splayNode
}

func NewSplayTree() *SplayTree {
	return &SplayTree{}
}

func (st *SplayTree) correctParent(n, t *splayNode) {
	if t.parent != nil {
		if n == t.parent.left {
			t.parent.left = t
		} else {
			t.parent.right = t
		}
	} else {
		st.root = t
	}
}

func (st *SplayTree) rotateLeft(n *splayNode) *splayNode {
	t := n.right
	t.parent = n.parent
	n.right = t.left

	if n.right != nil {
		n.right.parent = n
	}

	t.left = n
	n.parent = t
	st.correctParent(n, t)

	return t
}

func (st *SplayTree) rotateRight(n *splayNode) *splayNode {
	t := n.left
	t.parent = n.parent
	n.left = t.right

	if n.left != nil {
		n.left.parent = n
	}

	t.right = n
	n.parent = t
	st.correctParent(n, t)

	return t
}

func (st *SplayTree) splay(n *splayNode) {
	for n != st.root {
		p := n.parent
		if p == st.root {
			if n == p.left {
				st.rotateRight(p)
			} else if n == p.right {
				st.rotateLeft(p)
			}

			break
		}

		gp := p.parent
		nodeLeft := n == p.left
		nodeRight := n == p.right
		parentLeft := p == gp.left
		parentRight := p == gp.right

		if nodeLeft && parentLeft {
			// zig-zig right
			st.rotateRight(gp)
			st.rotateRight(p)
		} else if nodeRight && parentRight {
			// zig-zig left
			st.rotateLeft(gp)
			st.rotateLeft(p)
		} else if nodeRight && parentLeft {
			// zig-zag right
			st.rotateLeft(p)
			st.rotateRight(gp)
		} else if nodeLeft && parentRight {
			// zig-zag left
			st.rotateRight(p)
			st.rotateLeft(gp)
		}
	}
}

func (st *SplayTree) Insert(v Comparable) {
	if st.root == nil {
		st.root = &splayNode{
			value: v,
		}

		return
	}

	t := st.root
	var p *splayNode
	left := true
	for t != nil {
		p = t
		if t.value.CompareTo(v) > 0 {
			t = t.left
			left = true
		} else {
			t = t.right
			left = false
		}
	}

	n := &splayNode{
		parent: p,
		value:  v,
	}
	if left {
		p.left = n
	} else {
		p.right = n
	}

	st.splay(n)
}

func (st *SplayTree) find(v Comparable) *splayNode {
	n := st.root
	for n != nil {
		c := n.value.CompareTo(v)
		if c == 0 {
			return n
		} else if c > 0 {
			n = n.left
		} else {
			n = n.right
		}
	}

	return nil
}

func (st *SplayTree) transplant(n, nn *splayNode) {
	if n.parent == nil {
		st.root = nn
	} else if n == n.parent.left {
		n.parent.left = nn
	} else {
		n.parent.right = nn
	}
	if nn != nil {
		nn.parent = n.parent
	}
}

func (st *SplayTree) max(n *splayNode) *splayNode {
	t := n
	for t.right != nil {
		t = t.right
	}

	return t
}

func (st *SplayTree) Remove(v Comparable) bool {
	n := st.find(v)

	if n != nil {
		if n.left == nil {
			st.transplant(n, n.right)
		} else if n.right == nil {
			st.transplant(n, n.left)
		} else {
			st.splay(n)

			s := st.max(n.left)
			if s.parent.left == s {
				s.parent.left = nil
			} else {
				s.parent.right = nil
			}
			s.left = n.left
			s.left.parent = s
			s.right = n.right
			s.right.parent = s
			s.parent = nil
			st.root = s

		}

		return true
	}

	return false
}

func (st *SplayTree) Has(v Comparable) bool {
	n := st.find(v)
	if n != nil {
		st.splay(n)

		return true
	}

	return false
}

func (st *SplayTree) Traverse(f func(c Comparable, level int) bool) {
	st.traverse(st.root, 0, f)
}

func (st *SplayTree) traverse(from *splayNode, level int, f func(c Comparable, level int) bool) bool {
	if from == nil {
		return true
	}
	if !st.traverse(from.left, level+1, f) || !f(from.value, level) || !st.traverse(from.right, level+1, f) {
		return false
	}

	return true
}
