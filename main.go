package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/.well-known/terraform.json", func(w http.ResponseWriter, r *http.Request) {
		println(r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(`{"modules.v1": "https://127.0.0.1:6966/v1/00576966"}`))
		//w.Write([]byte(`{"login.v1": {"client": "terraform-cli","grant_types": ["authz_code"],"authz": "/oauth/authorization","token": "/oauth/token","ports": [2000, 10010]}}`))
	})
	r.Get("/oauth/authorization", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	r.Get("/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	r.Get("/v1/00576966/hashicorp/consul/aws/versions", func(w http.ResponseWriter, r *http.Request) {
		println(r.Header.Get("Authorization"))
		w.WriteHeader(401)
		w.Write([]byte(`{"error_code":"123","error_msg":"321"}`))
		//w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(200)
		//w.Write([]byte(`{"modules": [{"versions": [{"version": "1.0.0"},{"version": "1.1.0"},{"version": "2.0.0"}]}]}`))
	})
	r.Get("/v1/00576966/hashicorp/consul/aws/1.0.0/download", func(w http.ResponseWriter, r *http.Request) {
		println(r.Header.Get("Authorization"))
		w.Header().Set("X-Terraform-Get", "")
		w.WriteHeader(204)

	})

	http.ListenAndServeTLS(":6966", "", "", r)
}
