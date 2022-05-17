// Package main starts the server and handles routing.
package main

import (
	"fmt"
	"net/http"

	"GoSchool-Assignment4/account"
	"GoSchool-Assignment4/friend"
	"GoSchool-Assignment4/search"
	"GoSchool-Assignment4/userp"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", userp.Login)
	http.HandleFunc("/signup/", userp.Signup)

	http.HandleFunc("/friends/", friend.FriendsControl)
	http.HandleFunc("/addfriend/", friend.AddFriendToList)
	http.HandleFunc("/addgroup/", friend.AddGroup)
	http.HandleFunc("/editgroup/", friend.EditExistingGroup)
	http.HandleFunc("/deletegroup/", friend.DeleteGroup)

	http.HandleFunc("/accountmanagement/", account.AccountManagement)
	http.HandleFunc("/edituser/", account.EditUser)
	http.HandleFunc("/deleteuser/", account.DeleteUser)

	http.HandleFunc("/search/", search.SearchControl)
	http.HandleFunc("/deletefriend/", search.DeleteFriend)
	http.HandleFunc("/editfrienddetails/", search.EditFriendDetails)

	http.HandleFunc("/logout/", userp.Logout)

	http.ListenAndServe(":5221", nil)
}
