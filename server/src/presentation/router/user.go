package router

import (
	"database/sql"

	"github.com/Masuda-1246/shares/usecase"
	"github.com/Masuda-1246/shares/infrastructure/persistence"
	"github.com/Masuda-1246/shares/presentation/handler"
)

func (r *Router) InitUserRouter(db *sql.DB) {
	ur := persistence.NewUserRepository(db)
	uu := usecase.NewUserUseCase(ur)
	uh := handler.NewUserHandler(uu)

	g := r.Engine.Group("/user")
	g.GET("/:id", uh.GetByID)
	g.GET("/email/:email", uh.GetByEmail)
	g.POST("/", uh.CreateUser)
	g.POST("/login", uh.LoginUser)
	g.PUT("/:id", uh.UpdateUser)
	g.DELETE("/:id", uh.DeleteUser)
}