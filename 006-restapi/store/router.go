package store

import "net/http"

var controller = &Controller{Repository: Repository{}}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route {
		"Authentication",
		"POST",
		"/get-token",
		controller.GetToken
	}
}