package interfaces

import threadStructs "github.com/masharpik/ForumVKEducation/app/threads/utils/structs"

type IRepository interface {
	CreateThread(threadStructs.Thread) (threadStructs.Thread, error)
	GetThreadInfoBySlug(string) (threadStructs.Thread, error)
	GetThreadInfoById(int) (threadStructs.Thread, error)
	GetThreadIdBySlug(string) (int, error)
	GetForumSlugByThreadId(int) (string, error)
	UpdateThreadBySlug(string, string, string) (threadStructs.Thread, error)
	UpdateThreadById(int, string, string) (threadStructs.Thread, error)
	VoteByThreadId(int, string, int) (threadStructs.Thread, error)
	GetThreadsByForum(string, string, int, bool) (threadStructs.Threads, error)
	CheckExistenceThreadById(int) (bool, error)
}
