package app

import "github.com/Rifqi2000/go_shop/app/controllers"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
