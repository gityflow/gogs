package cmd

import (
	"github.com/urfave/cli"
	"gopkg.in/macaron.v1"
	"github.com/gityflow/githorse/routes.v2"
	"github.com/gityflow/githorse/pkg/form"
	"github.com/go-macaron/binding"
	"net/http"
)

func routesV2(c *cli.Context,m *macaron.Macaron, reqSignIn, ignSignIn, ignSignInAndCsrf, reqSignOut macaron.Handler){
	bindIgnErr := binding.BindIgnErr
	// FIXME: not all routes need go through same middlewares.
	// Especially some AJAX requests, we can reduce middleware number to improve performance.
	// Routers.

	fs := http.FileServer(http.Dir("static"))
	m.Get("/githorse/", http.StripPrefix("/githorse/", fs))

	m.Get("/v2", ignSignIn, routes_v2.Home)
	m.Combo("/v2/install", routes_v2.InstallInit).Get(routes_v2.Install).
		Post(bindIgnErr(form.Install{}), routes_v2.InstallPost)
}
