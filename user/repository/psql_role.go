package repository

import (
	"database/sql"
	"errors"

	"github.com/betsegawlemma/restaurant/entity"
)

// RoleRepositoryImpl implements RoleRepository interface
type RoleRepositoryImpl struct {
	conn *sql.DB
}

// NewRoleRepositoryImpl returns new RoleRepositoryImpl object
func NewRoleRepositoryImpl(Conn *sql.DB) *RoleRepositoryImpl {
	return &RoleRepositoryImpl{conn: Conn}
}

// StoreRole stores a given role object into a database
func (rr *RoleRepositoryImpl) StoreRole(role entity.Role) error {

	query := "INSERT INTO roles (name) values ($1)"
	_, err := rr.conn.Exec(query, role.Name)

	if err != nil {
		return errors.New("INSERT has failed")
	}

	return nil
}

func (rr *RoleRepositoryImpl) Roles() ([]entity.Role, error) {
	return []entity.Role{}, nil
}

func (rr *RoleRepositoryImpl) Role(id int) (entity.Role, error) {
	return entity.Role{}, nil
}

func (rr *RoleRepositoryImpl) UpdateRole(role entity.Role) error { return nil }

func (rr *RoleRepositoryImpl) DeleteRole(id int) error { return nil }
