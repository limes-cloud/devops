package logic

import (
	"context"
	"devops/common/tools"
	"devops/common/tools/tree"
	"devops/common/typex"
	"devops/user/models"
	"gorm.io/gorm"

	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuLogic {
	return &GetMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuLogic) GetMenu(req *types.GetMenuRequest) (resp *types.GetMenuResponse, err error) {
	var list []models.Menu
	resp = new(types.GetMenuResponse)

	menu := models.Menu{}
	list, _, err = menu.All(nil, func(db *gorm.DB) *gorm.DB {
		if req.IsFilter {
			db = db.Where("permission != ?", typex.BaseApiKey)
		}
		return db
	})

	var menuTree []MenuTree
	tools.Transform(list, &menuTree)
	nodeArray := make([]tree.Tree, len(menuTree))
	for i := 0; i < len(menuTree); i++ {
		nodeArray[i] = &menuTree[i]
	}
	//进行转菜单树
	root := tree.BuildTree(nodeArray)
	tools.Transform(root, resp)
	return
}

type MenuTree struct {
	types.GetMenuResponse
}

func (t *MenuTree) ID() int64 {
	return t.GetMenuResponse.ID
}

func (t *MenuTree) ParentID() int64 {
	return t.GetMenuResponse.ParentID
}

func (t *MenuTree) AppendChildren(node interface{}) {
	n := node.(*MenuTree)
	t.Children = append(t.Children, &n.GetMenuResponse)
}
