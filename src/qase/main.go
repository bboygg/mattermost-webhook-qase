package qase

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bboygg/mattermost-webhook-qase/src/request"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func ReceiveWebhook(c *gin.Context) {
	channel := c.Param("channel")
	var body BaseQaseTriggeredPayload
	if err := c.BindJSON(&body); err != nil {
		println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
		return
	}

	raw, err := json.Marshal(body.Payload)
	if err != nil {
		println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
		return
	}

	//Display name for Test Runner
	var tester string
	var memberId int8 = body.TeamMemberID // You requires to add your team Member from

	switch memberId {
	case 1:
		tester = "QA_1"
	case 2:
		tester = "QA_2"
	case 3:
		tester = "QA_3"
	default:
		tester = "QA"
	}

	// Display project name by ProjectCode
	var projectName string = body.ProjectCode
	switch projectName {
	case "PROJECT_1_CODE":
		projectName = "PROJECT_1_NAME"
	case "PROJECT_2_CODE":
		projectName = "PROJECT_2_NAME"
	default:
	}

	if body.EventName == "run.started" {
		var payload RunTestPayload

		if err := json.Unmarshal(raw, &payload); err != nil {
			println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
			return
		}
		// mapstructure.Decode(body.Payload, &payload)
		request.Qase(channel, gin.H{
			"attachments": []gin.H{
				{
					"title": fmt.Sprintf("%s Test Run Started by %s", projectName, tester),
					"text":  fmt.Sprintf("[%s](https://app.qase.io/run/%s/dashboard/%d)", payload.Title, body.ProjectCode, payload.ID),
					"fields": []gin.H{
						{
							"short": true,
							"title": "Cases Count",
							"value": payload.CasesCount,
						},
						{
							"short": true,
							"title": "Description",
							"value": payload.Description,
						},
						{
							"short": true,
							"title": "Environment",
							"value": payload.Environment,
						},
					},
				},
			},
		})
	} else if body.EventName == "run.completed" {
		var payload CompleteTestPayload
		mapstructure.Decode(body.Payload, &payload)

		//Convert ms to Human readable format
		var ms = payload.Duration / 1000
		var sec = ms % 60
		var min = (ms / 60) % 60
		var hr = (ms / 60 / 60) % 24
		var duration = strconv.Itoa(hr) + "h " + strconv.Itoa(min) + "m " + strconv.Itoa(sec) + "s"

		request.Qase(channel, gin.H{
			"attachments": []gin.H{
				{
					"title": fmt.Sprintf("%s Test Run Completed by %s", projectName, tester),
					"text":  fmt.Sprintf("[%s](https://app.qase.io/run/%s/dashboard/%d)", "See the Result", body.ProjectCode, payload.ID),
					"fields": []gin.H{
						{
							"short": true,
							"title": "cases",
							"value": payload.Cases,
						},
						{
							"short": true,
							"title": "failed",
							"value": payload.Failed,
						},
						{
							"short": true,
							"title": "passed",
							"value": payload.Passed,
						},
						{
							"short": true,
							"title": "blocked",
							"value": payload.Blocked,
						},
						{
							"short": true,
							"title": "duration",
							"value": duration,
						},
					},
				},
			},
		})

	}
	c.JSON(200, gin.H{})
}
