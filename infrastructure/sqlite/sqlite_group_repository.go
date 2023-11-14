package sqlite

import (
	"go-to-school/main/domain/group"
)

type GroupRepository struct {
}

func NewGroupRepository() *GroupRepository {
	return &GroupRepository{}
}

func (repo *GroupRepository) GetGroups() *[]group.Group {
	list := []group.Group{}
	return &list

}

func (repo *GroupRepository) CreateGroup(req *group.CreateGroupRequest) group.Group {

}

func (repo *GroupRepository) GetById(ID string) (*group.Group, error) {}
