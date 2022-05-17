// Package account controls all aspects of the account management menu.
package account

import (
	"GoSchool-Assignment4/data"
	"GoSchool-Assignment4/userp"
	"net/http"
)

// AccountManagement displays the account management template.
func AccountManagement(res http.ResponseWriter, req *http.Request) {
	if !userp.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	user := userp.GetUser(res, req)

	data.Tpl.ExecuteTemplate(res, "accountmanagement.gohtml", user.ProfileName)
}
