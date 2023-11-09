package main

import (
	"fmt"
	groups "go-to-school/main/internal/groups/domain"
	"time"
)

func main() {
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
