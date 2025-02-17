package internals

import (
	"context"
	"htmx-test/views"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

const appTimeout = time.Second * 10

func render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}

func (app *Config) indexPageHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()

		viewsTodos := []*views.Todo{
			{
				Id:          "1",
				Description: "Lorem Ipsum",
			},
			{
				Id:          "2",
				Description: "Batman",
			},
			{
				Id:          "3",
				Description: "Tyler durden",
			},
		}

		render(ctx, http.StatusOK, views.Index(viewsTodos))
	}
}
