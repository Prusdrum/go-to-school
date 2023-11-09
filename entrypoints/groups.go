package entrypoints

import (
	groups "go-to-school/main/internal/groups/domain"
	"log"
	"net/http"
)

type Groups struct {
	logger *log.Logger
}

func NewGroups(logger *log.Logger) *Groups {
	return &Groups{logger}
}

func (entrypoint *Groups) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	groups := groups.GetGroups()

	err := groups.ToJSON(res)

	if err != nil {
		http.Error(res, "parsing error", http.StatusInternalServerError)
	}
}
