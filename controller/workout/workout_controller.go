package workout

import (
	"afcmo/config"
	"afcmo/controller/workout/exercise"
	exercise3 "afcmo/controller/workout/plan"
	nutrition "afcmo/controller/workout/subscription"
	exercise2 "afcmo/controller/workout/type"
	exercise4 "afcmo/controller/workout/user"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

func WorkoutController(app *config.Env) http.Handler {
	mux := chi.NewMux()

	mux.Handle("/", Homeworkout(app))
	//mux.Handle("/{nutritionId}", Nutrition(app))
	mux.Mount("/exercise", exercise.ExerciseController(app))
	mux.Mount("/type", exercise2.WorkoutTypeController(app))
	mux.Mount("/plan", exercise3.WorkoutPlanController(app))
	mux.Mount("/user", exercise4.WorkoutUserController(app))
	mux.Mount("/subscription", nutrition.WorkoutSubscriptionController(app))
	return mux
}
func Homeworkout(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "workout/workout-home-page.html",
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
