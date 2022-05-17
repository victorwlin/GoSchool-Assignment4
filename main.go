package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	tpl         *template.Template
	mapSessions = map[string]string{}

	Info  *log.Logger
	Error *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Lmicroseconds|log.Llongfile)

	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", login)
	http.HandleFunc("/signup/", signup)

	http.HandleFunc("/friends/", friendsControl)
	http.HandleFunc("/addfriend/", addFriendToList)
	http.HandleFunc("/addgroup/", addGroup)
	http.HandleFunc("/editgroup/", editExistingGroup)
	http.HandleFunc("/deletegroup/", deleteGroup)

	http.HandleFunc("/accountmanagement/", accountManagement)
	http.HandleFunc("/edituser/", editUser)
	http.HandleFunc("/deleteuser/", deleteUser)

	http.HandleFunc("/search/", searchControl)
	http.HandleFunc("/deletefriend/", deleteFriend)
	http.HandleFunc("/editfrienddetails/", editFriendDetails)

	http.HandleFunc("/logout/", logout)

	http.ListenAndServe(":5221", nil)
}
