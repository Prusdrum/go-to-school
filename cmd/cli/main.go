package main

import (
	"fmt"
	"go-to-school/main/app"
	"go-to-school/main/infrastructure/sqlite"
	groups "go-to-school/main/internal/groups/domain"
	"time"
)

func main() {
	groupRepository := sqlite.NewGroupRepository()
	handleCreateGroupRequest := app.NewCreateGroupHandler(groupRepository)
	// repo, err := inmemory.GroupRepository

	fmt.Println("Hello, go to school")
	testGroup := groups.Group{
		ID:        "dsadas",
		Name:      "dsdasdas",
		CreatedAt: time.Now(),
	}
	fmt.Println("Group created", testGroup.Name)

	i := 20
	pointer := &i
	fmt.Println(i)
	fmt.Println(*pointer)
}
