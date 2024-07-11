/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nivesh-star/ondc/src/api"
	"github.com/nivesh-star/ondc/src/common/logger"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"api.OnCancelPost",
		strings.ToUpper("Post"),
		"/on_cancel",
		api.OnCancelPost,
	},

	Route{
		"api.OnConfirmPost",
		strings.ToUpper("Post"),
		"/on_confirm",
		api.OnConfirmPost,
	},

	Route{
		"api.OnInitPost",
		strings.ToUpper("Post"),
		"/on_init",
		api.OnInitPost,
	},

	Route{
		"api.OnRatingPost",
		strings.ToUpper("Post"),
		"/on_rating",
		api.OnRatingPost,
	},

	Route{
		"api.OnSearchPost",
		strings.ToUpper("Post"),
		"/on_search",
		api.OnSearchPost,
	},

	Route{
		"api.OnSelectPost",
		strings.ToUpper("Post"),
		"/on_select",
		api.OnSelectPost,
	},

	Route{
		"api.OnStatusPost",
		strings.ToUpper("Post"),
		"/on_status",
		api.OnStatusPost,
	},

	Route{
		"api.OnSupportPost",
		strings.ToUpper("Post"),
		"/on_support",
		api.OnSupportPost,
	},

	Route{
		"api.OnTrackPost",
		strings.ToUpper("Post"),
		"/on_track",
		api.OnTrackPost,
	},

	Route{
		"api.OnUpdatePost",
		strings.ToUpper("Post"),
		"/on_update",
		api.OnUpdatePost,
	},

	Route{
		"api.OnSearchPost",
		strings.ToUpper("Post"),
		"/on_search",
		api.OnSearchPost,
	},

	Route{
		"SearchPost",
		strings.ToUpper("Post"),
		"/search",
		api.SearchPost,
	},

	// Route{
	// 	"CancelPost",
	// 	strings.ToUpper("Post"),
	// 	"/cancel",
	// 	api.CancelPost,
	// },

	// Route{
	// 	"ConfirmPost",
	// 	strings.ToUpper("Post"),
	// 	"/confirm",
	// 	api.ConfirmPost,
	// },

	// Route{
	// 	"InitPost",
	// 	strings.ToUpper("Post"),
	// 	"/init",
	// 	api.InitPost,
	// },

	// Route{
	// 	"RatingPost",
	// 	strings.ToUpper("Post"),
	// 	"/rating",
	// 	api.RatingPost,
	// },

	// Route{
	// 	"SearchPost",
	// 	strings.ToUpper("Post"),
	// 	"/search",
	// 	api.SearchPost,
	// },

	// Route{
	// 	"SelectPost",
	// 	strings.ToUpper("Post"),
	// 	"/select",
	// 	api.SelectPost,
	// },

	// Route{
	// 	"StatusPost",
	// 	strings.ToUpper("Post"),
	// 	"/status",
	// 	api.StatusPost,
	// },

	// Route{
	// 	"SupportPost",
	// 	strings.ToUpper("Post"),
	// 	"/support",
	// 	api.SupportPost,
	// },

	// Route{
	// 	"TrackPost",
	// 	strings.ToUpper("Post"),
	// 	"/track",
	// 	api.TrackPost,
	// },

	// Route{
	// 	"UpdatePost",
	// 	strings.ToUpper("Post"),
	// 	"/update",
	// 	api.UpdatePost,
	//},
}
