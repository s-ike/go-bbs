package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templは１つのテンプレートを表す
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTPはHTTPリクエストを処理する
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	logs, err := loadLogs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data["Logs"] = logs
	t.templ.Execute(w, data)
}

func main() {
	w := newWriter()
	http.Handle("/", &templateHandler{filename: "bbs.html"})
	http.Handle("/write", w)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
