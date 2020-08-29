package controller

/**
* This is the main controller
* Every requests passe first here
* The purpose of this class is to direct the request(URL) coming from html to the respective controller classes
**/
import (
	"afcmo/config"
	"afcmo/controller/nutrition"
	"afcmo/controller/user"
	"afcmo/controller/workout"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"html/template"
	"net/http"
)

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(env.Session.LoadAndSave)

	mux.Handle("/", HomeHandler(env))
	mux.Handle("/about", AboutHandler(env))

	mux.Mount("/nutrition", nutrition.NutritionController(env))
	mux.Mount("/workout", workout.WorkoutController(env))
	mux.Mount("/user", user.UserController(env))
	//mux.Mount("/gallery", gallery.GalleryController(env))
	//mux.Mount("/workout", workout2.WorkoutController(env))
	//mux.Mount("/nutrition", nutrition2.Controller(env))

	fileServer := http.FileServer(http.Dir("view/assets/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/assets/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	return mux
}

func AboutHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "index.html",
			app.Path + "template/navbar.html",
			app.Path + "template/topbar.html",
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
func HomeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "index.html",
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
