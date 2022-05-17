package userp

import (
	"net/http"

	"GoSchool-Assignment4/data"
)

// GetUser retrieves the user profile of the current user based on the session cookie.
func GetUser(res http.ResponseWriter, req *http.Request) (user *(data.UserProfile)) {
	// get current session cookie
	cookie, err := req.Cookie("FriendTrackerCookie")

	// if cookie doesn't exist, redirect to login
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if _, ok := data.MapSessions[cookie.Value]; ok {
		username := data.MapSessions[cookie.Value]

		user = data.Users[username]
	}

	return user
}

// AlreadyLoggedIn checks if the current user is currently logged in.
func AlreadyLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("FriendTrackerCookie")
	if err != nil {
		return false
	}

	username := data.MapSessions[cookie.Value]
	_, ok := data.Users[username]

	return ok
}

// Logout allows the user to log out and delete the cookie as well as the entry in the MapSession map.
func Logout(res http.ResponseWriter, req *http.Request) {
	cookie, _ := req.Cookie("FriendTrackerCookie")

	// delete session
	delete(data.MapSessions, cookie.Value)

	// remove cookie
	cookie = &http.Cookie{
		Name:   "FriendTrackerCookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(res, cookie)

	http.Redirect(res, req, "/", http.StatusSeeOther)
}
