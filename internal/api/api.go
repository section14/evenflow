package api

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Server struct {
	tpls *template.Template
}

func NewServer(t *template.Template) *Server {
	return &Server{tpls: t}
}

func handlers(mux *chi.Mux, s *Server) {
	pageRoutes(mux, s)
}

/*
func findAndParseTemplates(rootDir string, funcMap template.FuncMap) (*template.Template, error) {
    cleanRoot := filepath.Clean(rootDir)
    pfx := len(cleanRoot)+1
    root := template.New("")

    err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
        if !info.IsDir() && strings.HasSuffix(path, ".html") {
            if e1 != nil {
                return e1
            }

            b, e2 := os.ReadFile(path)
            if e2 != nil {
                return e2
            }

            name := path[pfx:]
            t := root.New(name).Funcs(funcMap)
            _, e2 = t.Parse(string(b))
            if e2 != nil {
                return e2
            }
        }

        return nil
    })

    return root, err
}
*/

func findAndParseTemplates(files fs.FS, rootDir string, funcMap template.FuncMap) (*template.Template, error) {
	cleanRoot := filepath.Clean(rootDir)
	//pfx := len(cleanRoot) + 1
	root := template.New("")

    err := fs.WalkDir(files, cleanRoot, func(path string, d fs.DirEntry, e1 error) error {
        if !d.IsDir() && strings.HasSuffix(path, ".html") {
			if e1 != nil {
				return e1
			}

			//b, e2 := os.ReadFile(path)
			b, e2 := fs.ReadFile(files, path)
			if e2 != nil {
				return e2
			}

			//name := path[pfx:]
			name := path

            fmt.Println("path name: ", name)
			t := root.New(name).Funcs(funcMap)
			_, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}

        }

        return nil
    })

	return root, err
}

func Serve(templates embed.FS) {
	sub, err := fs.Sub(templates, "templates")
	if err != nil {
        log.Fatal("couldn't setup embedded templates directory: ", err)
	}

	//t, err := template.ParseFS(tpl)
	t, err := findAndParseTemplates(sub, "", nil)
	if err != nil {
        log.Fatal("couldn't parse templates: ", err)
	}

	mux := chi.NewRouter()
	//mux.Use(middleware.Logger)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	server := NewServer(t)

	handlers(mux, server)

	//addr := fmt.Sprintf("%s:%s", env.Location, env.Port)
	addr := fmt.Sprintf("%s:%s", "localhost", "8080")

	s := &http.Server{
		Addr:    addr,
		Handler: mux,
		// other settings omitted
	}

    fmt.Println("serving on localhost:8080...")
	log.Fatal(s.ListenAndServe())
}
