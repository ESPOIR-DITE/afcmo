package food

import (
	"afcmo/config"
	"afcmo/controller/util"
	image2 "afcmo/domain/image"
	"afcmo/domain/nutrition"
	"afcmo/io/image_io"
	"afcmo/io/nutrition_io/food_io"
	"bufio"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io/ioutil"
	"net/http"
)

func Controller(app *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Handle("/", FoodHomeHandler(app))
	mux.Post("/create", FoodCreateHandler(app))
	mux.Mount("/ingredient", IngredientController(app))
	return mux
}

func FoodCreateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var success,info,warning,error bool
		r.ParseForm()
		file, _, err := r.FormFile("image")
		foodName := r.PostFormValue("foodName")
		brief := r.PostFormValue("brief")
		ingredientIds := r.Form["ingredientIds"]
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		//image:= []io.Reader{file}
		reader := bufio.NewReader(file)
		content, _ := ioutil.ReadAll(reader)

		//Creating food
		if foodName!=""&&brief!=""&&ingredientIds!=nil{
			foodObject := nutrition.Food{"",foodName,brief}
			food,err := food_io.CreateFood(foodObject)
			if err != nil {
				error = true
				fmt.Println(err, "error creating food")
				util.CreateNotificationSession(app,r,util.Notification{success,info,warning,error})
				http.Redirect(w, r, "/nutrition/food", 301)
			}

			//Creating foodImgrdient
			for _,ingredientId := range ingredientIds{

				foodIngredientObject := nutrition.FoodIngredient{"",food.Id,ingredientId,""}

				_,err := food_io.CreateFoodIngredient(foodIngredientObject)
				if err != nil {
					error = true
					fmt.Println(err, " error creating food image")
				}

			}
			imageObject := image2.Images{"",content,brief}
			imageReturn,err := image_io.CreateImage(imageObject)
			if err != nil {
				error = true
				fmt.Println(err," error creating image")
				_,err :=food_io.DeleteFood(food.Id)
				if err != nil {
					fmt.Println(err," error deleting food")
				}
			}else {
				foodImage := nutrition.FoodImage{"",food.Id,imageReturn.Id}
				foodImage,err :=food_io.CreateFoodImage(foodImage)
				if err != nil {
					fmt.Println(err," error create foodImage")
					_,err :=food_io.DeleteFood(food.Id)
					if err != nil {
						fmt.Println(err," error deleting food")
					}
				}else {
					//If all goes well as expected
					success = true
				}
			}
		}
		//Adding notification to the session
		util.CreateNotificationSession(app,r,util.Notification{success,info,warning,error})
		http.Redirect(w, r, "/nutrition/food", 301)

	}
}

func FoodHomeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//Collecting every notification in the session
		notification := util.ReadNotificationSession(app,r)
		//Delete everything in the Session
		util.DeleteNotificationSession(app,r)
		// Read All the foods

		foods,err := food_io.ReadFoodAll()
		if err!=nil{
			fmt.Println("error reading foods")
		}
		ingredients,err := food_io.ReadIngredientAll()
		if err!=nil{
			fmt.Println("error reading ingredients")
		}
		type PageData struct {
			Foods []nutrition.Food
			Notification util.Notification
			Ingredients []nutrition.Ingredient
		}
		data := PageData{foods,notification,ingredients}
		files := []string{
			app.Path + "nutrition/food/food-home-page.html",
			app.Path + "template/navbar.html",
			app.Path + "template/topbar.html",
			app.Path + "template/cards.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
