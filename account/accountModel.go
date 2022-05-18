package account

import (
	"GoSchool-Assignment4/data"
	"GoSchool-Assignment4/userp"
	"fmt"
	"net/http"
	"strings"
)

// EditUser allows the user to edit their username.
func EditUser(res http.ResponseWriter, req *http.Request) {
	if !userp.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	user := userp.GetUser(res, req)

	oldName := user.ProfileName

	if req.Method == http.MethodPost {
		username := req.FormValue("username")

		if username == "" {
			data.Error.Printf("User %v tried to edit user without filling out all required fields.\n", user.ProfileName)
			http.Error(res, "All fields must be filled out.", http.StatusUnauthorized)
			return

		} else {

			// Check if username contains spaces. Username should not contain spaces.
			if strings.ContainsAny(username, " ") {
				data.Error.Printf("Unsuccessful edit user. User tried to enter new username with spaces: %v\n", username)
				http.Error(res, "Username cannot contain spaces.", http.StatusForbidden)
			}

			// Username should not be case sensitive, so before working with input, convert to lower case.
			username = strings.ToLower(username)

			// check if username exists
			exists := false
			for k := range data.Users {
				if k == username {
					exists = true
				}
			}

			if !exists {
				// new entry into users map
				data.Users[username] = &(*(data.Users[user.ProfileName]))
				fmt.Println(data.Users)
				// change profileName
				data.Users[username].ProfileName = username

				// update mapSessions
				cookie, _ := req.Cookie("FriendTrackerCookie")
				data.MapSessions[cookie.Value] = username

				// delete old userProfile
				delete(data.Users, oldName)

				data.Info.Printf("User %v successfully changed their username.\n", user.ProfileName)
				http.Redirect(res, req, "/accountmanagement/", http.StatusSeeOther)

			} else {
				data.Error.Printf("Unsuccessful edit user by %v. New username already exists.\n", username)
				http.Error(res, "Username already exists.", http.StatusUnauthorized)
				return
			}
		}

	}

	data.Tpl.ExecuteTemplate(res, "edituser.gohtml", nil)
}

// DeleteUser allows the user to delete their profile and returns them to the login page.
func DeleteUser(res http.ResponseWriter, req *http.Request) {
	if !userp.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	user := userp.GetUser(res, req)

	oldName := user.ProfileName

	delete(data.Users, oldName)

	data.Info.Printf("User %v successfully deleted their user profile and logged out.\n", user.ProfileName)

	http.Redirect(res, req, "/logout/", http.StatusSeeOther)
}
