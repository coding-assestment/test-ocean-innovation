package models

type (
	Team struct {
		ID          int           `json:"id"`
		Name        string        `json:"name"`
		TeamMembers []*TeamMember `json:"team_members"`
	}

	TeamMember struct {
		ID     int    `json:"id"`
		TeamId int    `json:"team_id"`
		Name   string `json:"name"`
	}
)
