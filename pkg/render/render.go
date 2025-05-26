package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"github.com/MichaelVarianK/bookings/pkg/config"
	"github.com/MichaelVarianK/bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func addDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// Create Template Cache
	var tc map[string]*template.Template

	if (app.UseCache) {
		tc = app.TemplateCache	
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Get Requested Template Cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Couldn't crete a new template")
	}

	buf := new(bytes.Buffer)
	err := t.Execute(buf, td)
	if (err != nil) {
		fmt.Println(err)
	}

	// Render Template Cache
	_, err = buf.WriteTo(w)
	if (err != nil) {
		fmt.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)

	pages, err := filepath.Glob("./templates/*.page.html")
	if (err != nil) {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if (err != nil) {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if (err != nil) {
			return myCache, err
		}

		if (len(matches) > 0) {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if (err != nil) {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}




// func RenderTemplate(w http.ResponseWriter, tmpl string) {
// 	temp, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.html")
// 	err := temp.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }

// Building a simple template cache
// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	_, inMap := tc[t]

// 	if (!inMap) {
// 		fmt.Println("Creating template...")
// 		err = createRenderTemplate(t)
// 		if (err != nil) {
// 			fmt.Println(err)
// 			return
// 		}
// 	} else {
// 		fmt.Println("Using template cache...")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)

// 	if (err != nil) {
// 		fmt.Println(err)
// 		return
// 	}
// }

// func createRenderTemplate(t string) error {
// 	var path = []string {
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.html",
// 	}

// 	temp, err := template.ParseFiles(path...)
// 	if (err != nil) {
// 		return err
// 	}

// 	tc[t] = temp
// 	return nil
// }