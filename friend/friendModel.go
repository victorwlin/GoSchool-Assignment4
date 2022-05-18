package friend

import (
	"GoSchool-Assignment4/data"
	"GoSchool-Assignment4/userp"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// AddFriendToList adds a friend to the current friend list.
func AddFriendToList(res http.ResponseWriter, req *http.Request) {
	if !userp.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	user := userp.GetUser(res, req)

	if len(user.Groups) < 1 {
		data.Error.Printf("User %v tried to add a friend without adding a group first.\n", user.ProfileName)
		http.Error(res, "Cannot add friend until a group is added first.", http.StatusUnauthorized)
		return
	}

	if req.Method == http.MethodPost {
		friendname := req.FormValue("friendname")
		group := req.FormValue("group")
		lastcontact := req.FormValue("lastcontact")
		desiredfreq := req.FormValue("desiredfreq")

		// check if fields have been filled out
		if friendname == "" || group == "" || desiredfreq == "" || lastcontact == "" {

			http.Error(res, "All fields must be filled out.", http.StatusUnauthorized)
			return

		} else {

			// check if friend exists
			if user.Friends.DoesFriendExist(friendname) {
				http.Error(res, "Friend already exists.", http.StatusUnauthorized)
				return
			}

			// conversions
			desiredfreqInt, _ := strconv.Atoi(desiredfreq)
			date, _ := time.Parse("2006-01-02", lastcontact)

			user.Friends.AddFriend(friendname, group, &data.Stack{&data.DateNode{date, nil}, 0}, desiredfreqInt)

			http.Redirect(res, req, "/friends/", http.StatusSeeOther)

		}
	}

	data.Tpl.ExecuteTemplate(res, "addfriend.gohtml", user.Groups)
}

// AddGroup allows the user to add a group to the group slice of the current user.
func AddGroup(res http.ResponseWriter, req *http.Request) {
	if !userp.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	user := userp.GetUser(res, req)

	if req.Method == http.MethodPost {
		group := req.FormValue("group")

		if group == "" {

			http.Error(res, "All fields must be filled out.", http.StatusUnauthorized)
			return

		} else {

			// check if group exists
			exists := false
			for _, v := range user.Groups {
				if v == group {
					exists = true
				}
			}

			if !exists {
				user.Groups = append(user.Groups, group)

				http.Redirect(res, req, "/friends/", http.StatusSeeOther)

			} else {
				http.Error(res, "Group already exists.", http.StatusUnauthorized)
				return
			}
		}

	}

	data.Tpl.ExecuteTemplate(res, "addgroup.gohtml", nil)
}

// EditExistingGroup allows the user to edit an existing group.
func EditExistingGroup(res http.ResponseWriter, req *http.Request) {
	user := userp.GetUser(res, req)

	if req.Method == http.MethodPost {
		group := req.FormValue("group")
		newgroup := req.FormValue("newgroup")

		if group == "" || newgroup == "" {

			http.Error(res, "All fields must be filled out.", http.StatusUnauthorized)
			return

		} else {

			runtime.GOMAXPROCS(2)
			var wg sync.WaitGroup
			wg.Add(2)

			// update all friends that are part of the changed group
			go func() {
				defer wg.Done()

				currentFriend := user.Friends.Head
				for {
					if currentFriend.Group == group {
						currentFriend.Group = newgroup
					}

					if currentFriend.Next == nil {
						break
					} else {
						currentFriend = currentFriend.Next
					}
				}
			}()

			// change group in the group slice
			go func() {
				defer wg.Done()

				for i, v := range user.Groups {
					if v == group {
						user.Groups[i] = newgroup
					}
				}
			}()

			wg.Wait()

			http.Redirect(res, req, "/friends/", http.StatusSeeOther)
		}
	}

	data.Tpl.ExecuteTemplate(res, "editgroup.gohtml", user.Groups)

}

// DeleteGroup allows the user to delete an existing group along with all friends that the group contains.
func DeleteGroup(res http.ResponseWriter, req *http.Request) {
	if !userp.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	user := userp.GetUser(res, req)

	if req.Method == http.MethodPost {
		group := req.FormValue("group")

		if group == "" {
			http.Error(res, "All fields must be filled out.", http.StatusUnauthorized)
			return
		} else {
			runtime.GOMAXPROCS(2)
			var wg sync.WaitGroup
			wg.Add(2)

			// delete all friends in group
			go func() {
				defer wg.Done()

				// identify all occurences
				deleteList := []int{}
				currentFriend := user.Friends.Head
				for index := 0; index < user.Friends.Size; index++ {
					if currentFriend.Group == group {
						deleteList = append(deleteList, index)
					}
					currentFriend = currentFriend.Next
				}

				// delete friends
				for i, v := range deleteList {
					user.Friends.RemoveFriend(v - i)
				}
			}()

			// delete group from group slice
			go func() {
				defer wg.Done()

				// identify group
				var index int
				for i, v := range user.Groups {
					if v == group {
						index = i
						break
					}
				}

				copy(user.Groups[index:], user.Groups[index+1:])
				user.Groups = user.Groups[:len(user.Groups)-1]
			}()

			wg.Wait()

			http.Redirect(res, req, "/friends/", http.StatusSeeOther)
		}
	}

	data.Tpl.ExecuteTemplate(res, "deletegroup.gohtml", user.Groups)
}
