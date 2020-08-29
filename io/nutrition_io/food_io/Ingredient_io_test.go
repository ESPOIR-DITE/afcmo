package food_io

import (
	"afcmo/domain/nutrition"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateIngredient(t *testing.T) {
	ingredientObject := nutrition.Ingredient{"","cucamba","green"}
	ingredient,err := CreateIngredient(ingredientObject)
	assert.Nil(t, err)
	fmt.Println("result: ",ingredient)
}
