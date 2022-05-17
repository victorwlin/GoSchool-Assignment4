package main

import "net/http"

func getUser(res http.ResponseWriter, req *http.Request) (user *userProfile) {
	// get current session cookie
	cookie, err := req.Cookie("FriendTrackerCookie")

	// if cookie doesn't exist, redirect to login
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if _, ok := mapSessions[cookie.Value]; ok {
		username := mapSessions[cookie.Value]

		user = users[username]
	}

	return user
}

func alreadyLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("FriendTrackerCookie")
	if err != nil {
		return false
	}

	username := mapSessions[cookie.Value]
	_, ok := users[username]

	return ok
}

func logout(res http.ResponseWriter, req *http.Request) {
	cookie, _ := req.Cookie("FriendTrackerCookie")

	// delete session
	delete(mapSessions, cookie.Value)

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
