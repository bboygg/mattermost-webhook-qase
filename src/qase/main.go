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
	var memberId int8 = body.TeamMemberID // George, DH, Juni, Dunkin, Sam, heejung

	switch memberId {
	case 1:
		tester = "George"
	case 2:
		tester = "DH"
	case 3:
		tester = "Juni"
	case 4:
		tester = "Dunkin"
	case 5:
		tester = "Sam"
	case 6:
		tester = "Heejung"
	default:
		tester = "QA"
	}

	// Display project name by ProjectCode
	var projectName string = body.ProjectCode
	switch projectName {
	case "PLUS":
		projectName = "Pivo+"
	case "PIVOLIVE":
		projectName = "Pivo Live"
	case "PIVOTOUR":
		projectName = "Pivo Tour"
	case "PLAY":
		projectName = "Pivo Play"
	case "CAST":
		projectName = "Pivo Cast"
	case "PRESENT":
		projectName = "Pivo Present"
	case "STUDIO":
		projectName = "Pivo Studio"
	case "BP":
		projectName = "Beamo Portal"
	case "BA":
		projectName = "Beamo App"
	case "MF":
		projectName = "Pivo Meet Frontend"
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
