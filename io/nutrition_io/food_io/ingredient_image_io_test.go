package food_io

import (
	"afcmo/domain/nutrition"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateIngredientImage(t *testing.T) {
	ingredientImage := nutrition.IngredientImage{"","00001","00002"}
	ingredient,err := CreateIngredientImage(ingredientImage)
	assert.Nil(t, err)
	fmt.Println("result: ",ingredient)
}
