package router

import (
	"instagram/pkg/engagementrate"

	"github.com/gin-gonic/gin"
)

func HandleHTTP(e *gin.Engine) {
	e.GET("/engagementRate/:name", engagementrate.EngagementRate)
}
