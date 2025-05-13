package recipes

import (
	"log"
)

type memStore struct {
}

func NewMemStore() *memStore {
	return &memStore{}
}

func (m *memStore) Add(name string, recipe Recipe) error {
	log.Printf("Added recipe %s\n", name)

	return nil
}

func (m *memStore) Get(name string) (Recipe, error) {
	panic("implement me")
}

func (m *memStore) Update(name string, recipe Recipe) error {
	panic("implement me")
}

func (m *memStore) List() (map[string]Recipe, error) {
	recipes := map[string]Recipe{
		"pasta": {
			Name: "Spaghetti Carbonara",
			Ingredients: []Ingredient{
				{Name: "tomatoes"},
				{Name: "onion"},
				{Name: "garlic"},
				{Name: "oil"},
			},
		},
		"soup": {
			Name: "Tomato Soup",
			Ingredients: []Ingredient{
				{Name: "tomatoes"},
				{Name: "onion"},
				{Name: "garlic"},
				{Name: "vegetable stock"},
				{Name: "cream"},
			},
		},
	}

	return recipes, nil
}

func (m *memStore) Remove(name string) error {
	panic("implement me")
}
