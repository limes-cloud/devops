package logic

import (
	"context"
	"devops/common/typex"
	"devops/user/models"
	"gorm.io/gorm"

	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRoleMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleMenuLogic {
	return &CreateRoleMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoleMenuLogic) CreateRoleMenu(req *types.AddRoleMenuRequest) error {
	role := models.Role{}
	role.ID = req.RoleID
	if err := role.OneByID(); err != nil {
		return err
	}

	menu := models.Menu{}
	menus, _, _ := menu.All(nil, func(db *gorm.DB) *gorm.DB {
		return db.Where("id in ?", req.MenuIds)
	})

	//在进行创建之前，首先删除此角色的所有的菜单
	m := models.RoleMenu{}
	_ = m.Delete(l.ctx, map[string]int64{"role_id": req.RoleID})
	l.svcCtx.Orm.Exec("delete from casbin_rule where v0 = ?", role.Keyword)

	//增加新的菜单权限
	var list []models.RoleMenu
	for _, item := range menus {
		list = append(list, models.RoleMenu{
			RoleID: req.RoleID,
			MenuID: item.ID,
		})
		if item.Type == "A" && item.Permission != typex.BaseApiKey {
			l.svcCtx.Rbac.AddPolicy(role.Keyword, item.Path, item.Method)
		}
	}
	var object models.RoleMenus = list
	return object.Create(l.ctx)
}
