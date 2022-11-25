package tree

type Tree interface {
	ID() int64
	Parent() int64
	AppendChildren(interface{})
	ChildrenNode() []Tree
}

func BuildTree(array []Tree) Tree {
	maxLen := len(array)
	var rootNode Tree = nil
	for i := 0; i < maxLen; i++ {
		count := 0
		for j := 0; j < maxLen; j++ {
			if array[j].ID() == array[i].Parent() {
				count++
				array[j].AppendChildren(array[i])
			}
		}
		if count == 0 && array[i].Parent() == 0 {
			rootNode = array[i]
		}
	}
	return rootNode
}

func BuildTreeByID(array []Tree, id int64) Tree {
	maxLen := len(array)
	var rootNode Tree = nil
	for i := 0; i < maxLen; i++ {
		count := 0
		for j := 0; j < maxLen; j++ {
			if array[j].ID() == array[i].Parent() {
				count++
				array[j].AppendChildren(array[i])
			}
		}
		if array[i].ID() == id {
			rootNode = array[i]
		}
	}
	return rootNode
}

func getTreeID(t Tree, ids *[]int64) {
	if t == nil {
		return
	}

	*ids = append(*ids, t.ID())

	for _, item := range t.ChildrenNode() {
		getTreeID(item, ids)
	}
}

func GetTreeID(tree Tree) []int64 {
	var ids []int64
	getTreeID(tree, &ids)
	return ids
}
