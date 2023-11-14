package group

type CreateGroupRequest struct {
	Name string
}

type Repository interface {
	GetGroups() *[]Group
	GetById(ID string) (*Group, error)
	CreateGroup(*CreateGroupRequest) *Group
}
