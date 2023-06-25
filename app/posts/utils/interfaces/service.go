package interfaces

import postStructs "github.com/masharpik/ForumVKEducation/app/posts/utils/structs"

type IService interface {
	CreatePostsByThreadId(int, string, postStructs.Posts) (postStructs.Posts, error)
	GetPostInfo(int, bool, bool, bool) (postStructs.PostFull, error)
	UpdatePostInfo(int, string) (postStructs.Post, error)
	GetPostsByThreadId(int, int, int, string, bool) (postStructs.Posts, error)
}
