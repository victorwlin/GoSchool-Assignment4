package main

import (
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func login(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/friends/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")

		// Username should not be case sensitive, so before working with input, convert to lower case.
		username = strings.ToLower(username)

		// check if user exists using entered username
		user, ok := users[username]
		if !ok {
			http.Error(res, "User does not exist.", http.StatusUnauthorized)
			return
		}

		// check if password matches our records
		err := bcrypt.CompareHashAndPassword(user.password, []byte(password))
		if err != nil {
			http.Error(res, "Username and password do not match.", http.StatusForbidden)
			return
		}

		// create session
		id := uuid.NewV4()
		cookie := &http.Cookie{
			Name:     "FriendTrackerCookie",
			Value:    id.String(),
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)

		mapSessions[cookie.Value] = username

		http.Redirect(res, req, "/friends/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/friends/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")

		if username == "" || password == "" {
			http.Error(res, "Both fields must contain values.", http.StatusForbidden)
			return
		} else {

			// Check if username contains spaces. Username should not contain spaces.
			if strings.ContainsAny(username, " ") {
				http.Error(res, "Username cannot contain spaces.", http.StatusForbidden)
			}

			// Username should not be case sensitive, so before working with input, convert to lower case.
			username = strings.ToLower(username)

			// check if username already exists
			if _, ok := users[username]; ok {
				http.Error(res, "Username already taken.", http.StatusForbidden)
				return
			}

			// create session
			id := uuid.NewV4()
			cookie := &http.Cookie{
				Name:     "FriendTrackerCookie",
				Value:    id.String(),
				Path:     "/",
				HttpOnly: true,
			}
			http.SetCookie(res, cookie)

			mapSessions[cookie.Value] = username

			// create password
			pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
			if err != nil {
				http.Error(res, "Internal server error", http.StatusInternalServerError)
				return
			}

			// create user profile
			users[username] = &userProfile{
				profileName: username,
				password:    pw,
				groups:      []string{},
				friends:     &friendList{nil, 0},
			}

			http.Redirect(res, req, "/friends/", http.StatusSeeOther)
			return
		}
	}

	tpl.ExecuteTemplate(res, "signup.gohtml", nil)
}
