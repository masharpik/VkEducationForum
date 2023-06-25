package interfaces

import threadStructs "github.com/masharpik/ForumVKEducation/app/threads/utils/structs"

type IService interface {
	CreateThread(threadStructs.Thread) (threadStructs.Thread, error)
	GetForumThreads(string, string, int, bool) (threadStructs.Threads, error)
	GetThreadIdBySlug(string) (int, error)
	CheckExistenceThreadById(int) (bool, error)
	VoteByThreadId(int, string, int) (threadStructs.Thread, error)
	GetThreadInfoBySlug(string) (threadStructs.Thread, error)
	GetThreadInfoById(int) (threadStructs.Thread, error)
	UpdateThreadInfoBySlug(string, string, string) (threadStructs.Thread, error)
	UpdateThreadInfoById(int, string, string) (threadStructs.Thread, error)
	GetForumSlugByThreadId(int) (string, error)
}
