package tree

type Tree interface {
	ID() int64
	ParentID() int64
	AppendChildren(interface{})
}

func BuildTree(array []Tree) Tree {
	maxLen := len(array)
	var rootNode Tree = nil
	for i := 0; i < maxLen; i++ {
		count := 0
		for j := 0; j < maxLen; j++ {
			if array[j].ID() == array[i].ParentID() {
				count++
				array[j].AppendChildren(array[i])
			}
		}
		if count == 0 {
			rootNode = array[i]
		}
	}
	return rootNode
}
