package app

import "go-to-school/main/domain/group"

type GetAllGroupsHandler struct {
	groupRepository group.Repository
}

func NewGetAllGroupsHandler(Repository group.Repository) *GetAllGroupsHandler {
	return &GetAllGroupsHandler{
		groupRepository: Repository,
	}
}

func (handler *CreateGroupHandler) HandleGetAllGroups() *[]group.Group {
	return handler.groupRepository.GetGroups()
}
