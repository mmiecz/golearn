package models

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	Id           int
	Email        string
	PasswordHash string
}

type NewUser struct {
	Email    string
	Password string // not hashed password yet
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(nu NewUser) (*User, error) {
	email := strings.ToLower(nu.Email)
	hash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash user password: %w", err)
	}
	passHash := string(hash)

	user := User{
		Email:        email,
		PasswordHash: passHash,
	}

	query := us.DB.QueryRow(
		`INSERT INTO users (email, password_hash)
			VALUES ($1, $2) RETURNING id`, email, passHash)
	err = query.Scan(&user.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}
	return &user, nil
}
