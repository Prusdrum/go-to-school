package inmemory

import (
	"go-to-school/main/domain/group"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
)

type GroupRepository struct {
	db *memdb.MemDB
}

func (repo *GroupRepository) New() (*GroupRepository, error) {
	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"groups": {
				Name: "group",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": {
						Name:    "age",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
					"created_at": {
						Name:    "age",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "CreatedAt"},
					},
				},
			},
		},
	}
	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}
	return &GroupRepository{db: db}, nil
}

func (repo *GroupRepository) GetGroups() *[]group.Group {
	list := []group.Group{}
	return &list

}

func (repo *GroupRepository) CreateGroup(req *group.CreateGroupRequest) group.Group {
	newGroup := &group.Group{
		ID:        uuid.New().String(),
		Name:      req.Name,
		CreatedAt: time.Now(),
	}
	txn := repo.db.Txn(true)

	err := txn.Insert("groups", newGroup)
	if err != nil {
		panic(err)
	}

	txn.Commit()
	return *newGroup
}
