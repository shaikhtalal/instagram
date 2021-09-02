package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AbortWithError(
	g *gin.Context,
	status int,
	code string,
	message string,
	err error,
) {
	data := gin.H{
		"code":    code,
		"status":  status,
		"success": false,
	}
	if err != nil {
		message = fmt.Sprintf("%s: %s", message, err)
	}
	data["message"] = message

	g.AbortWithStatusJSON(status, data)
}
