package handler

import (
	"github.com/limeschool/gin"
	"service/consts"
	"service/errors"
	"service/service"
	"service/types"
)

func AllCodeRegistries(ctx *gin.Context) {
	if resp, err := service.AllCodeRegistries(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func AllCodeRegistryFilter(ctx *gin.Context) {
	if resp, err := service.AllCodeRegistryFilter(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func AddCodeRegistry(ctx *gin.Context) {
	in := types.AddCodeRegistryRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddCodeRegistry(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateCodeRegistry(ctx *gin.Context) {
	in := types.UpdateCodeRegistryRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateCodeRegistry(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteCodeRegistry(ctx *gin.Context) {
	in := types.DeleteCodeRegistryRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.DeleteCodeRegistry(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func ConnectCodeRegistry(ctx *gin.Context) {
	in := types.ConnectCodeRegistryRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.ConnectCodeRegistry(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func AllCodeRegistryTypes(ctx *gin.Context) {
	ctx.RespData([]string{consts.GITEE, consts.GITLAB, consts.GITHUB})
}

func AllCodeRegistryCloneTypes(ctx *gin.Context) {
	ctx.RespData([]string{consts.SSHURL, consts.HTMLURL})
}

func GetCodeRegistryProject(ctx *gin.Context) {
	in := types.GetCodeRegistryProjectRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if data, err := service.GetCodeRegistryProject(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(data)
	}
}

func AllCodeRegistryBranches(ctx *gin.Context) {
	in := types.AllCodeRegistryBranchesRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if data, err := service.AllCodeRegistryBranches(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(data)
	}
}

func AllCodeRegistryTags(ctx *gin.Context) {
	in := types.AllCodeRegistryTagsRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if data, err := service.AllCodeRegistryTags(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(data)
	}
}
