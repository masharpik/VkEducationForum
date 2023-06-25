package threadRepository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	forumLiterals "github.com/masharpik/ForumVKEducation/app/forums/utils/literals"
	threadLiterals "github.com/masharpik/ForumVKEducation/app/threads/utils/literals"
	threadStructs "github.com/masharpik/ForumVKEducation/app/threads/utils/structs"
	mainLiterals "github.com/masharpik/ForumVKEducation/utils/literals"
)

func (repo Repository) CreateThread(thread threadStructs.Thread) (threadStructs.Thread, error) {
	var err error

	if thread.Slug == "" {
		createThreadQuery := `
			INSERT INTO threads (title, author, forum, message, created) 
			VALUES ($1, $2, $3, $4, $5) RETURNING id;
		`
		err = repo.conn.QueryRow(context.Background(), createThreadQuery,
			thread.Title,
			thread.Author,
			thread.Forum,
			thread.Message,
			thread.Created).
			Scan(&thread.Id)
	} else {
		createThreadQuery := `
			INSERT INTO threads (slug, title, author, forum, message, created) 
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;
		`

		err = repo.conn.QueryRow(context.Background(), createThreadQuery,
			thread.Slug,
			thread.Title,
			thread.Author,
			thread.Forum,
			thread.Message,
			thread.Created).
			Scan(&thread.Id)
	}

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case mainLiterals.CodeUniqConflict:
				err = errors.New(threadLiterals.ThreadAlreadyExists)
			}
		}

		return threadStructs.Thread{}, err
	}

	return thread, nil
}

func (repo Repository) GetThreadInfoBySlug(slug string) (thread threadStructs.Thread, err error) {
	getThreadBySlugQuery := `
		SELECT id, COALESCE(slug, '') AS slug, title, author, forum, message, votes, created FROM threads WHERE slug = $1;
	`

	var t time.Time
	err = repo.conn.QueryRow(context.Background(), getThreadBySlugQuery,
		slug).
		Scan(
			&thread.Id,
			&thread.Slug,
			&thread.Title,
			&thread.Author,
			&thread.Forum,
			&thread.Message,
			&thread.Votes,
			&t,
		)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.New(forumLiterals.ForumNotFound)
		}

		return
	}
	thread.Created = t.UTC().Format("2006-01-02T15:04:05.000Z")

	return
}

func (repo Repository) GetThreadInfoById(id int) (thread threadStructs.Thread, err error) {
	getThreadBySlugQuery := `
		SELECT id, COALESCE(slug, '') AS slug, title, author, forum, message, votes, created FROM threads WHERE id = $1;
	`

	var t time.Time
	err = repo.conn.QueryRow(context.Background(), getThreadBySlugQuery,
		id).
		Scan(
			&thread.Id,
			&thread.Slug,
			&thread.Title,
			&thread.Author,
			&thread.Forum,
			&thread.Message,
			&thread.Votes,
			&t,
		)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.New(forumLiterals.ForumNotFound)
		}

		return
	}
	thread.Created = t.UTC().Format("2006-01-02T15:04:05.000Z")

	return
}

func (repo Repository) GetThreadsByForum(slug, sinceDate string, limit int, desc bool) (threads threadStructs.Threads, err error) {
	var threadRows pgx.Rows
	threads = make(threadStructs.Threads, 0)
	if sinceDate != "" {
		oper := ">="
		order := "ASC"
		if desc {
			order = "DESC"
			oper = "<="
		}

		getThreadsQuery := fmt.Sprintf(`
			SELECT id, COALESCE(slug, '') AS slug, title, author, forum, message, votes, created 
			FROM threads 
			WHERE forum = $1 AND created %s $2 
			ORDER BY created %s 
			LIMIT $3;
		`, oper, order)

		threadRows, err = repo.conn.Query(context.Background(), getThreadsQuery,
			slug,
			sinceDate,
			limit)
	} else {
		order := "ASC"
		if desc {
			order = "DESC"
		}

		getThreadsQuery := fmt.Sprintf(`
			SELECT id, COALESCE(slug, '') AS slug, title, author, forum, message, votes, created 
			FROM threads 
			WHERE forum = $1
			ORDER BY created %s 
			LIMIT $2;
		`, order)

		threadRows, err = repo.conn.Query(context.Background(), getThreadsQuery,
			slug,
			limit)
	}

	if err != nil {
		return
	}

	for threadRows.Next() {
		var thread threadStructs.Thread

		var t time.Time
		err = threadRows.Scan(
			&thread.Id,
			&thread.Slug,
			&thread.Title,
			&thread.Author,
			&thread.Forum,
			&thread.Message,
			&thread.Votes,
			&t,
		)
		if err != nil {
			return
		}

		thread.Created = t.UTC().Format("2006-01-02T15:04:05.000Z")

		threads = append(threads, thread)
	}

	return
}

func (repo Repository) GetForumSlugByThreadId(threadId int) (slug string, err error) {
	getForumQuery := `
		SELECT forum FROM threads WHERE id = $1; 
	`

	err = repo.conn.QueryRow(context.Background(), getForumQuery, threadId).Scan(&slug)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.New(threadLiterals.ThreadNotFound)
		}

		return
	}

	return
}

func (repo Repository) GetThreadIdBySlug(slug string) (id int, err error) {
	getIdQuery := `
		SELECT id 
		FROM threads 
		WHERE slug = $1;
	`

	err = repo.conn.QueryRow(context.Background(), getIdQuery, slug).
		Scan(&id)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.New(threadLiterals.ThreadNotFound)
		}

		return
	}

	return
}

func (repo Repository) UpdateThreadBySlug(slug string, title, message string) (thread threadStructs.Thread, err error) {
	var updateMsgQuery string

	var titleNull bool = (title == "")
	var messageNull bool = (message == "")
	values := make([]interface{}, 0)
	if !titleNull && !messageNull {
		updateMsgQuery = "UPDATE threads SET title = $1, message = $2  WHERE slug = $3 RETURNING id, COALESCE(slug, '') AS slug, title, author, forum, message, votes, created;"
		values = append(values, title, message, slug)
	} else if !titleNull {
		updateMsgQuery = "UPDATE threads SET title = $1 WHERE slug = $2 RETURNING id, COALESCE(slug, '') AS slug, title, author, forum, message, votes, created;"
		values = append(values, title, slug)
	} else if !messageNull {
		updateMsgQuery = "UPDATE threads SET message = $1 WHERE slug = $2 RETURNING id, COALESCE(slug, '') AS slug, title, author, forum, message, votes, created;"
		values = append(values, message, slug)
	} else {
		updateMsgQuery = "SELECT id, COALESCE(slug, '') AS slug, title, author, forum, message, votes, created FROM threads WHERE slug = $1;"
		values = append(values, slug)
	}

	var t time.Time
	err = repo.conn.QueryRow(context.Background(), updateMsgQuery, values...).
		Scan(
			&thread.Id,
			&thread.Slug,
			&thread.Title,
			&thread.Author,
			&thread.Forum,
			&thread.Message,
			&thread.Votes,
			&t,
		)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.New(threadLiterals.ThreadNotFound)
		}

		return
	}
	thread.Created = t.UTC().Format("2006-01-02T15:04:05.000Z")

	return
}

func (repo Repository) UpdateThreadById(id int, title, message string) (thread threadStructs.Thread, err error) {
	var updateMsgQuery string

	var titleNull bool = (title == "")
	var messageNull bool = (message == "")
	values := make([]interface{}, 0)
	if !titleNull && !messageNull {
		updateMsgQuery = "UPDATE threads SET title = $1, message = $2  WHERE id = $3 RETURNING id, COALESCE(slug, '') AS slug, title, author, forum, message, votes, created;"
		values = append(values, title, message, id)
	} else if !titleNull {
		updateMsgQuery = "UPDATE threads SET title = $1  WHERE id = $2 RETURNING id, COALESCE(slug, '') AS slug, title, author, forum, message, votes, created;"
		values = append(values, title, id)
	} else if !messageNull {
		updateMsgQuery = "UPDATE threads SET message = $1  WHERE id = $2 RETURNING id, COALESCE(slug, '') AS slug, title, author, forum, message, votes, created;"
		values = append(values, message, id)
	} else {
		updateMsgQuery = "SELECT id, COALESCE(slug, '') AS slug, title, author, forum, message, votes, created FROM threads WHERE id = $1;"
		values = append(values, id)
	}

	var t time.Time
	err = repo.conn.QueryRow(context.Background(), updateMsgQuery, values...).
		Scan(
			&thread.Id,
			&thread.Slug,
			&thread.Title,
			&thread.Author,
			&thread.Forum,
			&thread.Message,
			&thread.Votes,
			&t,
		)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.New(threadLiterals.ThreadNotFound)
		}

		return
	}
	thread.Created = t.UTC().Format("2006-01-02T15:04:05.000Z")

	return
}

func (repo Repository) VoteByThreadId(id int, nickname string, vote int) (thread threadStructs.Thread, err error) {
	insertVoteQuery := `
		INSERT INTO votes (thread, nickname, vote) 
		VALUES ($1, $2, $3) 
		ON CONFLICT (thread, nickname) 
		DO UPDATE SET vote = $3;
	`

	getThreadQuery := `
		SELECT id, COALESCE(slug, '') AS slug, title, author, forum, message, votes, created FROM threads WHERE id = $1;
	`

	tx, err := repo.conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return
	}

	_, err = tx.Exec(context.Background(), insertVoteQuery, id, nickname, vote)
	if err != nil {
		tx.Rollback(context.Background())

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case mainLiterals.CodeForeignKeyNotFound:
				err = errors.New(threadLiterals.ForumHasNotThread)
			}
		}
		return
	}

	var t time.Time
	err = tx.QueryRow(context.Background(), getThreadQuery, id).
		Scan(
			&thread.Id,
			&thread.Slug,
			&thread.Title,
			&thread.Author,
			&thread.Forum,
			&thread.Message,
			&thread.Votes,
			&t,
		)
	if err != nil {
		tx.Rollback(context.Background())
		return
	}
	thread.Created = t.UTC().Format("2006-01-02T15:04:05.000Z")

	err = tx.Commit(context.Background())
	if err != nil {
		return
	}

	return
}

func (repo Repository) CheckExistenceThreadById(id int) (exists bool, err error) {
	checkExistenceThreadQuery := `
		SELECT EXISTS(SELECT 1 FROM threads WHERE id = $1) AS exists;
	`

	err = repo.conn.QueryRow(context.Background(), checkExistenceThreadQuery, id).
		Scan(&exists)
	if err != nil {
		return
	}

	return
}
