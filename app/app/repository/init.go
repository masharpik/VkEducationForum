package initRepository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	mainLiterals "github.com/masharpik/ForumVKEducation/utils/literals"
	"github.com/masharpik/ForumVKEducation/utils/logger"
)

func loadConfigUrl() string {
	// host := os.Getenv("PQ_APP_HOST")
	// port := os.Getenv("PQ_APP_PORT")
	// user := os.Getenv("PQ_APP_USER")
	// pass := os.Getenv("PQ_APP_PASS")
	// name := os.Getenv("PQ_APP_NAME")

	return "postgres://forum_user:forum_pass@localhost:5432/forum_name"
}

func GetConnectionDB() (conn *pgxpool.Pool, err error) {
	url := loadConfigUrl()

	ticker := time.NewTicker(1 * time.Second)
	timer := time.NewTimer(2 * time.Minute)

	for {
		select {
		case <-timer.C:
			ticker.Stop()
			err = fmt.Errorf(mainLiterals.LogConnDBTimeout)
			return
		case <-ticker.C:
			conn, err = pgxpool.New(context.Background(), url)

			if err == nil {
				err = conn.Ping(context.Background())
				if err == nil {
					ticker.Stop()
					timer.Stop()
					logger.LogOperationSuccess(fmt.Sprintf(mainLiterals.LogConnDBSuccess, url))
				}
				return
			}
		}
	}
}
