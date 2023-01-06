package model

import (
	"encoding/json"
	"github.com/limeschool/gin"
	"time"
	"ums/consts"
	"ums/tools/lock"
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
	Redirect   *string `json:"redirect"`
	ParentID   int64   `json:"parent_id"`
	Weight     *int    `json:"weight"`
	Hidden     *bool   `json:"hidden"`
	IsFrame    *bool   `json:"is_frame"`
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

	if u.Permission == consts.BaseApi {
		delayDelCache(ctx, consts.RedisBaseApi)
	}

	return transferErr(database(ctx).Table(u.Table()).Create(&u).Error)
}

func (u *Menu) One(ctx *gin.Context, cond ...interface{}) error {
	return transferErr(database(ctx).Table(u.Table()).First(u, cond...).Error)
}

func (u *Menu) GetBaseApiPath(ctx *gin.Context) []*Menu {
	// 从缓存中加载数据
	loadMenus := func(ctx *gin.Context) ([]*Menu, error) {
		str, err := cache(ctx).Get(ctx, consts.RedisBaseApi).Result()
		if err != nil {
			return nil, err
		}
		var resp []*Menu
		if json.Unmarshal([]byte(str), &resp) != nil {
			return nil, err
		}
		return resp, nil
	}

	// 存储数据到缓存
	storeMenus := func(ctx *gin.Context, list []*Menu) {
		byteData, _ := json.Marshal(list)
		ctx.Redis(consts.REDIS).Set(ctx, consts.RedisBaseApi, string(byteData), 1*time.Hour)
	}

	// 分布式锁的key
	lockKey := consts.RedisBaseApi + "_lock"

	// 获取缓存
	if list, err := loadMenus(ctx); err == nil {
		return list
	}

	// 获取锁
	lk := lock.NewLock(ctx, lockKey)
	lk.Acquire()
	defer lk.Release()

	// 重新获取缓存
	if list, err := loadMenus(ctx); err == nil {
		return list
	}

	// 查询数据库
	list, _ := u.All(ctx, "permission = ? and type = 'A'", consts.BaseApi)
	storeMenus(ctx, list)
	return list
}

func (u *Menu) OneById(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).First(u, "id = ?", id).Error)
}

func (u *Menu) All(ctx *gin.Context, cond ...interface{}) ([]*Menu, error) {
	var list []*Menu
	return list, transferErr(database(ctx).Table(u.Table()).Order("weight desc").Find(&list, cond...).Error)
}

func (u *Menu) Tree(ctx *gin.Context, cond ...interface{}) (tree.Tree, error) {
	list, err := u.All(ctx, cond...)
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

	if u.Permission == consts.BaseApi {
		delayDelCache(ctx, consts.RedisBaseApi)
	}

	return transferErr(database(ctx).Table(u.Table()).Updates(u).Error)
}

func (u *Menu) Delete(ctx *gin.Context, cond ...interface{}) error {
	if err := database(ctx).Table(u.Table()).First(u, cond...).Error; err != nil {
		return transferErr(err)
	}

	if u.Permission == consts.BaseApi {
		delayDelCache(ctx, consts.RedisBaseApi)
	}
	return transferErr(database(ctx).Table(u.Table()).Delete(u).Error)
}
