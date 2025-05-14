package recipes

import (
	"errors"
	"log"
)

var NotFoundErr = errors.New("not found")

type memStore struct {
	recipes map[string]Recipe
}

func NewMemStore() *memStore {
	recipes := make(map[string]Recipe)

	return &memStore{recipes: recipes}
}

func (m *memStore) Add(name string, recipe Recipe) error {
	m.recipes[name] = recipe
	log.Printf("Added recipe %s\n", name)

	return nil
}

func (m *memStore) Get(name string) (Recipe, error) {
	if recipe, ok := m.recipes[name]; ok {
		return recipe, nil
	}

	return Recipe{}, NotFoundErr
}

func (m *memStore) Update(name string, recipe Recipe) error {
	if _, ok := m.recipes[name]; ok {
		m.recipes[name] = recipe

		return nil
	}

	return NotFoundErr
}

func (m *memStore) List() (map[string]Recipe, error) {
	return m.recipes, nil
}

func (m *memStore) Remove(name string) error {
	delete(m.recipes, name)

	return nil
}
