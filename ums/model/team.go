package model

import (
	"encoding/json"
	"fmt"
	"github.com/limeschool/gin"
	"time"
	"ums/consts"
	"ums/errors"
	"ums/tools/lock"
	"ums/tools/tree"
)

type Team struct {
	gin.BaseModel
	Name        string  `json:"name"`
	Avatar      string  `json:"avatar"`
	Description string  `json:"description"`
	ParentID    int64   `json:"parent_id"`
	Operator    string  `json:"operator"`
	OperatorID  int64   `json:"operator_id"`
	Children    []*Team `json:"children" gorm:"-"`
}

func (u *Team) ID() int64 {
	return u.BaseModel.ID
}

func (u *Team) Parent() int64 {
	return u.ParentID
}

func (u *Team) AppendChildren(child any) {
	team := child.(*Team)
	u.Children = append(u.Children, team)
}

func (u *Team) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range u.Children {
		list = append(list, item)
	}
	return list
}

func (u *Team) Table() string {
	return "team"
}

// Create 创建部门
func (u *Team) Create(ctx *gin.Context) error {
	delayDelCache(ctx, u.cacheTeamsKey())
	u.clearTeamIdsCache(ctx)

	user := CurUser(ctx)
	u.OperatorID = user.ID
	u.Operator = user.Name

	return transferErr(database(ctx).Table(u.Table()).Create(&u).Error)
}

// Tree 获取部门树
func (u *Team) Tree(ctx *gin.Context) (tree.Tree, error) {
	var list []*Team
	err := database(ctx).Table(u.Table()).Find(&list).Error
	if err != nil {
		return nil, err
	}
	var treeList []tree.Tree
	for _, item := range list {
		treeList = append(treeList, item)
	}
	return tree.BuildTree(treeList), nil
}

// AllByCache 获取部门缓存
func (u *Team) AllByCache(ctx *gin.Context, key string) ([]*Team, error) {
	var teams []*Team
	resByte, err := cache(ctx).Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	if len(resByte) == 0 {
		return nil, errors.DBNotFoundError
	}
	if err = json.Unmarshal(resByte, &teams); err != nil {
		return nil, nil
	}
	return teams, nil
}

func (u *Team) cacheTeamsKey() string {
	return "ums_store_all_teams"
}

// All 获取全部部门
func (u *Team) All(ctx *gin.Context) ([]*Team, error) {
	key := "ums_team_lock"
	if teams, err := u.AllByCache(ctx, u.cacheTeamsKey()); err == nil {
		return teams, nil
	}

	// 加锁,防止缓存击穿
	rl := lock.NewLock(ctx, key)
	rl.Acquire()
	defer rl.Release()

	// 获取锁之后重新查询缓存
	if teams, err := u.AllByCache(ctx, u.cacheTeamsKey()); err == nil {
		return teams, nil
	}

	var list []*Team
	if err := database(ctx).Table(u.Table()).Find(&list).Error; err != nil {
		return nil, transferErr(err)
	}

	dataByte, _ := json.Marshal(list)
	cache(ctx).Set(ctx, u.cacheTeamsKey(), string(dataByte), time.Hour)
	return list, nil
}

// Update 更新部门信息
func (u *Team) Update(ctx *gin.Context) error {
	delayDelCache(ctx, u.cacheTeamsKey())
	u.clearTeamIdsCache(ctx)

	user := CurUser(ctx)
	u.Operator = user.Name
	u.OperatorID = user.ID
	return transferErr(database(ctx).Table(u.Table()).Updates(u).Error)
}

func (u *Team) DeleteByID(ctx *gin.Context, id int64) error {

	// 获取全部部门
	list, err := u.All(ctx)
	if err != nil {
		return transferErr(err)
	}

	// 组装成树状结构
	var treeList []tree.Tree
	for _, item := range list {
		treeList = append(treeList, item)
	}
	t := tree.BuildTreeByID(treeList, id)

	// 获取当前部门的下级树的id
	ids := tree.GetTreeID(t)

	// 延迟双删
	delayDelCache(ctx, u.cacheTeamsKey())
	u.clearTeamIdsCache(ctx)

	// 进行数据删除
	return transferErr(database(ctx).Table(u.Table()).Where("id in ?", ids).Delete(&u).Error)
}

func (u *Team) UserTeamIds(ctx *gin.Context, userId int64) []int64 {
	var ids = make([]int64, 0)
	user := User{}
	if user.OneByID(ctx, userId) != nil {
		return ids
	}

	// 判断用户权限
	if user.Role.DataScope == consts.CURTEAM {
		return []int64{user.TeamID}
	}
	if user.Role.DataScope == consts.CUSTOM {
		_ = json.Unmarshal([]byte(user.Role.TeamIds), &ids)
		return ids
	}

	// 获取部门
	teamList, err := u.All(ctx)
	if err != nil {
		return ids
	}
	var treeList []tree.Tree
	for _, item := range teamList {
		treeList = append(treeList, item)
	}

	switch user.Role.DataScope {
	case consts.ALLTEAM:
		teamTree := tree.BuildTreeByID(treeList, user.TeamID)
		ids = tree.GetTreeID(teamTree)
	case consts.DOWNTEAM:
		teamTree := tree.BuildTreeByID(treeList, user.TeamID)
		ids = tree.GetTreeID(teamTree)
		if len(ids) > 2 {
			ids = ids[1:]
		} else {
			ids = []int64{}
		}
	}
	return ids
}

func (u *Team) AllTeamIdsByCache(ctx *gin.Context, key string) ([]int64, error) {
	var respArr []int64
	resByte, err := cache(ctx).Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	if len(resByte) == 0 {
		return nil, errors.DBNotFoundError
	}
	if err = json.Unmarshal(resByte, &respArr); err != nil {
		return []int64{}, nil
	}
	return respArr, nil
}

func (u *Team) cacheTeamIdsKey(uid any) string {
	return fmt.Sprintf("ums_teams_ids_%v", uid)
}

func (u *Team) clearTeamIdsCache(ctx *gin.Context) {
	list, _ := cache(ctx).Keys(ctx, u.cacheTeamIdsKey("*")).Result()
	cache(ctx).Del(ctx, list...)
	go func() {
		time.Sleep(1 * time.Second)
		cache(ctx).Del(ctx, list...)
	}()
}

func CurUserTeamIds(ctx *gin.Context) []int64 {
	userId := CurUser(ctx).ID

	team := Team{}
	if teamIds, err := team.AllTeamIdsByCache(ctx, team.cacheTeamIdsKey(userId)); err == nil {
		return teamIds
	}

	ids := team.UserTeamIds(ctx, CurUser(ctx).ID)
	byteData, _ := json.Marshal(ids)
	cache(ctx).Set(ctx, team.cacheTeamIdsKey(userId), string(byteData), 5*time.Minute)
	return ids
}
