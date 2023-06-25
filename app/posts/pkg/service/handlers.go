package postService

import (
	forumStructs "github.com/masharpik/ForumVKEducation/app/forums/utils/structs"
	postStructs "github.com/masharpik/ForumVKEducation/app/posts/utils/structs"
	threadStructs "github.com/masharpik/ForumVKEducation/app/threads/utils/structs"
	userStructs "github.com/masharpik/ForumVKEducation/app/users/utils/structs"
)

func (service service) CreatePostsByThreadId(id int, forumSlug string, posts postStructs.Posts) (createdPosts postStructs.Posts, err error) {
	createdPosts, err = service.repo.CreatePostsByThreadId(id, forumSlug, posts)
	if err != nil {
		return
	}

	return
}

func (service service) GetPostInfo(id int, relatedUser, relatedForum, relatedThread bool) (postFull postStructs.PostFull, err error) {
	post, err := service.repo.GetPostInfo(id)
	if err != nil {
		return
	}
	postFull.Post = post

	if relatedUser {
		var user userStructs.User
		user, err = service.userRepo.GetUserInfo(post.Author)
		if err != nil {
			return
		}

		postFull.Author = &user
	}

	if relatedForum {
		var forum forumStructs.Forum
		forum, err = service.forumRepo.GetForumInfoBySlug(post.Forum)
		if err != nil {
			return
		}

		postFull.Forum = &forum
	}

	if relatedThread {
		var thread threadStructs.Thread
		thread, err = service.threadRepo.GetThreadInfoById(int(post.Thread))
		if err != nil {
			return
		}

		postFull.Thread = &thread
	}

	return
}

func (service service) UpdatePostInfo(id int, newMessage string) (post postStructs.Post, err error) {
	if newMessage == "" {
		post, err = service.repo.GetPostInfo(id)
		return
	}

	post, err = service.repo.UpdatePostInfo(id, newMessage)
	return
}

func (service service) GetPostsByThreadId(id, since, limit int, sort string, desc bool) (posts postStructs.Posts, err error) {
	switch sort {
	case "tree":
		posts, err = service.repo.GetTreePostsByThreadId(id, since, limit, sort, desc)
	case "parent_tree":
		posts, err = service.repo.GetPostsParentTreeByThreadId(id, since, limit, sort, desc)
	case "flat":
		posts, err = service.repo.GetFlatPostsByThreadId(id, since, limit, sort, desc)
	}
	if err != nil {
		return
	}

	return
}
