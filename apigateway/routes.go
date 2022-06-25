package apigateway

import (
	post "thesayedirfan/socialapi/services/postservice/handler"
	user "thesayedirfan/socialapi/services/userservice/handler"
	"thesayedirfan/socialapi/utils/middlerware"
)

func PublicRoutes() {

	V1.Post("SignUp", user.SignUp)
	V1.Post("SignIn", user.SignIn)
}

func PrivateRoutes() {

	userV1 := V1.Group("/user")
	postV1 := V1.Group("/post")

	V1.Use(middlerware.TokenValid)

	userV1.Post("refreshToken", user.Refresh)
	userV1.Get("view", user.UserView)
	userV1.Get("list", user.UserList)
	userV1.Post("follow", user.UserFollow)
	userV1.Post("unfollow", user.UserUnFollow)
	userV1.Get("following", user.GetUserFollowing)

	postV1.Post("create", post.CreatePost)
	postV1.Get("view", post.ViewPost)
	postV1.Post("delete", post.DeletePost)
	postV1.Get("user", post.GetUserPost)

	postV1.Post("comment", post.CommentToPost)
	postV1.Delete("comment", post.DeleteComment)
	postV1.Get("comments", post.GetPostComments)

}
