package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/dongil91/module-test/domain"
)

type UserRepository struct {
	DB *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) domain.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) FindById(ctx context.Context, id int64) (*domain.User, error) {
	row := u.DB.QueryRow("SELECT id, name, email, last_modified_at, created_at FROM user WHERE id = ?", id)
	log.Println(row)
	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.LastModifiedAt, &user.CreatedAt)
	switch err {
	case sql.ErrNoRows:
		return nil, sql.ErrNoRows
	case nil:
		return &user, nil
	default:
		panic(err)
	}
}

func (u *UserRepository) FindAll() ([]*domain.User, error) {
	rows, err := u.DB.Query("SELECT id, name, email, last_modified_at, created_at FROM user")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	log.Println(rows)
	var users []*domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.LastModifiedAt, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
