package recipes

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemStore_Add(t *testing.T) {
	tests := []struct {
		name        string
		recipes     map[string]Recipe
		inputName   string
		inputRecipe Recipe
		expectedErr error
	}{
		{
			name:        "Add a new recipe",
			recipes:     map[string]Recipe{},
			inputName:   "Pancakes",
			inputRecipe: Recipe{Name: "Pancakes", Ingredients: []Ingredient{{Name: "Flour"}}},
			expectedErr: nil,
		},
		{
			name:        "Add recipe with empty name",
			recipes:     map[string]Recipe{},
			inputName:   "",
			inputRecipe: Recipe{Name: "EmptyName", Ingredients: []Ingredient{{Name: "Milk"}}},
			expectedErr: nil,
		},
		{
			name:        "Add recipe with empty ingredients",
			recipes:     map[string]Recipe{},
			inputName:   "EmptyIngredients",
			inputRecipe: Recipe{Name: "EmptyIngredients", Ingredients: []Ingredient{}},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &memStore{recipes: tt.recipes}
			err := store.Add(tt.inputName, tt.inputRecipe)

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Add() error = %v, expectedErr = %v", err, tt.expectedErr)
			}
		})
	}
}

func TestMemStore_Get(t *testing.T) {
	tests := []struct {
		name        string
		recipes     map[string]Recipe
		inputName   string
		expectedErr error
	}{
		{
			name: "Find pasta recipe",
			recipes: map[string]Recipe{
				"Pasta": {Name: "Pasta", Ingredients: []Ingredient{{Name: "Flour"}}},
			},
			inputName:   "Pasta",
			expectedErr: nil,
		},
		{
			name:        "Find recipe with empty name",
			recipes:     map[string]Recipe{},
			inputName:   "",
			expectedErr: NotFoundErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &memStore{recipes: tt.recipes}
			_, err := store.Get(tt.inputName)

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Get() error = %v, expectedErr = %v", err, tt.expectedErr)
			}
		})
	}
}

func TestMemStore_List(t *testing.T) {
	tests := []struct {
		name                string
		recipes             map[string]Recipe
		inputName           string
		expectedErr         error
		expectedTotalRecord int
	}{
		{
			name: "List recipe",
			recipes: map[string]Recipe{
				"pasta": {
					Name:        "Pasta",
					Ingredients: []Ingredient{{Name: "Flour"}},
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
			},
			expectedErr:         nil,
			expectedTotalRecord: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &memStore{recipes: tt.recipes}
			_, err := store.List()

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("List() error = %v, expectedErr = %v", err, tt.expectedErr)
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedTotalRecord, len(tt.recipes))
		})
	}
}

func TestMemStore_Update(t *testing.T) {
	tests := []struct {
		name        string
		recipes     map[string]Recipe
		inputName   string
		inputRecipe Recipe
		expectedErr error
	}{
		{
			name: "Update pasta recipe",
			recipes: map[string]Recipe{
				"pasta": {Name: "Pasta", Ingredients: []Ingredient{{Name: "Flour"}}},
			},
			inputName:   "pasta",
			inputRecipe: Recipe{Name: "Pasta updated", Ingredients: []Ingredient{{Name: "Flour"}}},
			expectedErr: nil,
		},
		{
			name: "Update pasta recipe",
			recipes: map[string]Recipe{
				"pasta": {Name: "Pasta", Ingredients: []Ingredient{{Name: "Flour"}}},
			},
			inputName:   "pasta1",
			inputRecipe: Recipe{Name: "Pasta updated", Ingredients: []Ingredient{{Name: "Flour"}}},
			expectedErr: NotFoundErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &memStore{recipes: tt.recipes}
			err := store.Update(tt.inputName, tt.inputRecipe)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.inputRecipe, tt.recipes[tt.inputName])
			}
		})
	}
}

func TestMemStore_Remove(t *testing.T) {
	tests := []struct {
		name                string
		recipes             map[string]Recipe
		inputName           string
		inputRecipe         Recipe
		expectedErr         error
		expectedTotalRecord int
	}{
		{
			name: "Remove pasta recipe",
			recipes: map[string]Recipe{
				"pasta": {Name: "Pasta", Ingredients: []Ingredient{{Name: "Flour"}}},
			},
			inputName:   "pasta",
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &memStore{recipes: tt.recipes}
			err := store.Remove(tt.inputName)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, tt.recipes, 0)
			}
		})
	}
}
