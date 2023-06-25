package threadRepository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masharpik/ForumVKEducation/app/threads/utils/interfaces"
)

type Repository struct {
	conn *pgxpool.Pool
}

func NewRepository(conn *pgxpool.Pool) (repo interfaces.IRepository) {
	repo = Repository{
		conn: conn,
	}
	return
}
