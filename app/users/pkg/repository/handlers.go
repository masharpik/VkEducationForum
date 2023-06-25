package userRepository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	userLiterals "github.com/masharpik/ForumVKEducation/app/users/utils/literals"
	userStructs "github.com/masharpik/ForumVKEducation/app/users/utils/structs"
	mainLiterals "github.com/masharpik/ForumVKEducation/utils/literals"
)

func (repo Repository) CreateUser(user userStructs.User) (createdUser userStructs.User, err error) {
	createUserQuery := `
		INSERT INTO users (nickname, fullname, about, email)
		VALUES ($1, $2, $3, $4)
		RETURNING *;
	`
	err = repo.conn.QueryRow(context.Background(), createUserQuery,
		user.Nickname,
		user.Fullname,
		user.About,
		user.Email).
		Scan(&createdUser.Nickname,
			&createdUser.Fullname,
			&createdUser.About,
			&createdUser.Email)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case mainLiterals.CodeUniqConflict:
				err = fmt.Errorf(userLiterals.LogUserAlreadyExists)
			}
		}

		return
	}

	return
}

func (repo Repository) GetUsersByNickOrMail(nickname, email string) (createdUsers userStructs.Users, err error) {
	getUsersQuery := `
		SELECT * FROM users 
		WHERE nickname = $1 OR email = $2;
	`

	var userRows pgx.Rows
	userRows, err = repo.conn.Query(context.Background(), getUsersQuery,
		nickname,
		email)
	if err != nil {
		return
	}

	for userRows.Next() {
		var user userStructs.User

		err = userRows.Scan(&user.Nickname, &user.Fullname, &user.About, &user.Email)
		if err != nil {
			return
		}

		createdUsers = append(createdUsers, user)
	}

	return
}

func (repo Repository) GetNickByNick(nickname string) (nick string, err error) {
	getUsersQuery := `
		SELECT nickname FROM users 
		WHERE nickname = $1;
	`

	err = repo.conn.QueryRow(context.Background(), getUsersQuery,
		nickname).Scan(&nick)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.New(userLiterals.UserNotFound)
		}

		return
	}

	return
}

func (repo Repository) GetUserInfo(nickname string) (user userStructs.User, err error) {
	getUserQuery := `
		SELECT * FROM users WHERE nickname = $1;
	`

	err = repo.conn.QueryRow(context.Background(), getUserQuery, nickname).
		Scan(&user.Nickname,
			&user.Fullname,
			&user.About,
			&user.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.New(userLiterals.UserNotFound)
		}

		return
	}

	return
}

func (repo Repository) UpdateUserInfo(user userStructs.User) (err error) {
	updateUserQuery := `
		UPDATE users 
		SET fullname = $1, about = $2, email = $3
		WHERE nickname = $4;
	`

	_, err = repo.conn.Exec(context.Background(), updateUserQuery,
		user.Fullname,
		user.About,
		user.Email,
		user.Nickname)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case mainLiterals.CodeUniqConflict:
				err = fmt.Errorf(userLiterals.LogUserAlreadyExists)
			}
		}

		return
	}

	return
}
