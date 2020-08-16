package config

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"runtime/debug"
)

type Env struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Path     string
	Session  *scs.SessionManager
}

type templateData struct {
	Title string
	Data  map[string]string
}

func (app *Env) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Env) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}

func (app *Env) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
