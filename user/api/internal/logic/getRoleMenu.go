package logic

import (
	"context"
	"devops/common/errorx"
	"devops/common/meta"
	"devops/common/tools"
	"devops/common/tools/tree"
	"devops/user/models"
	"errors"
	"gorm.io/gorm"

	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleMenuLogic {
	return &GetRoleMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleMenuLogic) GetRoleMenu(req *types.GetRoleMenusRequest) (resp *types.GetMenuResponse, err error) {

	if meta.IsSuper(l.ctx) { // 系统角色
		menuServer := NewGetMenuLogic(l.ctx, l.svcCtx)
		return menuServer.GetMenu(&types.GetMenuRequest{
			IsFilter: true,
		})
	}

	// 客户端调用
	var list []models.Menu
	resp = new(types.GetMenuResponse)
	var menuIds []int64

	//获取当前用户的menu_id
	roleMenu := models.RoleMenu{}
	if req.RoleID == 0 {
		req.RoleID = meta.UserRoleId(l.ctx)
	}

	roleMenus, _, _ := roleMenu.All(req)
	for _, item := range roleMenus {
		menuIds = append(menuIds, item.MenuID)
	}

	menu := models.Menu{}
	list, _, err = menu.AllCall(func(db *gorm.DB) *gorm.DB {
		return db.Where("id in ?", menuIds)
	})
	if err != nil {
		return nil, errors.New(errorx.SqlErr)
	}

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
