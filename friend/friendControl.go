// Package friend controls the functions of the friends menu.
package friend

import (
	"net/http"

	"GoSchool-Assignment4/data"
	"GoSchool-Assignment4/search"
	"GoSchool-Assignment4/userp"
)

// FriendsControl displays the friends menu template and allows the user to sort the friends list.
func FriendsControl(res http.ResponseWriter, req *http.Request) {
	if !userp.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	user := userp.GetUser(res, req)

	sortType := req.FormValue("sort")
	sortControl(user.Friends, sortType)

	if req.Method == http.MethodPost {
		friend := req.FormValue("friend")

		// check if fields have been filled out
		if friend == "" {

			http.Error(res, "Search field must not be blank.", http.StatusUnauthorized)
			return

		} else {

			if user.Friends.Head != nil { // only do search if there is a friends list
				// check if friend exists
				friendNode, _ := search.SeqSearch(user.Friends, friend)
				if friendNode == nil {
					http.Error(res, "Friend does not exist.", http.StatusUnauthorized)
					return
				} else {
					http.Redirect(res, req, "/search/?friend="+friend, http.StatusSeeOther)
				}
			} else {
				data.Info.Println("User tried to search for a friend in an empty friend list.")
				http.Error(res, "No friends to search for.", http.StatusUnauthorized)
				return
			}

		}
	}

	data.Tpl.ExecuteTemplate(res, "friends.gohtml", user.Friends.MakeFriendsSlice())
}
