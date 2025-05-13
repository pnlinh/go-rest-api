package recipes

import (
	"log"
)

type MemStore struct {
}

func NewMemStore() *MemStore {
	return &MemStore{}
}

func (m *MemStore) Add(name string, recipe Recipe) error {
	log.Printf("Added recipe %s\n", name)

	return nil
}

func (m *MemStore) Get(name string) (Recipe, error) {
	panic("implement me")
}

func (m *MemStore) Update(name string, recipe Recipe) error {
	panic("implement me")
}

func (m *MemStore) List() (map[string]Recipe, error) {
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

func (m *MemStore) Remove(name string) error {
	panic("implement me")
}
