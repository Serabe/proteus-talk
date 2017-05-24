package gomad

import (
	"errors"
)

//proteus:generate
type Pizza struct {
	ID          int32
	Ingredients []*IngredientSelection
}

type IngredientSelection struct {
	Ingredient Ingredient
	Portions   int32
}

//proteus:generate
type Ingredient int

const (
	Pepperoni Ingredient = iota
	Pineapple
	Ham
	Bacon
	Chicken
	Lamb
	NoFood
)

var _ = Pineapple
var _ = Ham
var _ = Bacon
var _ = Chicken
var _ = Lamb
var _ = NoFood

//proteus:generate
func AddIngredient(p *Pizza, ing Ingredient, portions int32) *Pizza {
	for _, i := range p.Ingredients {
		if i.Ingredient == ing {
			i.Portions += portions
			return p
		}
	}

	p.Ingredients = append(p.Ingredients, &IngredientSelection{
		Ingredient: ing,
		Portions:   portions,
	})

	return p
}

//proteus:generate
func RemoveIngredient(p *Pizza, ing Ingredient, portions int32) (*Pizza, error) {
	for idx, i := range p.Ingredients {
		if i.Ingredient == ing {
			if i.Portions < portions {
				return nil, errors.New("there were not enough portions to remove")
			}

			i.Portions -= portions

			if i.Portions == 0 {
				p.Ingredients = append(p.Ingredients[:idx], p.Ingredients[idx+1:]...)
			}

			return p, nil
		}
	}

	return nil, errors.New("there was no ingredient")
}

//proteus:generate
func NutritionalReport(p *Pizza) string {
	return "Who cares? Just enjoy it, Javier!"
}
