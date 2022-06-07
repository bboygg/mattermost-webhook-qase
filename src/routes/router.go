package routes

import (
	"github.com/bboygg/mattermost-webhook-qase/src/qase"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup) {
	InitQase(r.Group("/qase"))
}

func InitQase(r *gin.RouterGroup) {
	r.POST("/:channel", qase.ReceiveWebhook)
}
