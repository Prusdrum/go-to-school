package main

import (
	"fmt"
	"go-to-school/main/app"
	"go-to-school/main/infrastructure/sqlite"
)

func main() {
	groupRepository := sqlite.NewGroupRepository()
	handleCreateGroupRequest := app.NewCreateGroupHandler(groupRepository)
	// repo, err := inmemory.GroupRepository

	createGroupRequest := &app.CreateGroupRequest{
		Name: "Test group Mateusz",
	}
	handleCreateGroupRequest.HandleCreateGroup(createGroupRequest)

	groups := handleCreateGroupRequest.HandleGetAllGroups()
	fmt.Println(groups)
}
