// Package main starts the server and handles routing.
package main

import (
	"fmt"
	"net/http"

	"GoSchool-Assignment4/account"
	"GoSchool-Assignment4/friend"
	"GoSchool-Assignment4/search"
	"GoSchool-Assignment4/userp"

	"github.com/gorilla/mux"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	r := mux.NewRouter()

	r.Handle("/favicon.ico", http.NotFoundHandler())

	r.HandleFunc("/", userp.Login)
	r.HandleFunc("/signup/", userp.Signup)

	r.HandleFunc("/friends/", friend.FriendsControl)
	r.HandleFunc("/addfriend/", friend.AddFriendToList)
	r.HandleFunc("/addgroup/", friend.AddGroup)
	r.HandleFunc("/editgroup/", friend.EditExistingGroup)
	r.HandleFunc("/deletegroup/", friend.DeleteGroup)

	r.HandleFunc("/accountmanagement/", account.AccountManagement)
	r.HandleFunc("/edituser/", account.EditUser)
	r.HandleFunc("/deleteuser/", account.DeleteUser)

	r.HandleFunc("/search/", search.SearchControl)
	r.HandleFunc("/deletefriend/", search.DeleteFriend)
	r.HandleFunc("/editfrienddetails/", search.EditFriendDetails)

	r.HandleFunc("/logout/", userp.Logout)

	http.Handle("/", r)

	http.ListenAndServe(":5221", nil)
}
