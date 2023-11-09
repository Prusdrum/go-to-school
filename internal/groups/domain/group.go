package groups

import (
	"time"
)

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// omit from output
	CreatedAt time.Time `json:"-"`
}

func GetGroups() []*Group {
	return testGroups
}

var testGroups = []*Group{
	{
		ID:        "test-id-1",
		Name:      "Group 1",
		CreatedAt: time.Date(2023, time.August, 25, 13, 23, 0, 0, time.UTC),
	},
	{
		ID:        "test-id-2",
		Name:      "Group 2",
		CreatedAt: time.Date(2023, time.August, 25, 13, 23, 0, 0, time.UTC),
	},
}
