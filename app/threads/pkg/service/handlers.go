package threadService

import (
	"time"

	threadLiterals "github.com/masharpik/ForumVKEducation/app/threads/utils/literals"
	threadStructs "github.com/masharpik/ForumVKEducation/app/threads/utils/structs"
)

func (service service) CreateThread(thread threadStructs.Thread) (createdThread threadStructs.Thread, err error) {
	if thread.Created == "" {
		thread.Created = time.Now().Format(time.RFC3339)
	}

	createdThread, err = service.repo.CreateThread(thread)
	if err != nil {
		if err.Error() == threadLiterals.ThreadAlreadyExists {
			var infoErr error

			createdThread, infoErr = service.repo.GetThreadInfoBySlug(thread.Slug)
			if infoErr != nil {
				return threadStructs.Thread{}, infoErr
			}
		}
		return
	}

	return
}

func (service service) GetForumThreads(slug, since string, limit int, desc bool) (threads threadStructs.Threads, err error) {
	threads, err = service.repo.GetThreadsByForum(slug, since, limit, desc)
	if err != nil {
		return
	}

	return
}

func (service service) GetThreadIdBySlug(slug string) (id int, err error) {
	id, err = service.repo.GetThreadIdBySlug(slug)
	if err != nil {
		return
	}

	return
}

func (service service) CheckExistenceThreadById(id int) (exists bool, err error) {
	exists, err = service.repo.CheckExistenceThreadById(id)
	if err != nil {
		return
	}

	return
}

func (service service) VoteByThreadId(id int, nickname string, vote int) (thread threadStructs.Thread, err error) {
	thread, err = service.repo.VoteByThreadId(id, nickname, vote)
	if err != nil {
		return
	}

	return
}

func (service service) GetThreadInfoBySlug(slug string) (thread threadStructs.Thread, err error) {
	thread, err = service.repo.GetThreadInfoBySlug(slug)
	if err != nil {
		return
	}

	return
}

func (service service) GetThreadInfoById(id int) (thread threadStructs.Thread, err error) {
	thread, err = service.repo.GetThreadInfoById(id)
	if err != nil {
		return
	}

	return
}

func (service service) UpdateThreadInfoBySlug(slug, title, message string) (thread threadStructs.Thread, err error) {
	thread, err = service.repo.UpdateThreadBySlug(slug, title, message)
	if err != nil {
		return
	}

	return
}

func (service service) UpdateThreadInfoById(id int, title, message string) (thread threadStructs.Thread, err error) {
	thread, err = service.repo.UpdateThreadById(id, title, message)
	if err != nil {
		return
	}

	return
}

func (service service) GetForumSlugByThreadId(threadId int) (slug string, err error) {
	slug, err = service.repo.GetForumSlugByThreadId(threadId)
	if err != nil {
		return
	}

	return
}
