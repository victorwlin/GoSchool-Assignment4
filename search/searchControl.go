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
			data.Error.Printf("User %v submitted a blank search.\n", user.ProfileName)
			http.Error(res, "Field must not be blank.", http.StatusUnauthorized)
			return

		} else {

			date, _ := time.Parse("2006-01-02", newLastContact)

			if date.Before(user.Friends.Head.LastContact.Top.Date) {
				data.Error.Printf("User %v tried to enter a new date of last contact that is earlier than the current date of last contact.\n", user.ProfileName)
				http.Error(res, "New date of last contact must have happened earlier than current date of last contact.", http.StatusUnauthorized)
				return
			}

			friendNode.LastContact.Push(date)
			data.Info.Printf("User %v successfully added a date of last contact to one of their friends.\n", user.ProfileName)
		}
	}

	data.Tpl.ExecuteTemplate(res, "search.gohtml", user.Friends.MakeSearchStruct(friendNode, user))
}
