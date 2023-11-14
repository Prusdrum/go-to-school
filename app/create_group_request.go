package app

import "go-to-school/main/domain/group"

type CreateGroupRequest struct {
	Name string
}

type CreateGroupHandler struct {
	groupRepository group.Repository
}

func NewCreateGroupHandler(Repository group.Repository) *CreateGroupHandler {
	return &CreateGroupHandler{
		groupRepository: Repository,
	}
}

func (handler *CreateGroupHandler) HandleCreateGroup(createGroupRequest *CreateGroupRequest) {
	newGroup := group.CreateGroupRequest{
		Name: createGroupRequest.Name,
	}
	handler.groupRepository.CreateGroup(&newGroup)
}

func (handler *CreateGroupHandler) HandleGetAllGroups() *[]group.Group {
	return handler.groupRepository.GetGroups()
}
