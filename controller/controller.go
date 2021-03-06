package controller

import(
	"html/template"
	"net/http"
)

var(
	homeController home
	shopController shop
)

func Startup (templates map[string]*template.Template){

	homeController.homeTemplate = templates["home.html"]
	homeController.loginTemplate = templates["login.html"]

	homeController.registerRoutes()

	shopController.shopTemplate = templates["shop.html"]
	shopController.registerRoutes()


	shopController.categoryTemplate = templates["shop_details.html"]

	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}