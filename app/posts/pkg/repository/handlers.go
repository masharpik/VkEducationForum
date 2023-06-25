package postRepository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	postLiterals "github.com/masharpik/ForumVKEducation/app/posts/utils/literals"
	postStructs "github.com/masharpik/ForumVKEducation/app/posts/utils/structs"
	mainLiterals "github.com/masharpik/ForumVKEducation/utils/literals"
)

func (repo Repository) CreatePostsByThreadId(id int, forumSlug string, posts postStructs.Posts) (postStructs.Posts, error) {
	if len(posts) == 0 {
		return make(postStructs.Posts, 0), nil
	}

	parentFirst := posts[0].Parent
	if parentFirst != 0 {
		var exists bool

		existenceParentQuery := `
			SELECT EXISTS(SELECT 1 FROM posts WHERE id = $1 AND thread = $2) AS exists;
		`

		err := repo.conn.QueryRow(context.Background(), existenceParentQuery, parentFirst, id).Scan(&exists)
		if err != nil {
			return make(postStructs.Posts, 0), err
		}
		if !exists {
			err = errors.New(postLiterals.ParentNotInThread)
			return make(postStructs.Posts, 0), err
		}
	}

	var now string
	if posts[0].Created == "" {
		now = time.Now().Format(time.RFC3339)
	} else {
		now = posts[0].Created
	}

	createPostsQuery := `
		INSERT INTO posts (parent, author, message, forum, thread, created) VALUES`

	const formatString string = ` ($%d, $%d, $%d, $%d, $%d, $%d)`
	var formatPosts []string
	var values []interface{}
	for i, post := range posts {
		formatPosts = append(formatPosts, fmt.Sprintf(formatString, i*6+1, i*6+2, i*6+3, i*6+4, i*6+5, i*6+6))

		var parent interface{}
		postParent := post.Parent
		if postParent == 0 {
			parent = nil
		} else {
			parent = postParent
		}

		values = append(values,
			parent,
			post.Author,
			post.Message,
			forumSlug,
			id,
			now)
	}

	createPostsQuery += strings.Join(formatPosts, ",")
	createPostsQuery += " RETURNING id;"

	postRows, err := repo.conn.Query(context.Background(), createPostsQuery, values...)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case mainLiterals.CodeForeignKeyNotFound:
				err = errors.New(postLiterals.ParentOrAuthorFKError)
			}
		}

		return make(postStructs.Posts, 0), err
	}

	var i uint = 0
	idInt32 := int32(id)
	for postRows.Next() {
		posts[i].Forum = forumSlug
		posts[i].Thread = idInt32
		posts[i].Created = now
		err = postRows.Scan(&posts[i].Id)
		if err != nil {
			return make(postStructs.Posts, 0), err
		}
		i++
	}

	if err = postRows.Err(); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case mainLiterals.CodeForeignKeyNotFound:
				err = errors.New(postLiterals.ParentOrAuthorFKError)
			}
		}

		return make(postStructs.Posts, 0), err
	}

	return posts, nil
}

func (repo Repository) GetPostInfo(id int) (post postStructs.Post, err error) {
	getInfoQuery := `
		SELECT id, COALESCE(parent, 0), author, message, isEdited, forum, thread, created FROM posts WHERE id = $1;
	`
	var t time.Time
	err = repo.conn.QueryRow(context.Background(), getInfoQuery, id).
		Scan(
			&post.Id,
			&post.Parent,
			&post.Author,
			&post.Message,
			&post.IsEdited,
			&post.Forum,
			&post.Thread,
			&t,
		)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.New(postLiterals.PostNotFound)
		}

		return
	}
	post.Created = t.Format(time.RFC3339)

	return
}

func (repo Repository) UpdatePostInfo(id int, newMessage string) (post postStructs.Post, err error) {
	updateMsgQuery := `
		UPDATE posts SET message = $1 WHERE id = $2 RETURNING id, COALESCE(parent, 0), author, message, isEdited, forum, thread, created;
	`

	var t time.Time
	err = repo.conn.QueryRow(context.Background(), updateMsgQuery, newMessage, id).
		Scan(
			&post.Id,
			&post.Parent,
			&post.Author,
			&post.Message,
			&post.IsEdited,
			&post.Forum,
			&post.Thread,
			&t,
		)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.New(postLiterals.PostNotFound)
		}

		return
	}
	post.Created = t.Format(time.RFC3339)

	return
}

func (repo Repository) GetFlatPostsByThreadId(id, since, limit int, sort string, desc bool) (posts postStructs.Posts, err error) {
	posts = make(postStructs.Posts, 0)
	oper := ">"
	order := "ASC"

	if desc {
		oper = "<"
		order = "DESC"
	}

	var getMessagesQuery string

	if since == 0 {
		getMessagesQuery = fmt.Sprintf(`
SELECT id, COALESCE(parent, 0), author, message, isEdited, forum, thread, created 
FROM posts 
WHERE thread = $1  
ORDER BY created %s, id %s
LIMIT $2;
`, order, order)
	} else {
		getMessagesQuery = fmt.Sprintf(`
SELECT id, COALESCE(parent, 0), author, message, isEdited, forum, thread, created 
FROM posts 
WHERE thread = $1 AND id %s $2 
ORDER BY created %s, id %s
LIMIT $3;
`, oper, order, order)
	}

	var postRows pgx.Rows
	if since == 0 {
		postRows, err = repo.conn.Query(context.Background(), getMessagesQuery, id, limit)
	} else {
		postRows, err = repo.conn.Query(context.Background(), getMessagesQuery, id, since, limit)
	}
	if err != nil {
		return
	}

	for postRows.Next() {
		var post postStructs.Post
		var t time.Time

		err = postRows.Scan(
			&post.Id,
			&post.Parent,
			&post.Author,
			&post.Message,
			&post.IsEdited,
			&post.Forum,
			&post.Thread,
			&t,
		)
		if err != nil {
			return
		}
		post.Created = t.UTC().Format("2006-01-02T15:04:05.000Z")

		posts = append(posts, post)
	}

	return
}

func (repo Repository) GetTreePostsByThreadId(id, since, limit int, sort string, desc bool) (posts postStructs.Posts, err error) {
	posts = make(postStructs.Posts, 0)
	oper := ">"
	order := "ASC"

	if desc {
		oper = "<"
		order = "DESC"
	}

	var getMessagesQuery string

	if since == 0 {
		getMessagesQuery = fmt.Sprintf(`
SELECT id, COALESCE(parent, 0), author, message, isEdited, forum, thread, created 
FROM posts 
WHERE thread = $1 
ORDER BY num %s 
LIMIT $2;
`, order)
	} else {
		getMessagesQuery = fmt.Sprintf(`
WITH since_table AS (
	SELECT num FROM posts WHERE id = $1 LIMIT 1
)
SELECT id, COALESCE(parent, 0), author, message, isEdited, forum, thread, created
FROM posts
WHERE thread = $2 AND num %s (SELECT num FROM since_table)
ORDER BY num %s 
LIMIT $3;
`, oper, order)
	}

	var postRows pgx.Rows
	if since == 0 {
		postRows, err = repo.conn.Query(context.Background(), getMessagesQuery, id, limit)
	} else {
		postRows, err = repo.conn.Query(context.Background(), getMessagesQuery, since, id, limit)
	}
	if err != nil {
		return
	}

	for postRows.Next() {
		var post postStructs.Post
		var t time.Time

		err = postRows.Scan(
			&post.Id,
			&post.Parent,
			&post.Author,
			&post.Message,
			&post.IsEdited,
			&post.Forum,
			&post.Thread,
			&t,
		)
		if err != nil {
			return
		}
		post.Created = t.UTC().Format("2006-01-02T15:04:05.000Z")

		posts = append(posts, post)
	}

	return
}

func (repo Repository) GetPostsParentTreeByThreadId(id, since, limit int, sort string, desc bool) (posts postStructs.Posts, err error) {
	posts = make(postStructs.Posts, 0)
	order := "ASC"

	if desc {
		order = "DESC"
	}

	var getMessagesQuery string

	if since == 0 {
		getMessagesQuery = fmt.Sprintf(`
WITH root_posts AS (
	SELECT id, num 
	FROM posts 
	WHERE thread = $1 AND parent IS NULL 
	ORDER BY id %s 
	LIMIT $2 
)


SELECT p.id, COALESCE(parent, 0), author, message, isEdited, forum, thread, created 
FROM root_posts rp 
	LEFT JOIN posts p 
		ON rp.id = p.parent OR p.num <@ rp.num 
ORDER BY subpath(p.num, 0, 1) %s, p.num ASC;
`, order, order)
	} else {
		oper := ">"
		if desc {
			oper = "<"
		}
		getMessagesQuery = fmt.Sprintf(`
WITH since_table AS (
    SELECT subpath(num, 0, 1)::text::integer AS num FROM posts WHERE thread = $1 AND id = $2 LIMIT 1
), parent_table AS (
    SELECT id, num FROM posts WHERE thread = $1 AND parent IS NULL AND id %s (SELECT num FROM since_table) ORDER BY id %s LIMIT $3
)

SELECT p.id, COALESCE(parent, 0) AS parent, author, message, isedited, forum, thread, created 
FROM parent_table pt
	LEFT JOIN posts p ON p.num <@ pt.num
ORDER BY subpath(p.num, 0, 1) %s, p.num ASC
`, oper, order, order)
	}

	var postRows pgx.Rows
	if since == 0 {
		postRows, err = repo.conn.Query(context.Background(), getMessagesQuery, id, limit)
	} else {
		postRows, err = repo.conn.Query(context.Background(), getMessagesQuery, id, since, limit)
	}
	if err != nil {
		return
	}

	for postRows.Next() {
		var post postStructs.Post
		var t time.Time

		err = postRows.Scan(
			&post.Id,
			&post.Parent,
			&post.Author,
			&post.Message,
			&post.IsEdited,
			&post.Forum,
			&post.Thread,
			&t,
		)
		if err != nil {
			return
		}
		post.Created = t.UTC().Format("2006-01-02T15:04:05.000Z")

		posts = append(posts, post)
	}

	return
}
