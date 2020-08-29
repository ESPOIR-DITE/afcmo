package nutrition

import (
	"afcmo/config"
	"afcmo/controller/nutrition/food"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

//mux := chi.NewMux()
//
//mux.Mount("/api", api.InstitutionAPI(app))
//
//return mux
func NutritionController(app *config.Env) http.Handler {
	mux := chi.NewMux()

	mux.Handle("/", HomeNutrition(app))
	mux.Handle("/{nutritionId}", Nutrition(app))
	mux.Mount("/food", food.Controller(app))
	mux.Mount("/user", NutritionUserController(app))
	mux.Mount("/video", NutritionVideoController(app))
	mux.Mount("/type", NutritionTypeController(app))
	mux.Mount("/image", NutritionPlaneController(app))
	mux.Mount("/plan", NutritionPlaneController(app))
	mux.Mount("/subscription", NutritionSubscriptionController(app))

	return mux
}

func Nutrition(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "nutrition/food-details.html",
			app.Path + "template/header.html",
			app.Path + "template/footer.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func HomeNutrition(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "nutrition/nutrition-home-page.html",
			app.Path + "template/navbar.html",
			app.Path + "template/topbar.html",
			app.Path + "template/cards.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
