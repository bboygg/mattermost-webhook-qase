package qase

type BaseQaseTriggeredPayload struct {
	EventName    string                 `json:"event_name"`
	Timestamp    int32                  `json:"timestamp"`
	ProjectCode  string                 `json:"project_code"`
	TeamMemberID int8                   `json:"team_member_id"`
	Payload      map[string]interface{} `json:"payload"`
}

type RunTestPayload struct {
	ID   int `json:"id"`
	Plan struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}
	Title       string `json:"title"`
	CasesCount  int    `json:"cases_count"`
	Description string `json:"description"`
	Environment string `json:"environment"`
}

type QaseTriggeredRunTestPayload struct {
	EventName    string         `json:"event_name"`
	Timestamp    int32          `json:"timestamp"`
	ProjectCode  string         `json:"project_code"`
	TeamMemberID int8           `json:"team_member_id"`
	Payload      RunTestPayload `json:"payload"`
}

type CompleteTestPayload struct {
	ID       int   `json:"id"`
	Cases    uint8 `json:"cases"`
	Failed   uint8 `json:"failed"`
	Passed   uint8 `json:"passed"`
	Blocked  uint8 `json:"blocked"`
	Duration int   `json:"duration"`
}
