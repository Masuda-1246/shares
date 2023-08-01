package main

import (
	"github.com/Masuda-1246/shares/infrastructure/database"
	"github.com/Masuda-1246/shares/presentation/router"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r := router.NewRouter()
	r.InitUserRouter(db)
	r.InitPostRouter(db)
	r.Serve()
}