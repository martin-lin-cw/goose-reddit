package goreddit

type Thread struct {
	Id          int64  `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

type Post struct {
	Id       int64  `db:"id"`
	ThreadId int64  `db:"thread_id"`
	Title    string `db:"title"`
	Content  string `db:"content"`
	Votes    int    `db:"votes"`
}

type Comment struct {
	Id      int64  `db:"id"`
	PostId  int64  `db:"post_id"`
	Content string `db:"content"`
	Votes   int    `db:"votes"`
}

type ThreadStore interface {
	Thread(id int64) (Thread, error)
	Threads() ([]Thread, error)
	CreateThread(t *Thread) error
	UpdateThread(t *Thread) error
	DeleteThread(id int64) error
}

type PostStore interface {
	Post(id int64) (Post, error)
	PostsByThread(threadId int64) ([]Post, error)
	CreatePost(p *Post) error
	UpdatePost(p *Post) error
	DeletePost(id int64) error
}

type CommentStore interface {
	Comment(id int64) (Comment, error)
	CommentsByPost(postId int64) ([]Comment, error)
	CreateComment(c *Comment) error
	UpdateComment(c *Comment) error
	DeleteComment(id int64) error
}

type Store interface {
	ThreadStore
	PostStore
	CommentStore
}
