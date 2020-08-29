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

func IngredientController(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", IngredientHomeHandler(app))
	r.Get("/edit/{ingredientId}", IngredientEditHandler(app))
	r.Get("/delete/{ingredientId}", IngredientDeleteHandler(app))
	r.Post("/create",CreateIngredientHandler(app))
	r.Post("/update",UpdateIngredientHandler(app))
	return r
}

func IngredientDeleteHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var success,info,warning,error bool
		ingredientId:=chi.URLParam(r,"ingredientId")

		//checking the ingredient
		_,err := food_io.ReadIngredient(ingredientId)
		if err != nil {
			error = true
			fmt.Println(err," error reading ingredient")
			//Adding notification to the session
			util.CreateNotificationSession(app,r,util.Notification{success,info,warning,error})
			http.Redirect(w, r, "/nutrition/food/ingredient", 301)
			return
		}

		//Deleting the ingredient
		_,err = food_io.DeleteIngredient(ingredientId)
		if err != nil {
			error = true
			fmt.Println(err," error deleting ingredient")
			//Adding notification to the session
			util.CreateNotificationSession(app,r,util.Notification{success,info,warning,error})
			http.Redirect(w, r, "/nutrition/food/ingredient", 301)
			return
		}

		//checking ingredient image
		ingredientImage,err := food_io.ReadIngredientImageOf(ingredientId)
		if err != nil {
			fmt.Println(err," error reading ingredient image, This error may occur if the ingredient had not an image")
		}else {
			_,err := food_io.DeleteIngredientImage(ingredientImage.Id)
			if err != nil {
				warning = true
				fmt.Println(err," error deleting ingredient image")
			}else{
				_,err := image_io.DeleteImage(ingredientImage.ImageId)
				if err != nil {
					warning = true
					fmt.Println(err," error deleting image")
				}else {
					info= true
				}
			}
		}
		util.CreateNotificationSession(app,r,util.Notification{success,info,warning,error})
		http.Redirect(w, r, "/nutrition/food/ingredient", 301)
	}
}

func UpdateIngredientHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var isFileToUpdate bool
		var content []byte
		var success,info,warning,error bool
		r.ParseForm()
		file, _, err := r.FormFile("image")
		name:=r.PostFormValue("name")
		brief:=r.PostFormValue("brief")
		ingredientId:=r.PostFormValue("ingredientId")
		imageId:=r.PostFormValue("imageId")
		//need to putt a validation

		//image:= []io.Reader{file}
		if err==nil{
		reader := bufio.NewReader(file)
		content, _ = ioutil.ReadAll(reader)
		isFileToUpdate = true
		}

		fmt.Println("creating an ingredient", content," <image","name>",name,"brief> ",brief)

		if brief!=""&&name!=""&&ingredientId!=""&&imageId!=""{
			ingredientObject:= nutrition.Ingredient{ingredientId,name,brief}
			_,err := food_io.UpdateIngredient(ingredientObject)
			if err != nil {
				warning = true
				fmt.Println(err," error updating ingredient")
				util.CreateNotificationSession(app,r,util.Notification{success,info,warning,error})
				http.Redirect(w, r, "/nutrition/food/ingredient", 301)

			}else if isFileToUpdate == true{
				fmt.Println("updating the image")
				imageObject := image2.Images{imageId,content,brief}
				_,err := image_io.UpdateImage(imageObject)
				if err != nil {
					warning = true
					fmt.Println(err," error updating image")
					util.CreateNotificationSession(app,r,util.Notification{success,info,warning,error})
					http.Redirect(w, r, "/nutrition/food/ingredient", 301)
				}else {
						//If all goes well as expected
						success = true
				}
			}
		}
		//Adding notification to the session
		util.CreateNotificationSession(app,r,util.Notification{success,info,warning,error})
		http.Redirect(w, r, "/nutrition/food/ingredient/edit/"+ingredientId, 301)
	}
}

func IngredientEditHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var myImage image2.ImageHelper
		var success,info,warning,error bool
		ingredientId := chi.URLParam(r,"ingredientId")

		//check the ingredient
		ingredient,err := food_io.ReadIngredient(ingredientId)
		if err != nil {
			warning = true
			fmt.Println(err," error reading ingredient")
			//Adding notification to the session
			util.CreateNotificationSession(app,r,util.Notification{success,info,warning,error})
			http.Redirect(w, r, "/nutrition/food/ingredient", 301)
			return
		}
		//reading the image
		ingredientImage,err := food_io.ReadIngredientImageOf(ingredientId)
		if err != nil {
			fmt.Println(err, " error reading ingredient")
		}else {
				image,err:= image_io.ReadImage(ingredientImage.ImageId)
				if err != nil {
					fmt.Println(err, " error reading image")
				}else {
					myImage = image2.ImageHelper{image.Id,util.ConvertToString(image.Image),image.Description}
				}

		}

		notification := util.ReadNotificationSession(app,r)
		//Delete everything in the Session
		util.DeleteNotificationSession(app,r)

		type PageData struct {
			Ingredient nutrition.Ingredient
			Notification util.Notification
			MyImage image2.ImageHelper
		}
		data := PageData{ingredient,notification,myImage}
		files := []string{
			app.Path + "nutrition/ingredient/ingredient-edit-page.html",
			app.Path + "template/navbar.html",
			app.Path + "template/topbar.html",
			app.Path + "template/cards.html",
			app.Path + "template/notification.html",
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

func CreateIngredientHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var success,info,warning,error bool
		r.ParseForm()
		file, _, err := r.FormFile("image")
		name:=r.PostFormValue("name")
		brief:=r.PostFormValue("brief")
		//need to putt a validation
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		//image:= []io.Reader{file}
		reader := bufio.NewReader(file)
		content, _ := ioutil.ReadAll(reader)

		fmt.Println("creating an ingredient", content," <image","name>",name,"brief> ",brief)

		if brief!=""&&name!=""{
			ingredientObject:= nutrition.Ingredient{"",name,brief}
			ingredient,err := food_io.CreateIngredient(ingredientObject)
			if err != nil {
				error = true
				fmt.Println(err," error creating ingredient")
			}else{
				imageObject := image2.Images{"",content,brief}
				imageReturn,err := image_io.CreateImage(imageObject)
				if err != nil {
					error = true
					fmt.Println(err," error creating image")
					_,err := food_io.DeleteIngredient(ingredient.Id)
					if err != nil {
						fmt.Println(err," error deleting ingredient")
					}
				}else {
					imageIngredeintImage:=nutrition.IngredientImage{"",ingredient.Id,imageReturn.Id}
					_,err := food_io.CreateIngredientImage(imageIngredeintImage)
					if err != nil {
						_,err := food_io.DeleteIngredient(ingredient.Id)
						if err != nil {
							fmt.Println(err," error deleting ingredient")
						}
						//Deleting image
						_,errx := image_io.DeleteImage(imageReturn.Id)
						if errx != nil {
							fmt.Println(err," error deleting ingredient")
						}

						error = true
						fmt.Println(err," error creating IngredientImage")
					}else {
						//If all goes well as expected
						success = true
					}
				}
			}
		}
		//Adding notification to the session
		util.CreateNotificationSession(app,r,util.Notification{success,info,warning,error})
		http.Redirect(w, r, "/nutrition/food/ingredient", 301)
	}
}

func IngredientHomeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var success,info,warning,error bool
		//Collecting every notification in the session
		notification := util.ReadNotificationSession(app,r)
		//Delete everything in the Session
		util.DeleteNotificationSession(app,r)

		//Read ingredients
		ingredients,err := food_io.ReadIngredientAll()
		if err != nil {
			fmt.Println(err," error reading ingredients")
		}

		type PageData struct {
			Notification util.Notification
			Ingredients []nutrition.Ingredient
		}

		data := PageData{notification,ingredients}
		files := []string{
			app.Path + "nutrition/ingredient/ingredient-home-page.html",
			app.Path + "template/navbar.html",
			app.Path + "template/topbar.html",
			app.Path + "template/cards.html",
			app.Path + "template/notification.html",
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
