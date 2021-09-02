package engagementrate

import (
	"encoding/json"
	"fmt"
	"instagram/pkg/model"
	"instagram/pkg/param"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-rod/rod"
)

func EngagementRate(c *gin.Context) {
	p := param.Eject(c)
	defer func() {
		if err := recover(); err != nil {
			p.LogrusEntry.Warn("retry again:net::ERR_NETWORK_CHANGED")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "retry again,",
			})
			return
		}
	}()
	userName := c.Param("name")
	if userName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username is missing",
			"code":    http.StatusBadRequest,
		})
	}
	url := fmt.Sprintf("%s/%s", p.URL, userName)
	page := rod.New().MustConnect().MustPage(url).MustWaitLoad()
	if page == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "retry again,",
		})
	}
	page.MustScreenshot("me.png")
	scripts := page.MustElements("script")
	var windowSharedData string
	for _, script := range scripts {
		if strings.Contains(script.MustText(), "window._sharedData =") {
			windowSharedData = script.MustEval(`window._sharedData`).JSON("", "  ")
		}
	}
	var instagram model.Instagram
	if err := json.Unmarshal([]byte(windowSharedData), &instagram); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "unable to parse json",
			"code":    http.StatusUnprocessableEntity,
			"error":   err,
		})
		return
	}
	var commentsCount int
	var likesCount int
	var follower int
	for _, data := range instagram.EntryData.ProfilePage {
		follower = data.Graphql.User.EdgeFollowedBy.Count
		for _, edge := range data.Graphql.User.EdgeOwnerToTimelineMedia.Edges {
			commentsCount += edge.Node.EdgeMediaToComment.Count
			likesCount += edge.Node.EdgeLikedBy.Count
		}
	}
	likeCommentTotal := (likesCount + commentsCount)
	percentage := float64(likeCommentTotal) / float64(follower)
	c.JSON(http.StatusOK, gin.H{
		"message": "engagement rate caculated successfully",
		"code":    http.StatusOK,
		"result": model.EngagementRate{
			UserName:   userName,
			Percentage: percentage,
		},
	})
}
