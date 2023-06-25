package forumRepository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	forumLiterals "github.com/masharpik/ForumVKEducation/app/forums/utils/literals"
	forumStructs "github.com/masharpik/ForumVKEducation/app/forums/utils/structs"
	userStructs "github.com/masharpik/ForumVKEducation/app/users/utils/structs"
	mainLiterals "github.com/masharpik/ForumVKEducation/utils/literals"
)

func (repo Repository) CreateForum(forum forumStructs.Forum) (createdForum forumStructs.Forum, err error) {
	createForumQuery := `
		INSERT INTO forums (slug, title, author) 
		VALUES ($1, $2, $3) RETURNING slug, title, author;
	`

	err = repo.conn.QueryRow(context.Background(),
		createForumQuery,
		forum.Slug,
		forum.Title,
		forum.User).
		Scan(
			&createdForum.Slug,
			&createdForum.Title,
			&createdForum.User,
		)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case mainLiterals.CodeUniqConflict:
				err = errors.New(forumLiterals.ForumAlreadyExists)
			case mainLiterals.CodeForeignKeyNotFound:
				err = errors.New(forumLiterals.UserFKNotFound)
			}
		}

		return
	}

	return
}

func (repo Repository) GetForumInfoBySlug(slug string) (forum forumStructs.Forum, err error) {
	getForumQuery := `
		SELECT * 
		FROM forums 
		WHERE slug = $1;
	`

	err = repo.conn.QueryRow(context.Background(), getForumQuery,
		slug).
		Scan(
			&forum.Slug,
			&forum.Title,
			&forum.User,
			&forum.Threads,
			&forum.Posts,
		)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.New(forumLiterals.ForumNotFound)
		}

		return
	}

	return
}

func (repo Repository) GetUsers(slug, sinceNickname string, limit int, desc bool) (users userStructs.Users, err error) {
	users = make(userStructs.Users, 0)

	oper := ">"
	order := "ASC"
	if desc {
		order = "DESC"
		oper = "<"
	}

	var userRows pgx.Rows

	if sinceNickname == "" {
		getUsersQuery := fmt.Sprintf(`
SELECT nickname, fullname, about, email 
FROM user_forums 
WHERE forum = $1
ORDER BY nickname %s 
LIMIT $2;`, order)
		userRows, err = repo.conn.Query(context.Background(), getUsersQuery,
			slug,
			limit)
	} else {
		getUsersQuery := fmt.Sprintf(`
SELECT nickname, fullname, about, email 
FROM user_forums 
WHERE forum = $1 AND nickname %s $2 
ORDER BY nickname %s 
LIMIT $3;`, oper, order)
		userRows, err = repo.conn.Query(context.Background(), getUsersQuery,
			slug,
			sinceNickname,
			limit)
	}

	if err != nil {
		return
	}

	for userRows.Next() {
		var user userStructs.User

		err = userRows.Scan(&user.Nickname, &user.Fullname, &user.About, &user.Email)
		if err != nil {
			return
		}

		users = append(users, user)
	}

	return
}

func (repo Repository) CheckExistenceForum(slug string) (exists bool, err error) {
	checkExistenceForumQuery := `
		SELECT EXISTS(SELECT 1 FROM forums WHERE slug = $1) AS exists;
	`

	err = repo.conn.QueryRow(context.Background(), checkExistenceForumQuery, slug).
		Scan(&exists)
	if err != nil {
		return
	}

	return
}

func (repo Repository) GetForumSlugBySlug(slug string) (originSlug string, err error) {
	checkExistenceForumQuery := `
		SELECT slug FROM forums WHERE slug = $1;
	`

	err = repo.conn.QueryRow(context.Background(), checkExistenceForumQuery, slug).
		Scan(&originSlug)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.New(forumLiterals.ForumNotFound)
		}

		return
	}

	return
}
