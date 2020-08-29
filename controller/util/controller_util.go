package util

import (
	"afcmo/config"
	"net/http"
)

const (
	Success = "success"
	Info = "info"
	Warning = "warning"
	Error = "error"
)

type Notification struct {
	Success bool
	Info    bool
	Warning bool
	Error   bool
}
func ConvertToString(bytes []byte) string{
	return string(bytes)
}
func GetNotification(Success, Info, Warning, Error bool)Notification{
	return Notification{Success,Info,Warning,Error}
}
//This method help create a new session
func CreateNotificationSession(app *config.Env, r *http.Request,notification Notification){
	app.Session.Put(r.Context(),Success,notification.Success)
	app.Session.Put(r.Context(),Info,notification.Info)
	app.Session.Put(r.Context(),Warning,notification.Warning)
	app.Session.Put(r.Context(),Error,notification.Error)
}
func ReadNotificationSession(app *config.Env, r *http.Request)Notification{
	var success,info,warning,error bool
	success = app.Session.GetBool(r.Context(),Success)
	info = app.Session.GetBool(r.Context(),Info)
	warning =app.Session.GetBool(r.Context(),Warning)
	error = app.Session.GetBool(r.Context(),Error)
return Notification{success,info,warning,error}
}
func DeleteNotificationSession(app *config.Env, r *http.Request){
	app.Session.Remove(r.Context(),Success)
	app.Session.Remove(r.Context(),Info)
	app.Session.Remove(r.Context(),Warning)
	app.Session.Remove(r.Context(),Error)
}