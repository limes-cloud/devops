package model

import (
	"encoding/json"
	"fmt"
	"github.com/limeschool/gin"
	"time"
	"ums/consts"
	"ums/tools"
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

const (
	ALL_TEAMS_KEY = "ums_store_all_teams"
)

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

func (u *Team) Create(ctx *gin.Context) error {
	defer func() {
		tools.DelRedis(ctx, ALL_TEAMS_KEY)
	}()
	ctx.Redis(consts.REDIS).Del(ctx, ALL_TEAMS_KEY)
	user := CurUser(ctx)
	u.OperatorID = user.ID
	u.Operator = user.Name
	return transferErr(database(ctx).Table(u.Table()).Create(&u).Error)
}

func (u *Team) Tree(ctx *gin.Context, conds ...interface{}) (tree.Tree, error) {
	var list []*Team
	err := database(ctx).Table(u.Table()).Find(&list, conds...).Error
	if err != nil {
		return nil, err
	}
	var treeList []tree.Tree
	for _, item := range list {
		treeList = append(treeList, item)
	}
	return tree.BuildTree(treeList), nil
}

func (u *Team) All(ctx *gin.Context, conds ...interface{}) ([]*Team, error) {
	var list []*Team
	if store, _ := ctx.Redis(consts.REDIS).Get(ctx, ALL_TEAMS_KEY).Result(); store != "" {
		if json.Unmarshal([]byte(store), &list) != nil {
			return list, nil
		}
	}
	if err := database(ctx).Table(u.Table()).Find(&list, conds...).Error; err != nil {
		return nil, transferErr(err)
	}
	dataByte, _ := json.Marshal(list)
	ctx.Redis(consts.REDIS).Set(ctx, ALL_TEAMS_KEY, string(dataByte), time.Hour)
	return list, nil
}

func (u *Team) UpdateByID(ctx *gin.Context) error {
	defer func() {
		tools.DelRedis(ctx, ALL_TEAMS_KEY)
	}()
	ctx.Redis(consts.REDIS).Del(ctx, ALL_TEAMS_KEY)

	user := CurUser(ctx)
	u.Operator = user.Name
	u.OperatorID = user.ID
	return transferErr(database(ctx).Table(u.Table()).Updates(u).Error)
}

func (u *Team) DeleteByID(ctx *gin.Context, id int64) error {
	defer func() {
		tools.DelRedis(ctx, ALL_TEAMS_KEY)
	}()
	ctx.Redis(consts.REDIS).Del(ctx, ALL_TEAMS_KEY)

	var list []*Team
	err := database(ctx).Table(u.Table()).Find(&list).Error
	if err != nil {
		return err
	}
	var treeList []tree.Tree
	for _, item := range list {
		treeList = append(treeList, item)
	}
	t := tree.BuildTreeByID(treeList, id)
	ids := tree.GetTreeID(t)
	return transferErr(database(ctx).Table(u.Table()).Where("id in ?", ids).Delete(&u).Error)
}

func UserTeamIds(ctx *gin.Context, userId int64) []int64 {
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
	team := Team{}
	teamList, err := team.All(ctx)
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

func CurUserTeamIds(ctx *gin.Context) []int64 {
	var ids []int64
	userId := CurUser(ctx).ID
	redisKey := fmt.Sprintf("ums_user_ids_%v", userId)
	if result, _ := ctx.Redis(consts.REDIS).Get(ctx, redisKey).Result(); result != "" {
		if json.Unmarshal([]byte(result), &ids) == nil {
			return ids
		}
	}
	ids = UserTeamIds(ctx, CurUser(ctx).ID)
	byteData, _ := json.Marshal(ids)
	ctx.Redis(consts.REDIS).Set(ctx, redisKey, string(byteData), 5*time.Minute)
	return ids
}
