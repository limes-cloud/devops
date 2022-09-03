package service

import (
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"ums/errors"
	"ums/model"
	"ums/tools"
	"ums/tools/tree"
	"ums/types"
)

func AllTeam(ctx *gin.Context) (tree.Tree, error) {
	team := model.Team{}
	return team.Tree(ctx)
}

func AddTeam(ctx *gin.Context, in *types.AddTeamRequest) error {
	team := model.Team{}
	if !tools.InList(model.CurUserTeamIds(ctx), in.ParentID) {
		return errors.NotAddTeamError
	}

	if copier.Copy(&team, in) != nil {
		return errors.AssignError
	}
	return team.Create(ctx)
}

func UpdateTeam(ctx *gin.Context, in *types.UpdateTeamRequest) error {
	team := model.Team{}
	if in.ParentID != 0 && in.ID == in.ParentID {
		return errors.TeamParentIdError
	}
	if !tools.InList(model.CurUserTeamIds(ctx), in.ID) {
		return errors.NotEditTeamError
	}
	if copier.Copy(&team, in) != nil {
		return errors.AssignError
	}
	return team.UpdateByID(ctx)
}

func DeleteTeam(ctx *gin.Context, in *types.DeleteTeamRequest) error {
	team := model.Team{}
	if !tools.InList(model.CurUserTeamIds(ctx), in.ID) {
		return errors.NotDelTeamError
	}
	return team.DeleteByID(ctx, in.ID)
}
