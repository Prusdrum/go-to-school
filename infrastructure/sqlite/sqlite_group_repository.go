package sqlite

import (
	"database/sql"
	"go-to-school/main/domain/group"
	"log"
	"time"

	uuid "github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type GroupRepository struct {
	db *sql.DB
}

func NewGroupRepository() *GroupRepository {
	db, err := sql.Open("sqlite3", "../../database/go-to-school-db.db")
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("CREATE TABLE IF NOT EXISTS groups (id TEXT, name TEXT, created_at INTEGER)")

	return &GroupRepository{
		db: db,
	}
}

func (repo *GroupRepository) GetGroups() *[]group.Group {
	stmt, err := repo.db.Prepare("SELECT id, name, created_at FROM groups")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	groups := []group.Group{}

	for rows.Next() {
		var id string
		var name string
		var createdAt int

		err = rows.Scan(&id, &name, &createdAt)
		if err != nil {
			log.Fatal(err)
		}
		groups = append(groups, group.Group{
			ID:        id,
			Name:      name,
			CreatedAt: time.Unix(int64(createdAt), 0),
		})
	}
	return &groups
}

func (repo *GroupRepository) CreateGroup(req *group.CreateGroupRequest) *group.Group {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()
	stmt, err := tx.Prepare("insert into groups (id, name, created_at) values(?, ?, ?)")
	defer stmt.Close()
	group := group.Group{
		ID:        uuid.New().String(),
		Name:      req.Name,
		CreatedAt: time.Now(),
	}
	_, err = stmt.Exec(group.ID, group.Name, group.CreatedAt.Unix())
	if err != nil {
		log.Fatal(err)
	}

	return &group
}

func (repo *GroupRepository) GetById(ID string) (*group.Group, error) {
	return nil, nil
}
