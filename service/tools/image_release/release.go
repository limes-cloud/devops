package image_release

import (
	"github.com/limeschool/gin"
)

type ImageRelease interface {
	DeleteFormYaml(ctx *gin.Context, applyYaml string) (err error)
	UpdateFromYaml(ctx *gin.Context, applyYaml string) (err error)
}
