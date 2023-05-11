# Learn Goose migration with Reddit example

## Setup project

1. Init go project with

   `go mod init github.com/...`

2. Install goose with

   `go install github.com/pressly/goose/v3/cmd/goose@latest`

## Create Reddit model

1. Define structure for

   - Thread, Post, and Comment

2. Create interface for manage those model
   - ThreadStore
     - Thread: get a thread by id
     - Threads: get all threads
     - CreateThread
     - UpdateThread
     - DeleteThread
   - PostStore
     - Post: get a post by id
     - PostsByThread: get all post belong to a thread id
     - CreatePost
     - UpdatePost
     - DeletePost
   - CommentStore:
     - Comment: get a comment by id
     - CommentsByPost: get all comments belong to a post id
     - CreateComment
     - UpdateComment
     - DeleteComment
