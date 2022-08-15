package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/drawflow_app/database"
	"github.com/drawflow_app/models"
	"github.com/go-chi/chi"
)

type DiagramsResource struct{}

func (rs DiagramsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.GetAllPrograms)
	r.Post("/", rs.AddProgram)
	r.Put("/", rs.AddProgram)
	r.Post("/runCode", rs.RunCode)
	r.Delete("/", rs.DeleteProgram)

	r.Route("/{uid}", func(r chi.Router) {
		r.Use(ProgramCtx)
		r.Get("/", rs.GetOneProgram)
	})

	return r
}

func (rs DiagramsResource) GetAllPrograms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := database.DgraphClient.GetPrograms(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res.Json)
}

func (rs DiagramsResource) GetOneProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.Context().Value("uid").(string)
	resp, err := database.DgraphClient.GetWithId(id, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(resp.Json)
}

func (rs DiagramsResource) RunCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	buf := new(bytes.Buffer)

	buf.ReadFrom(r.Body)
	runCode := buf.String()
	file, err := os.Create("code.py")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(file, runCode)
	file.Close()

	cmd := exec.Command("python", "code.py")
	out, err := cmd.CombinedOutput()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(out)

}

func (rs DiagramsResource) AddProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var add models.Program

	err := json.NewDecoder(r.Body).Decode(&add)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pb, _ := json.Marshal(add)
	res, err := database.DgraphClient.Add(pb, r.Context())

	if err != nil {
		log.Fatal(err)
	}

	json, _ := json.Marshal(res)
	w.Write(json)
}

func (rs DiagramsResource) DeleteProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var delete models.Program
	err := json.NewDecoder(r.Body).Decode(&delete)

	if err != nil {
		log.Fatal(err)
	}
	pb, _ := json.Marshal(delete)
	resp, err := database.DgraphClient.Delete(pb, r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json, _ := json.Marshal(resp)
	w.Write(json)

}

func ProgramCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "uid", chi.URLParam(r, "uid"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
