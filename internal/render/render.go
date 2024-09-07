package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/babtun123/MelodyV2/internal/config"
	"github.com/babtun123/MelodyV2/internal/models"
	"github.com/justinas/nosurf"
)

// Used when rendering Page???
var functions = template.FuncMap{}

var app *config.AppConfig
var pathToTemplates = "./templates"

// NewTemplates sets the config for the template package. "Used in main"
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData is used to pass data to cached page
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate to render html template using html/template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
		// fmt.Println("Cache was used")
	} else {
		tc, _ = CreateTemplateCache()
		// app.UseCache = true
	}

	// Get requested template cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)
	_ = t.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// CreateTemplateCache caches every tmpl file we have.
func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{} // Same thing as the code above

	// get all the of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	// Range through each file in files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page) // "ts" is template set
		if err != nil {
			return myCache, err
		}

		// Returns a string slice containing files ending  with *.layout.tmpl
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))

		// As long as we have a file, parse it with the above files.
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		}
		if err != nil {
			return myCache, err
		}

		// Add ts to cache map
		myCache[name] = ts
	}

	return myCache, nil
}
