package main

import (
	"database/sql"
	"fmt"
	"go-to-school/main/app"
	"go-to-school/main/infrastructure/sqlite"
	"log"
	"path/filepath"
)

func main() {
	path, err := filepath.Abs("database/go-to-school.sqlite3")
	if err != nil {
		log.Fatal("get db path: ", err)
	}
	db, err := sql.Open("sqlite3", path)
	groupRepository := sqlite.NewGroupRepository(db)
	handleCreateGroupRequest := app.NewCreateGroupHandler(groupRepository)

	createGroupRequest := &app.CreateGroupRequest{
		Name: "Test group Mateusz",
	}
	handleCreateGroupRequest.HandleCreateGroup(createGroupRequest)

	groups := handleCreateGroupRequest.HandleGetAllGroups()
	fmt.Println(groups)
}
