package model

import (
	"github.com/limeschool/gin"
	"ums/tools/tree"
)

type Menu struct {
	Title      string  `json:"title"`
	Icon       string  `json:"icon"`
	Path       string  `json:"path"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	Permission string  `json:"permission"`
	Method     string  `json:"method"`
	Component  string  `json:"component"`
	Redirect   string  `json:"redirect"`
	ParentID   int64   `json:"parent_id"`
	Weight     int     `json:"weight"`
	Hidden     bool    `json:"hidden"`
	IsFrame    bool    `json:"is_frame"`
	Operator   string  `json:"operator"`
	OperatorID int64   `json:"operator_id"`
	Children   []*Menu `json:"children" gorm:"-"`
	gin.BaseModel
}

func (u *Menu) ID() int64 {
	return u.BaseModel.ID
}

func (u *Menu) Parent() int64 {
	return u.ParentID
}

func (u *Menu) AppendChildren(child any) {
	menu := child.(*Menu)
	u.Children = append(u.Children, menu)
}

func (u *Menu) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range u.Children {
		list = append(list, item)
	}
	return list
}

func (u *Menu) Table() string {
	return "menu"
}

func (u *Menu) Create(ctx *gin.Context) error {
	user := CurUser(ctx)
	u.OperatorID = user.ID
	u.Operator = user.Name
	return database(ctx).Table(u.Table()).Create(&u).Error
}

func (u *Menu) One(ctx *gin.Context, conds ...interface{}) error {
	return database(ctx).Table(u.Table()).First(u, conds...).Error
}

func (u *Menu) OneById(ctx *gin.Context, id int64) error {
	return database(ctx).Table(u.Table()).First(u, "id = ?", id).Error
}

func (u *Menu) All(ctx *gin.Context, conds ...interface{}) ([]*Menu, error) {
	var list []*Menu
	return list, database(ctx).Table(u.Table()).Order("weight desc").Find(&list, conds...).Error
}

func (u *Menu) Tree(ctx *gin.Context, conds ...interface{}) (tree.Tree, error) {
	list, err := u.All(ctx, conds...)
	if err != nil {
		return nil, err
	}
	var treeList []tree.Tree
	for _, item := range list {
		treeList = append(treeList, item)
	}
	return tree.BuildTree(treeList), nil
}

func (u *Menu) UpdateByID(ctx *gin.Context) error {
	user := CurUser(ctx)
	u.Operator = user.Name
	u.OperatorID = user.ID
	return database(ctx).Table(u.Table()).Updates(u).Error
}

func (u *Menu) Delete(ctx *gin.Context, conds ...interface{}) error {
	return database(ctx).Table(u.Table()).Delete(u, conds...).Error
}
