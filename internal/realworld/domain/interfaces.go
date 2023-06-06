package domain

type UserService interface {
	Login(email string, password string) (User, error)
	Register(email string, password string, username string) (User, error)
	GetUser(id int) (User, error)
	GetUserByEmail(email string) (User, error)
	GetUserByUsername(username string) (User, error)
	UpdateUser(id int, email string, username string, bio string, image string) (User, error)
}

type ArticleService interface {
	GetArticle(id int) (Article, error)
	GetArticles(offset int, limit int, tag string, author string, favorited string) ([]Article, int, error)
	GetArticlesFeed(offset int, limit int, tag string, author string, favorited string, userID int) ([]Article, int, error)
	CreateArticle(userID int, title string, description string, body string, tagList []string) (Article, error)
	UpdateArticle(id int, title string, description string, body string) (Article, error)
	DeleteArticle(id int) error
	AddArticleToFavorites(articleID int, userID int) (Article, error)
	RemoveArticleFromFavorites(articleID int, userID int) (Article, error)
	GetArticleComments(articleID int) ([]Comment, error)
	AddComment(articleID int, userID int, body string) (Comment, error)
	DeleteComment(articleID int, commentID int) error
}

type TagService interface {
	GetTags() ([]string, error)
}

type FollowService interface {
	FollowUser(followerID int, followedID int) (Profile, error)
	UnfollowUser(followerID int, followedID int) (Profile, error)
	GetFollowers(userID int) ([]Profile, error)
	GetFollowed(userID int) ([]Profile, error)
}

type FavoriteService interface {
	GetFavorites(userID int) ([]Article, error)
}
