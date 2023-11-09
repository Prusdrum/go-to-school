package entrypoints

import (
	"encoding/json"
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
	jsonRes, err := json.Marshal(groups)
	if err != nil {
		http.Error(res, "parsing error", http.StatusInternalServerError)
	}

	res.Write(jsonRes)
}
