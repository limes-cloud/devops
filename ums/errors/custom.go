package errors

import "github.com/limeschool/gin"

var (
	//基础相关
	ParamsError           = &gin.CustomError{Code: 100002, Msg: "参数验证失败"}
	AssignError           = &gin.CustomError{Code: 100003, Msg: "数据赋值失败"}
	DBError               = &gin.CustomError{Code: 100004, Msg: "数据库操作失败"}
	DBNotFoundError       = &gin.CustomError{Code: 100005, Msg: "未查询到指定数据"}
	UserNameNotFoundError = &gin.CustomError{Code: 100006, Msg: "账号不存在"}
	UserNameDisableError  = &gin.CustomError{Code: 100007, Msg: "账号已被禁用"}
	PasswordError         = &gin.CustomError{Code: 100008, Msg: "账号密码错误"}
	RsaPasswordError      = &gin.CustomError{Code: 100009, Msg: "非法账号密码"}
	IpLimitLoginError     = &gin.CustomError{Code: 100010, Msg: "当前设备登陆错误次数过多,今日已被限制登陆"}
	SuperAdminEditError   = &gin.CustomError{Code: 100011, Msg: "超级管理员不允许修改"}
	SuperAdminDelError    = &gin.CustomError{Code: 100012, Msg: "超级管理员不允许删除"}
	//auth相关
	NotResourcePower     = &gin.CustomError{Code: 4003, Msg: "暂无接口资源权限"}
	TokenExpiredError    = &gin.CustomError{Code: 4001, Msg: "登陆信息已过期，请重新登陆"}
	RefTokenExpiredError = &gin.CustomError{Code: 4000, Msg: "太长时间未登陆，请重新登陆"}
	DulDeviceLoginError  = &gin.CustomError{Code: 4000, Msg: "你已在其他设备登陆"}
	TokenValidateError   = &gin.CustomError{Code: 4000, Msg: "token验证失败"}
	TokenEmptyError      = &gin.CustomError{Code: 4000, Msg: "token信息不存在"}

	//menu相关
	DulMenuNameError  = &gin.CustomError{Code: 1000030, Msg: "菜单name值不能重复"}
	MenuParentIdError = &gin.CustomError{Code: 1000031, Msg: "父菜单不能为自己"}

	//team相关
	NotAddTeamError      = &gin.CustomError{Code: 1000040, Msg: "暂无此部门的下级部门创建权限"}
	NotEditTeamError     = &gin.CustomError{Code: 1000041, Msg: "暂无此部门的修改权限"}
	NotDelTeamError      = &gin.CustomError{Code: 1000042, Msg: "暂无此部门的删除权限"}
	NotAddTeamUserError  = &gin.CustomError{Code: 1000043, Msg: "暂无此部门的人员创建权限"}
	NotEditTeamUserError = &gin.CustomError{Code: 1000044, Msg: "暂无此部门的人员修改权限"}
	NotDelTeamUserError  = &gin.CustomError{Code: 1000045, Msg: "暂无此部门的人员删除权限"}
	TeamParentIdError    = &gin.CustomError{Code: 1000046, Msg: "父部门不能为自己"}

	//role相关
	DulKeywordError = &gin.CustomError{Code: 1000050, Msg: "角色标志符已存在"}
)
