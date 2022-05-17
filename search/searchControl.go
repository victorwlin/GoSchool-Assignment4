// Package search controls all aspects of the search menu.
package search

import (
	"net/http"
	"time"

	"GoSchool-Assignment4/data"
	"GoSchool-Assignment4/userp"
)

// SearchControl displays the search menu and allows user to add a new date of last contact to the current friend.
func SearchControl(res http.ResponseWriter, req *http.Request) {
	if !userp.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	user := userp.GetUser(res, req)

	friend := req.FormValue("friend")
	friendNode, _ := SeqSearch(user.Friends, friend)

	if req.Method == http.MethodPost {
		newLastContact := req.FormValue("newLastContact")

		// check if fields have been filled out
		if newLastContact == "" {

			http.Error(res, "Field must not be blank.", http.StatusUnauthorized)
			return

		} else {

			date, _ := time.Parse("2006-01-02", newLastContact)
			friendNode.LastContact.Push(date)

		}
	}

	data.Tpl.ExecuteTemplate(res, "search.gohtml", user.Friends.MakeSearchStruct(friendNode, user))
}
