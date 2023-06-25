package interfaces

import postStructs "github.com/masharpik/ForumVKEducation/app/posts/utils/structs"

type IRepository interface {
	CreatePostsByThreadId(int, string, postStructs.Posts) (postStructs.Posts, error)
	GetPostInfo(int) (postStructs.Post, error)
	UpdatePostInfo(int, string) (postStructs.Post, error)
	GetFlatPostsByThreadId(int, int, int, string, bool) (postStructs.Posts, error)
	GetTreePostsByThreadId(int, int, int, string, bool) (postStructs.Posts, error)
	GetPostsParentTreeByThreadId(int, int, int, string, bool) (postStructs.Posts, error)
}
