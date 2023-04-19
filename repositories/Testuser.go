package repositories

type (
	TestUser struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	Team struct {
		ID          int          `json:"id"`
		Name        string       `json:"name"`
		TeamMembers []TeamMember `json:"team_members"`
	}

	TeamMember struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)
