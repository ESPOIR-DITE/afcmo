package exercise

import (
	"afcmo/config"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

func WorkoutPlanController(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", WorkoutPlanHomeHandler(app))
	return r
}

func WorkoutPlanHomeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "workout/plan/workout-plan-home-page.html",
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
