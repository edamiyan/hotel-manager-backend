package repository

import (
	"fmt"
	hotelManager "github.com/edamiyan/hotel-manager"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user hotelManager.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (hotelManager.User, error) {
	var user hotelManager.User
	query := fmt.Sprintf("SELECT id from %s WHERE username=$1 and password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
