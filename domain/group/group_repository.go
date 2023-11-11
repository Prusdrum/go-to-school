package group

type CreateGroupRequest struct {
	Name string
}

type Repository interface {
	GetGroups() *[]Group
	CreateGroup(*CreateGroupRequest) *Group
}
