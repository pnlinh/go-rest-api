package main

import (
	"encoding/json"
	"errors"
	"github.com/gosimple/slug"
	"go-rest-api/pkg/recipes"
	"net/http"
	"regexp"
)

func main() {
	// Create the Store and Recipe Handler
	store := recipes.NewMemStore()
	recipesHandler := NewRecipesHandler(store)

	mux := http.NewServeMux()

	mux.Handle("/", &homeHandler{})
	mux.Handle("/recipes", recipesHandler)
	mux.Handle("/recipes/", recipesHandler)

	http.ListenAndServe("localhost:8080", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

var (
	RecipeRe       = regexp.MustCompile(`^/recipes/*$`)
	RecipeReWithID = regexp.MustCompile(`^/recipes/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)

func (h *RecipesHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe recipes.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	// Convert the name of the recipe into URL friendly string
	resourceID := slug.Make(recipe.Name)
	// Call the store to add the recipe
	if err := h.store.Add(resourceID, recipe); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	// Set the status code to 201
	w.WriteHeader(http.StatusCreated)
}
func (h *RecipesHandler) ListRecipes(w http.ResponseWriter, r *http.Request) {
	resources, _ := h.store.List()
	jsonBytes, err := json.Marshal(resources)

	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
func (h *RecipesHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	matches := RecipeReWithID.FindStringSubmatch(r.URL.Path)

	if len(matches) < 2 {
		InternalServerErrorHandler(w, r)
		return
	}

	// Retrieve recipe from the store
	recipe, err := h.store.Get(matches[1])
	if err != nil {
		// Special case of NotFound Error
		if errors.Is(err, recipes.NotFoundErr) {
			NotFoundHandler(w, r)
			return
		}

		// Every other error
		InternalServerErrorHandler(w, r)
		return
	}

	// Convert the struct into JSON payload
	jsonBytes, err := json.Marshal(recipe)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	// Write the results
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
func (h *RecipesHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {}
func (h *RecipesHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {}

func (h *RecipesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && RecipeRe.MatchString(r.URL.Path):
		h.CreateRecipe(w, r)
		return
	case r.Method == http.MethodGet && RecipeRe.MatchString(r.URL.Path):
		h.ListRecipes(w, r)
		return
	case r.Method == http.MethodGet && RecipeReWithID.MatchString(r.URL.Path):
		h.GetRecipe(w, r)
		return
	case r.Method == http.MethodPut && RecipeReWithID.MatchString(r.URL.Path):
		h.UpdateRecipe(w, r)
		return
	case r.Method == http.MethodDelete && RecipeReWithID.MatchString(r.URL.Path):
		h.DeleteRecipe(w, r)
		return
	default:
		return
	}
}

type recipeStore interface {
	Add(name string, recipe recipes.Recipe) error
	Get(name string) (recipes.Recipe, error)
	Update(name string, recipe recipes.Recipe) error
	List() (map[string]recipes.Recipe, error)
	Remove(name string) error
}

// RecipesHandler implements http.Handler and dispatches requests to the store
type RecipesHandler struct {
	store recipeStore
}

func NewRecipesHandler(s recipeStore) *RecipesHandler {
	return &RecipesHandler{
		store: s,
	}
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}
