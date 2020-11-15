package tree

// Node represents a node in a rooted ordered tree.
type Node struct {
	Data interface{}

	Parent *Node

	Children []*Node
}

// New creates a new bare node holding any desired sort of data.
func New(data interface{}) *Node {
	return &Node{
		Data: data,
	}
}

// IsRoot returns true if and only if the specified node has no parent.
func (u *Node) IsRoot() bool {
	return u.Parent == nil
}

// IsLeaf returns true and if only if the specified node has no children.
func (u *Node) IsLeaf() bool {
	return u.Children == nil || len(u.Children) == 0
}

// AppendChild creates a parent-child relationship between an input node and
// another given node, appending it to its ordered list of children.
func (u *Node) AppendChild(child *Node) {
	if u.Children == nil {
		u.Children = make([]*Node, 1)
	}
	u.Children = append(u.Children, child)
	child.Parent = u
}

type callbackFunc func(u *Node)

// PreOrderTraversal performs pre-order traversal of the tree rooted at a given
// node, applying the callback function on each node before its children.
// This method uses a stack explicitly to track the nodes in the traversal.
func (u *Node) PreOrderTraversal(callback callbackFunc) {
	stack := []*Node{u}
	for len(stack) > 0 {
		currNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if currNode == nil {
			continue
		}
		callback(currNode)
		stack = append(stack, currNode.Children...)
	}
}

// PreOrderTraversalRecursive performs pre-order traversal of the tree rooted
// at a given node, applying the callback function on each node before its
// children. This method recursively calls itself in order to perform the traversal.
func (u *Node) PreOrderTraversalRecursive(callback callbackFunc) {
	if u == nil {
		return
	}
	callback(u)
	for _, child := range u.Children {
		child.PreOrderTraversalRecursive(callback)
	}
}

// PostOrderTraversal performs post-order traversal of the tree rooted at a
// given node, applying the callback function on each node after its children.
// This method uses a stack explicitly to track the nodes in the traversal.
func (u *Node) PostOrderTraversal(callback callbackFunc) {
	stack1 := []*Node{u}
	stack2 := []*Node{}
	for len(stack1) > 0 {
		currNode := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		if currNode == nil {
			continue
		}
		stack2 = append(stack2, currNode)
		if currNode.Children == nil {
			continue
		}
		stack1 = append(stack1, currNode.Children...)
	}
	for len(stack2) > 0 {
		currNode := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		callback(currNode)
	}
}

// PostOrderTraversal performs post-order traversal of the tree rooted at a
// given node, applying the callback function on each node after its children.
// This method recursively calls itself in order to perform the traversal.
func (u *Node) PostOrderTraversalRecursive(callback callbackFunc) {
	if u == nil {
		return
	}
	for _, child := range u.Children {
		child.PostOrderTraversalRecursive(callback)
	}
	callback(u)
}

// getHeightCollector returns a callback which collects the height of subtrees rooted
// at nodes of a tree in a given map.
func getHeightAccumulator(heightMap map[*Node]uint32) callbackFunc {
	return func(v *Node) {
		if v.IsLeaf() {
			heightMap[v] = 1
		} else {
			h := uint32(0)
			for _, child := range v.Children {
				if h < heightMap[child] {
					h = heightMap[child]
				}
			}
			heightMap[v] = h + 1
		}
	}
}

// Height computes the height of the tree rooted at a given node. Takes time
// linear in the size of that subtree.
func (u *Node) Height() uint32 {
	heightMap := map[*Node]uint32{}
	u.PostOrderTraversal(getHeightAccumulator(heightMap))
	return heightMap[u]
}

// HeightRecursive computes the height of the tree rooted at a given node,
// using a recursive function call. Takes time linear in the size of that subtree.
func (u *Node) HeightRecursive() uint32 {
	heightMap := map[*Node]uint32{}
	u.PostOrderTraversalRecursive(getHeightAccumulator(heightMap))
	return heightMap[u]
}

// getSizeCollector returns a callback which collects the size of subtrees rooted
// at nodes of a tree in a given map.
func getSizeAccumulator(sizeMap map[*Node]uint32) callbackFunc {
	return func(v *Node) {
		sizeMap[v] = 1
		for _, child := range v.Children {
			sizeMap[v] += sizeMap[child]
		}
	}
}

// Size computes the number of nodes in the subtree rooted at a given node.
// Takes time linear in the size of the subtree.
func (u *Node) Size() uint32 {
	sizeMap := map[*Node]uint32{}
	u.PostOrderTraversal(getSizeAccumulator(sizeMap))
	return sizeMap[u]
}

// SizeRecursive computes the number of nodes in the subtree rooted at a given
// node, using a recursive function call. Takes time linear in the size of the
// subtree.
func (u *Node) SizeRecursive() uint32 {
	sizeMap := map[*Node]uint32{}
	u.PostOrderTraversalRecursive(getSizeAccumulator(sizeMap))
	return sizeMap[u]
}
