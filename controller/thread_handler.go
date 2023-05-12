package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martin-lin-cw/goose-reddit/goreddit"
)

type ThreadController interface {
	GetThread(ctx *gin.Context) (goreddit.Thread, error)
}
type threadController struct {
	goreddit.ThreadStore
}

func NewThreadController(store goreddit.Store) ThreadController {
	return &threadController{ThreadStore: store}
}

func (c *threadController) GetThread(ctx *gin.Context) (goreddit.Thread, error) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return goreddit.Thread{}, fmt.Errorf("error GetThread parseInt: %w", err)
	}
	t, err := c.ThreadStore.Thread(int64(id))
	if err != nil {
		return goreddit.Thread{}, fmt.Errorf("error GetThread store: %w", err)
	}
	return t, nil
}
