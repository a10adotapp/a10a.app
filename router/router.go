package router

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/a10adotapp/a10a.app/ent"
	"github.com/a10adotapp/a10a.app/service/finschia"
	"github.com/a10adotapp/a10a.app/service/kusogeeeeeenft"
	"github.com/a10adotapp/a10a.app/service/linenft"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func NewRouter(entClient *ent.Client) *chi.Mux {
	router := chi.NewRouter()

	router.Use(chimiddleware.Timeout(10 * time.Second))

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("{\"message\":\"method not allowed\"}"))
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{\"message\":\"not found\"}"))
	})

	router.Get("/b", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://bookmarks.a10a.app", http.StatusMovedPermanently)
	})

	router.Route("/finschia", FinschiaRoute(finschia.NewFinschiaService(entClient)))
	router.Route("/kusogeeeeeenft", KusogeeeeeeNFTRoute(kusogeeeeeenft.NewKusogeeeeeeNFTService(entClient)))
	router.Route("/line-nft", LINENFTRoute(linenft.NewLINENFTService(entClient)))

	router.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("{\"message\":\"%s\"}\n", err.Error())))
			return
		}

		defer file.Close()

		dirname := filepath.Join(os.Getenv("UPLOADED_FILE_DIR"), time.Now().Format("20060102"))
		if len(dirname) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"message\":\"no uploaded file directory specified\"}\n"))
			return
		}

		if err := os.MkdirAll(dirname, os.ModePerm); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("{\"message\":\"%s\"}\n", err.Error())))
			return
		}

		filePath := filepath.Join(dirname, fileHeader.Filename)

		fileInfo, err := os.Stat(filePath)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("{\"message\":\"%s\"}\n", err.Error())))
			return
		}

		if fileInfo != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"message\":\"file already exists\"}\n"))
			return
		}

		fileToSave, err := os.Create(filePath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("{\"message\":\"%s\"}\n", err.Error())))
			return
		}

		defer fileToSave.Close()

		offset := int64(0)

		for {
			buf := make([]byte, 100)
			n, err := file.Read(buf)
			buf = buf[:n]
			if err == io.EOF {
				buf = buf[:n]
			}

			if n == 0 {
				break
			}

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("{\"message\":\"%s\"}\n", err.Error())))
				return
			}

			fileToSave.WriteAt(buf, offset)

			offset += int64(n)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"message\":\"OK\"}\n"))
	})

	return router
}
