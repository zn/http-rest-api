package sqlstore

import (
	"database/sql"

	"github.com/zn/http-rest-api/internal/app/model"
	"github.com/zn/http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users(email, encrypted_password) VALUES($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	err := r.store.db.QueryRow("SELECT * FROM users WHERE email = $1", email).
		Scan(&u.ID, &u.Email, &u.EncryptedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	err := r.store.db.QueryRow("SELECT * FROM users WHERE id = $1", id).
		Scan(&u.ID, &u.Email, &u.EncryptedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}
