package logic

import (
	"configure/common/tools"
	"configure/common/tools/tree"
	"configure/models"
	"context"

	"configure/api/internal/svc"
	"configure/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTeamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeamLogic {
	return &GetTeamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTeamLogic) GetTeam() (resp *types.GetTeamResponse, err error) {
	var list []models.Team
	resp = new(types.GetTeamResponse)

	team := models.Team{}
	tb := l.svcCtx.Orm.Table(team.Table())
	if err = tb.Find(&list).Error; err != nil {
		return nil, err
	}

	var teamTree []TeamTree
	tools.Transform(list, &teamTree)

	nodeArray := make([]tree.Tree, len(teamTree))
	for i := 0; i < len(teamTree); i++ {
		nodeArray[i] = &teamTree[i]
	}

	//进行转菜单树
	root := tree.BuildTree(nodeArray)
	tools.Transform(root, resp)
	return
}

type TeamTree struct {
	types.GetTeamResponse
}

func (t *TeamTree) ID() int64 {
	return t.GetTeamResponse.ID
}

func (t *TeamTree) ParentID() int64 {
	return t.GetTeamResponse.ParentID
}

func (t *TeamTree) AppendChildren(node interface{}) {
	n, _ := node.(TeamTree)
	t.Children = append(t.Children, n.GetTeamResponse)
}
