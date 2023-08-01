package router

import (
	"database/sql"

	"github.com/Masuda-1246/shares/usecase"
	"github.com/Masuda-1246/shares/infrastructure/persistence"
	"github.com/Masuda-1246/shares/presentation/handler"
)

func (r *Router) InitCommentRouter(db *sql.DB) {
	cr := persistence.NewCommentRepository(db)
	cu := usecase.NewCommentUsecase(cr)
	ch := handler.NewCommentHandler(cu)

	g := r.Engine.Group("/comment")
	g.GET("/:id", ch.GetByID)
	g.GET("/user/:userID", ch.GetByUserID)
	g.GET("/post/:postID", ch.GetByPostID)
	g.GET("/", ch.GetAll)
	g.POST("/", ch.CreateComment)
	g.DELETE("/:id", ch.DeleteComment)
}