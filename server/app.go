package main

import (
	"github.com/iris-contrib/iris-starter-kit/server/data/static"
	"github.com/iris-contrib/iris-starter-kit/server/data/templates"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/middleware/requestid"

	"github.com/olebedev/config"
)

// App struct.
// There is no singleton anti-pattern,
// all variables defined locally inside
// this struct.
type App struct {
	Server *iris.Application
	Conf   *config.Config
	React  *React
	API    *API
}

// NewApp returns initialized struct
// of main server application.
func NewApp(opts ...AppOptions) *App {
	options := AppOptions{}
	for _, i := range opts {
		options = i
		break
	}

	options.init()

	// Parse config yaml string from ./conf.go
	conf, err := config.ParseYaml(confString)
	Must(err)

	// Set config variables delivered from main.go:11
	// Variables defined as ./conf.go:3
	conf.Set("debug", debug)
	conf.Set("commitHash", commitHash)

	// Parse environ variables for defined
	// in config constants
	conf.Env()

	// Make an engine
	srv := iris.New()

	// Use precompiled embedded templates: remove any existing .go files before:
	// go-bindata -o ./data/templates/templates.go -pkg templates -fs -prefix "data/templates" ./data/templates/...
	srv.RegisterView(iris.HTML(templates.AssetFile(), ".html"))

	// Set up debug level for iris logger
	if conf.UBool("debug") {
		srv.Logger().SetLevel("debug")
	}

	// Middlewares that should be registered
	// everywhere, all routes, and even not founds.
	//
	// Logger, will execute the next and then log the request-response lifecycle.
	ac := accesslog.File("./access.log") // its Close is handled on CTRL/CMD+C automatically.
	srv.UseRouter(ac.Handler)
	// Map app and uuid for every requests
	srv.UseRouter(requestid.New())
	// Recover middleware.
	srv.UseRouter(recover.New())

	// Favicon
	srv.Favicon("./data/static/images/favicon.ico")

	// Initialize the application
	app := &App{
		Conf:   conf,
		Server: srv,
		React: NewReact(
			conf.UString("duktape.path"),
			conf.UBool("debug"),
			srv,
		),
	}

	// Serve static via bindata ./data/static
	// go-bindata -o ./data/static/static.go -pkg static -fs -prefix "data/static" ./data/static/...
	srv.HandleDir("/static", static.AssetFile())

	api := NewAPI(app)

	// Bind api hadling for URL api.prefix
	api.Bind(
		app.Server.Party(
			app.Conf.UString("api.prefix"),
		),
	)

	// Handle via react app
	//
	// Registers / and /anything/here on GET and HEAD http methods.
	// srv.HandleMany("GET HEAD", "/ /{p:path}", app.React.Handle)
	// Or:
	//
	// handle root with react
	srv.Get("/", app.React.Handle)
	// handle anything expect /static/ and /api/v1/conf with react as well
	srv.Get("/{p:path}", app.React.Handle)

	return app
}

// Run runs the app
func (app *App) Run() {
	Must(app.Server.Listen(":" + app.Conf.UString("port")))
}

// AppOptions is options struct
type AppOptions struct{}

func (ao *AppOptions) init() { /* write your own*/ }
