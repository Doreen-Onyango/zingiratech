package utils

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// TemplateCache holds precompiled templates for efficient rendering.
var (
	TemplateCache = make(map[string]*template.Template)
	fn            = template.FuncMap{}
	mu            sync.RWMutex
)

// RenderServerErrorTemplate renders a fallback error page for server errors.
func RenderServerErrorTemplate(w http.ResponseWriter, statusCode int, errMsg string) {
	data := struct {
		StatusCode int
		Error      string
	}{
		StatusCode: statusCode,
		Error:      errMsg,
	}

	tmpl := `
<!DOCTYPE html>
<html>
<head><title>Error {{.StatusCode}}</title></head>
<body>
    <h1>Error {{.StatusCode}}</h1>
    <p>{{.Error}}</p>
</body>
</html>`

	t, err := template.New("error").Parse(tmpl)
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	_ = t.Execute(w, data)
}

// RenderTemplate renders a cached template with the given data.
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	mu.RLock()
	t, ok := TemplateCache[tmpl]
	mu.RUnlock()
	if !ok {
		errMsg := fmt.Sprintf("Template %s not found", tmpl)
		log.Printf("ERROR: %s", errMsg)
		RenderServerErrorTemplate(w, http.StatusNotFound, errMsg)
		return
	}

	if err := t.Execute(w, data); err != nil {
		errMsg := fmt.Sprintf("Error rendering template: %v", err)
		log.Printf("ERROR: %s", errMsg)
		RenderServerErrorTemplate(w, http.StatusInternalServerError, errMsg)
	}
}

// LoadTemplates loads and caches templates from the specified directory.
func LoadTemplates() error {
	cache := map[string]*template.Template{}

	baseDir, err := GetProjectRootPath("frontend", "templates")
	if err != nil {
		return fmt.Errorf("could not find project root: %w", err)
	}

	pagePattern := filepath.Join(baseDir, "*.page.html")
	pages, err := filepath.Glob(pagePattern)
	if err != nil {
		return fmt.Errorf("error finding page templates: %w", err)
	}

	layoutPattern := filepath.Join(baseDir, "*.layout.html")
	layouts, err := filepath.Glob(layoutPattern)
	if err != nil {
		return fmt.Errorf("error finding layout templates: %w", err)
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(fn).ParseFiles(page)
		if err != nil {
			return fmt.Errorf("error parsing page template %s: %w", name, err)
		}

		if len(layouts) > 0 {
			ts, err = ts.ParseGlob(layoutPattern)
			if err != nil {
				return fmt.Errorf("error parsing layout templates: %w", err)
			}
		}

		cache[name] = ts
		log.Printf("Loaded template: %s", name)
	}

	mu.Lock()
	TemplateCache = cache
	mu.Unlock()

	return nil
}

// RegisterFunc registers a custom function for use in templates.
func RegisterFunc(name string, function interface{}) {
	mu.Lock()
	fn[name] = function
	mu.Unlock()
}
