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

## Create migration SQL file

create a migration SQL file by<br>
`goose create file_name sql`

put SQL below `-- +goose Up` for migrate up<br>
put SQL below `-- +goose Down` for migrate down

for example:<br>
create the TABLE post when migrate up<br>
drop the TABLE post when migrate down

```sql
-- +goose Up
CREATE TABLE post(
   id BIGINT(20) PRIMARY KEY
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +goose Down
DROP TABLE post;
```

## Using goose to migrate

goose syntax `goose DRIVER DBSTRING CMD`

for this example:

- dirver=mysql
- dbsrting="user:password@tcp(localhost:3306)/dbname?parseTime=true"

migrate up, CMD=up<br>

```console
goose mysql "user:password@tcp(localhost:3306)/dbname?parseTime=true" up
```

migrate down, CMD=down<br>

```console
goose mysql "user:password@tcp(localhost:3306)/dbname?parseTime=true" dwon
```

check migration status, CMD=status

```console
goose mysql "user:password@tcp(localhost:3306)/dbname?parseTime=true" status
```

## Implements store interface inside mysql package

1. Install sqlx with

   ```console
   go get github.com/jmoiron/sqlx
   ```

2. Implement ThreadStore, PostStore, and CommentStore with mySql. Then use Store to merge every single store.

3. Create and connect db with sqlx

   ```go
   func NewStore(dataSourceName string) (*Store, error) {
      db, err := sqlx.Open("mysql", dataSourceName)
      if err != nil {
         return nil, fmt.Errorf("error opening database: %w", err)
      }
      if err := db.Ping(); err != nil {
         return nil, fmt.Errorf("error connecting to database: %w", err)
      }
      return &Store{
         ThreadStore:  NewThreadStore(db),
         PostStore:    NewPostStore(db),
         CommentStore: NewCommentStore(db),
      }, nil
   }
   ```

4. Install dependency lib for mySql driver

   ```console
   $ go get -u github.com/go-sql-driver/mysql
   ```

   and import in store

   ```go
   import (
      "github.com/jmoiron/sqlx"
      _ "github.com/go-sql-driver/mysql"
   )
   ```
