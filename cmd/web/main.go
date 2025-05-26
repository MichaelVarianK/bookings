package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MichaelVarianK/bookings/pkg/config"
	"github.com/MichaelVarianK/bookings/pkg/handlers"
	"github.com/MichaelVarianK/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var session *scs.SessionManager
var app config.AppConfig

func main() {
	// changes this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if (err != nil) {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = true

    repo := handlers.NewRepo(&app)
    handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if (err != nil) {
		log.Fatal(err)
	}
}
