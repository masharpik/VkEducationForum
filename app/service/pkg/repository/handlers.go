package serviceRepository

import (
	"context"
	"sync"

	serviceStructs "github.com/masharpik/ForumVKEducation/app/service/utils/structs"
)

func (repo Repository) ClearDB() (err error) {
	clearDBQuery := `TRUNCATE votes, posts, threads, forums, users CASCADE;`
	_, err = repo.conn.Exec(context.Background(), clearDBQuery)

	return
}

func (repo Repository) GetInfoDB() (status serviceStructs.Status, err error) {
	var wg sync.WaitGroup
	queries := []string{"SELECT COUNT(*) AS user FROM users;", "SELECT COUNT(*) AS forum, COALESCE(SUM(threads), 0) AS thread, COALESCE(SUM(posts), 0) AS post FROM forums;"}
	errs := make(chan error, 2)

	wg.Add(1)
	go func(q string) {
		defer wg.Done()
		err := repo.conn.QueryRow(context.Background(), q).Scan(&status.User)
		if err != nil {
			errs <- err
		}
	}(queries[0])

	wg.Add(1)
	go func(q string) {
		defer wg.Done()
		err := repo.conn.QueryRow(context.Background(), q).Scan(&status.Forum, &status.Thread, &status.Post)
		if err != nil {
			errs <- err
		}
	}(queries[1])

	wg.Wait()
	close(errs)

	for e := range errs {
		if e != nil {
			err = e
			return
		}
	}

	return
}
