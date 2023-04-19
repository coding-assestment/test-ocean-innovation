package handlers

import (
	"fmt"
	"strconv"
	"sync"
	"test-ocean-innovation/models"
	"test-ocean-innovation/responseMessage"

	"github.com/labstack/echo/v4"
)

var (
	TeamsRepositories   = []*models.Team{}
	MembersRepositories = []*models.TeamMember{}
	seq                 = 0
	seqMember           = 0
	lock                = sync.Mutex{}
)

func filterMembers(data []*models.TeamMember, f func(*models.TeamMember) bool) []*models.TeamMember {

	fltd := make([]*models.TeamMember, 0)

	for _, e := range data {

		if f(e) {
			fltd = append(fltd, e)
		}
	}

	return fltd
}

func filterTeams(data []*models.Team, f func(*models.Team) bool) []*models.Team {

	fltd := make([]*models.Team, 0)

	for _, e := range data {

		if f(e) {
			fltd = append(fltd, e)
		}
	}

	return fltd
}

func RemoveIndex(slicep []*models.Team, i int) []*models.Team {
	TeamsRepositories = append(slicep[:i], slicep[i+1:]...)
	return TeamsRepositories
}

func CreateTeam(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := &models.Team{
		ID: seq,
	}
	members := []*models.TeamMember{}
	u.TeamMembers = members

	if err := c.Bind(u); err != nil {
		return responseMessage.Error(c, nil, err)
	}
	TeamsRepositories = append(TeamsRepositories, u)
	seq++

	return responseMessage.NewSuccess(c, u, "Success created")
}

func GetTeam(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, _ := strconv.Atoi(c.Param("id"))

	selectedItems := models.Team{}
	hasData := false
	for i := 0; i < len(TeamsRepositories); i++ {
		if TeamsRepositories[i].ID == id {
			hasData = true
			selectedItems = *TeamsRepositories[i]
		}
	}
	if !hasData {
		return responseMessage.NewNotFound(c, selectedItems, fmt.Sprintf("Item %v not found", id))
	}

	return responseMessage.NewSuccess(c, selectedItems, fmt.Sprintf("Success get item %v", id))
}

func UpdateTeam(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := new(models.Team)
	if err := c.Bind(u); err != nil {
		return responseMessage.Error(c, nil, err)
	}
	id, _ := strconv.Atoi(c.Param("id"))

	/* TeamsRepositories[id].Name = u.Name */
	selectedItems := models.Team{}
	hasData := false
	for i := 0; i < len(TeamsRepositories); i++ {
		if TeamsRepositories[i].ID == id {
			hasData = true
			TeamsRepositories[i].Name = u.Name
			selectedItems = *TeamsRepositories[i]
		}
	}
	if !hasData {
		return responseMessage.ListSuccess(c, selectedItems, fmt.Sprintf("Failed Updated item %v", id))
	}

	return responseMessage.NewSuccess(c, selectedItems, fmt.Sprintf("Updated item %v", id))
}

func DeleteTeam(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))

	deletedItem := filterTeams(TeamsRepositories, func(team *models.Team) bool {
		return team.ID == id
	})
	if len(deletedItem) == 0 {
		return responseMessage.NewSuccess(c, deletedItem, fmt.Sprintf("Delete failed, Item %v not found", id))
	}
	TeamsRepositories = filterTeams(TeamsRepositories, func(team *models.Team) bool {
		return team.ID != id
	})

	return responseMessage.NewSuccess(c, deletedItem, fmt.Sprintf("Deleted item %v", id))
}

func GetAllTeams(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	for i := 0; i < len(TeamsRepositories); i++ {
		members := filterMembers(MembersRepositories, func(member *models.TeamMember) bool {
			return member.TeamId == TeamsRepositories[i].ID
		})
		TeamsRepositories[i].TeamMembers = members
	}

	return responseMessage.ListSuccess(c, TeamsRepositories, "Success")
}
