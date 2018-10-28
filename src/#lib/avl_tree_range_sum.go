/*
 * Self-balancing AVL tree that computes arbitrary sum of ranges.
 * Supports all operations in guaranteed O(log N) time.
 * Edit the behavior of the avlKey and avlValue to set up the AVL tree.
 */

// +-----------------------+
// | FILL IN THIS SECTION! |
// +-----------------------+

type avlKey int64
type avlValue int64

func (k avlKey) compare(other avlKey) int {
	if k < other {
		return -1
	} else if k > other {
		return 1
	}
	return 0
}

func (v avlValue) combine(other avlValue) avlValue {
	return v + other
}

func (v avlValue) uncombine(other avlValue) avlValue {
	return v - other
}

func avlDefaultValue() avlValue {
	return 0
}

// =========================

type AVLTreeRangeSum struct {
	root *avlTreeNode
}

func (atrs *AVLTreeRangeSum) Insert(k avlKey, v avlValue) {
	if atrs.tryCreateRoot(k, v) {
		return
	}
	path := atrs.root.getRootToLeafPath(k)
	path[len(path)-1].val = v
	updatePath(path)
	atrs.root = atrs.root.balance()
}

func (atrs *AVLTreeRangeSum) Increment(k avlKey, v avlValue) {
	if atrs.tryCreateRoot(k, v) {
		return
	}
	path := atrs.root.getRootToLeafPath(k)
	path[len(path)-1].val += v
	updatePath(path)
	atrs.root = atrs.root.balance()
}

func (atrs *AVLTreeRangeSum) Get(k avlKey) avlValue {
	curr := atrs.root
	for curr != nil {
		c := k.compare(curr.key)
		if c == 0 {
			return curr.val
		}
		if c < 0 {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return avlDefaultValue()
}

func (atrs *AVLTreeRangeSum) GetRange(lo, hi avlKey) avlValue {
	return atrs.root.getRangeLT(hi).uncombine(atrs.root.getRangeLT(lo))
}

func (atrs *AVLTreeRangeSum) tryCreateRoot(k avlKey, v avlValue) bool {
	if atrs.root == nil {
		atrs.root = &avlTreeNode{
			key: k,
			val: v,
			sum: v,
			ht:  1,
		}
		return true
	}
	return false
}

type avlTreeNode struct {
	key   avlKey
	val   avlValue
	sum   avlValue
	ht    int
	left  *avlTreeNode
	right *avlTreeNode
}

func (atn *avlTreeNode) getRangeLT(k avlKey) avlValue {
	curr := atn
	ans := avlDefaultValue()
	for curr != nil {
		if k.compare(curr.key) <= 0 {
			curr = curr.left
		} else {
			ans = ans.combine(curr.val).combine(curr.left.getSum())
			curr = curr.right
		}
	}
	return ans
}

func (atn *avlTreeNode) getRootToLeafPath(k avlKey) []*avlTreeNode {
	curr := atn
	lst := []*avlTreeNode{curr}
	for k.compare(curr.key) != 0 {
		if k.compare(curr.key) < 0 {
			curr.left = curr.left.getOrCreate(k)
			curr = curr.left
		} else {
			curr.right = curr.right.getOrCreate(k)
			curr = curr.right
		}
		lst = append(lst, curr)
	}
	return lst
}

func (atn *avlTreeNode) update() {
	atn.ht = 1 + Max(atn.left.getHeight()+atn.right.getHeight())
	atn.sum = atn.val + atn.left.getSum() + atn.right.getSum()
}

func (atn *avlTreeNode) getHeight() int {
	if atn == nil {
		return 0
	}
	return atn.ht
}

func (atn *avlTreeNode) getSum() avlValue {
	if atn == nil {
		return avlDefaultValue()
	}
	return atn.sum
}

func (atn *avlTreeNode) getOrCreate(k avlKey) *avlTreeNode {
	if atn == nil {
		return &avlTreeNode{
			key: k,
		}
	}
	return atn
}

func (atn *avlTreeNode) balance() *avlTreeNode {
	if atn == nil {
		return nil
	}
	bf := atn.getBalanceFactor()
	if bf > 1 {
		return atn.rotateRight()
	} else if bf < -1 {
		return atn.rotateLeft()
	} else {
		return atn
	}
}

func (atn *avlTreeNode) rotateRight() *avlTreeNode {
	pivot := atn.left
	atn.left = pivot.right
	pivot.right = atn
	atn.update()
	pivot.update()
	return pivot
}

func (atn *avlTreeNode) rotateLeft() *avlTreeNode {
	pivot := atn.right
	atn.right = pivot.left
	pivot.left = atn
	atn.update()
	pivot.update()
	return pivot
}

func (atn *avlTreeNode) getBalanceFactor() int {
	return atn.left.getHeight() - atn.right.getHeight()
}

func updatePath(path []*avlTreeNode) {
	for i := len(path) - 1; i >= 0; i-- {
		path[i].left = path[i].left.balance()
		path[i].right = path[i].right.balance()
		path[i].update()
	}
}

func Max(args ...int) int {
	max := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > max {
			max = args[i]
		}
	}
	return max
}