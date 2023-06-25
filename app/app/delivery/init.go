package initDelivery

import (
	"github.com/gorilla/mux"

	"github.com/masharpik/ForumVKEducation/app/app/middleware"
	initRepository "github.com/masharpik/ForumVKEducation/app/app/repository"
	forumDelivery "github.com/masharpik/ForumVKEducation/app/forums/pkg/delivery"
	forumRepository "github.com/masharpik/ForumVKEducation/app/forums/pkg/repository"
	forumService "github.com/masharpik/ForumVKEducation/app/forums/pkg/service"
	postDelivery "github.com/masharpik/ForumVKEducation/app/posts/pkg/delivery"
	postRepository "github.com/masharpik/ForumVKEducation/app/posts/pkg/repository"
	postService "github.com/masharpik/ForumVKEducation/app/posts/pkg/service"
	serviceDelivery "github.com/masharpik/ForumVKEducation/app/service/pkg/delivery"
	serviceRepository "github.com/masharpik/ForumVKEducation/app/service/pkg/repository"
	serviceService "github.com/masharpik/ForumVKEducation/app/service/pkg/service"
	threadDelivery "github.com/masharpik/ForumVKEducation/app/threads/pkg/delivery"
	threadRepository "github.com/masharpik/ForumVKEducation/app/threads/pkg/repository"
	threadService "github.com/masharpik/ForumVKEducation/app/threads/pkg/service"
	userDelivery "github.com/masharpik/ForumVKEducation/app/users/pkg/delivery"
	userRepository "github.com/masharpik/ForumVKEducation/app/users/pkg/repository"
	userService "github.com/masharpik/ForumVKEducation/app/users/pkg/service"
)

func RegisterHandlers(r *mux.Router) error {
	r.Use(middleware.JSONMiddleware)

	conn, err := initRepository.GetConnectionDB()
	if err != nil {
		return err
	}

	userRepo := userRepository.NewRepository(conn)
	postRepo := postRepository.NewRepository(conn)
	threadRepo := threadRepository.NewRepository(conn)
	forumRepo := forumRepository.NewRepository(conn)
	serviceRepo := serviceRepository.NewRepository(conn)

	userService := userService.NewService(userRepo)
	postService := postService.NewService(postRepo, userRepo, forumRepo, threadRepo)
	threadService := threadService.NewService(threadRepo)
	forumService := forumService.NewService(forumRepo, userRepo)
	serviceService := serviceService.NewService(serviceRepo)

	userRouter := r.PathPrefix("/user").Subrouter()
	userDelivery.RegisterHandlers(userRouter, userService)

	postRouter := r.PathPrefix("/post").Subrouter()
	postDelivery.RegisterHandlers(postRouter, postService)

	threadRouter := r.PathPrefix("/thread").Subrouter()
	threadDelivery.RegisterHandlers(threadRouter, threadService, postService)

	forumRouter := r.PathPrefix("/forum").Subrouter()
	forumDelivery.RegisterHandlers(forumRouter, forumService, threadService, userService)

	serviceRouter := r.PathPrefix("/service").Subrouter()
	serviceDelivery.RegisterHandlers(serviceRouter, serviceService)

	return nil
}
