package handlers

import (
	"fmt"
	"strconv"
	"test-ocean-innovation/models"
	"test-ocean-innovation/responseMessage"

	"github.com/labstack/echo/v4"
)

func AddTeamPlayer(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	teamId, _ := strconv.Atoi(c.Param("id"))

	tm := &models.TeamMember{
		ID:     seqMember,
		TeamId: teamId,
	}

	if err := c.Bind(tm); err != nil {
		return responseMessage.Error(c, nil, err)
	}

	MembersRepositories = append(MembersRepositories, tm)

	teamMembers := filterMembers(MembersRepositories, func(member *models.TeamMember) bool {
		return member.TeamId == teamId
	})

	TeamsRepositories[teamId].TeamMembers = teamMembers

	seqMember++
	return responseMessage.NewSuccess(c, TeamsRepositories[teamId], "Success created")
}

func GetPlayer(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	u := new(models.Team)
	if err := c.Bind(u); err != nil {
		return responseMessage.Error(c, nil, err)
	}
	id, _ := strconv.Atoi(c.Param("id"))

	selectedItems := models.TeamMember{}
	hasData := false
	for i := 0; i < len(MembersRepositories); i++ {
		if MembersRepositories[i].ID == id {
			hasData = true
			selectedItems = *MembersRepositories[i]
		}
	}
	if !hasData {
		return responseMessage.NewNotFound(c, selectedItems, fmt.Sprintf("Item not found %v", id))
	}

	return responseMessage.NewSuccess(c, selectedItems, fmt.Sprintf("Success get item %v", id))

}
