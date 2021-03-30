package repository

import (
	"database/sql"
	"log"

	"github.com/dongil91/module-test/domain"
)

type SellerRepository struct {
	DB *sql.DB
}

func NewMysqlSellerRepository(db *sql.DB) domain.SellerRepository {
	return &SellerRepository{
		DB: db,
	}
}

func (u *SellerRepository) FindById(id int64) (*domain.Seller, error) {
	row := u.DB.QueryRow("SELECT id, name, email, last_modified_at, created_at FROM user WHERE id = ?", id)
	log.Println(row)
	var user domain.Seller
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

func (u *SellerRepository) FindAll() ([]*domain.Seller, error) {
	rows, err := u.DB.Query("SELECT id, name, email, last_modified_at, created_at FROM user")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	log.Println(rows)
	var users []*domain.Seller
	for rows.Next() {
		var user domain.Seller
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.LastModifiedAt, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (u *SellerRepository) Create(name string, email string) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec("INSERT INTO user(name, email) VALUE(?, ?)", name, email)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	log.Println(result)
	return nil
}

func (u *SellerRepository) Update(name string, id int64) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Exec("UPDATE user SET name = ? WHERE id = ?", name, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	log.Println(result)
	return nil
}

func (u *SellerRepository) Delete(id int) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return nil
	}

	result, err := tx.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	log.Println(result)
	return nil
}
