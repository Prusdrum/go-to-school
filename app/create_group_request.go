package app

import "go-to-school/main/domain/group"

type CreateGroupRequest struct {
	Name string
}

type CreateGroupHandler struct {
	GroupRepository group.Repository
}

func (handler *CreateGroupHandler) New(Repository group.Repository) *CreateGroupHandler {
	return &CreateGroupHandler{
		GroupRepository: Repository,
	}
}

func (handler *CreateGroupHandler) HandleCreateGroup(createGroupRequest *CreateGroupRequest) {
	newGroup := group.CreateGroupRequest{
		Name: createGroupRequest.Name,
	}
	handler.GroupRepository.CreateGroup(&newGroup)
}
