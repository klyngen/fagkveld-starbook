package presentation

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	oidcdiscovery "github.com/klyngen/golang-oidc-discovery"
	"github.com/klyngen/packages/backend/starbook-auth/repository"
)

const userIdKey = "userId"

type AuthenticationConfig struct {
	Domain   string
	Scopes   string
	Issuer   string // Really not needed
	Audience string // Really not needed
}

type api struct {
	repository *repository.Repository
	router     *chi.Mux
}

func NewApi(repository *repository.Repository, config AuthenticationConfig) *api {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(authenticationMiddleware)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	api := api{
		repository: repository,
	}

	router.Route("/person", func(r chi.Router) {
		r.Get("/", api.handleGetPersons)
		r.Post("/", api.handlePostPerson)
	})

	router.Route("/star", func(r chi.Router) {
		r.Get("/", api.handleGetStars)
		r.Post("/", api.handlePostStar)
	})

	api.router = router

	return &api
}

func (a *api) Serve(port string) {
	fmt.Printf("Listening on port: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), a.router)
}

func getPublicKeys(config AuthenticationConfig) ([]oidcdiscovery.PublicKey, error) {
	client, err := oidcdiscovery.NewOidcDiscoveryClient(config.Domain)

	if err != nil {
		return nil, err
	}

	return client.GetCertificates()
}

func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			context := context.WithValue(r.Context(), userIdKey, "klingen")
			next.ServeHTTP(w, r.WithContext(context))
		})
}

func (a *api) handleGetPersons(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userIdKey).(string)

	persons, err := a.repository.GetPersonsByUserId(userId)

	if err != nil {
		fmt.Fprint(w, "Something wen to shit")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(persons)
}

func (a *api) handlePostPerson(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userIdKey).(string)

	decoder := json.NewDecoder(r.Body)
	var person repository.Person
	err := decoder.Decode(&person)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Unable to read body")
		return
	}

	person.BelongsTo = userId

	err = a.repository.CreatePerson(&person)

	if err != nil {
		log.Println(err.Error())
		fmt.Fprint(w, "Something went to shit")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(&person)
}

func (a *api) handleGetStars(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userIdKey).(string)

	stars, err := a.repository.GetStarsByUser(userId)

	if err != nil {
		log.Println(err.Error())
		fmt.Fprint(w, "Something wen to shit")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(stars)
}

func (a *api) handlePostStar(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userIdKey).(string)

	decoder := json.NewDecoder(r.Body)
	var star repository.Star
	err := decoder.Decode(&star)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Unable to read body")
		return
	}

	star.UserID = userId

	err = a.repository.CreateStar(&star)

	if err != nil {
		log.Println(err.Error())
		fmt.Fprint(w, "Something went to shit")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(&star)
}
