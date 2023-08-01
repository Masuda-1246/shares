package router

import (
	"database/sql"

	"github.com/Masuda-1246/shares/infrastructure/persistence"
	"github.com/Masuda-1246/shares/presentation/handler"
	"github.com/Masuda-1246/shares/usecase"
)

func (r *Router) InitPostRouter(db *sql.DB) {
	pr := persistence.NewPostRepository(db)
	pu := usecase.NewPostUsecase(pr)
	ph := handler.NewPostHandler(pu)

	g := r.Engine.Group("/post")
	g.GET("/:id", ph.GetByID)
	g.GET("/user/:userID", ph.GetByUserID)
	g.GET("/", ph.GetAll)
	g.POST("/", ph.CreatePost)
	g.PUT("/:id", ph.UpdatePost)
	g.DELETE("/:id", ph.DeletePost)
}
