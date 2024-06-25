package helper

import "net/http"

func HttpErrorHandler(w http.ResponseWriter,err error){
	http.Error(w,err.Error(),http.StatusInternalServerError)
	return
}