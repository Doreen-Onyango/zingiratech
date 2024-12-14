package utils

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"text/template"
)

// TemplateCache holds precompiled templates for efficient rendering.
var (
	TemplateCache = make(map[string]*template.Template)
	fn            = template.FuncMap{}
	mu            sync.RWMutex // Ensure concurrency safety
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
