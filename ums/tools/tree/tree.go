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

func GetTreeID(tree Tree) []int64 {
	var ids []int64
	ids = append(ids, tree.ID())
	// 遍历菜单树
	rangeTree := func(tree Tree, ids *[]int64) {
		if len(tree.ChildrenNode()) == 0 {
			return
		}
		for _, item := range tree.ChildrenNode() {
			*ids = append(*ids, item.ID())
		}
	}
	rangeTree(tree, &ids)
	return ids
}
