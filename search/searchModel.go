package search

import (
	"GoSchool-Assignment4/data"
	"GoSchool-Assignment4/userp"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// SeqSearch sequentially searches through a friend list for a specific friend and returns the node and index.
func SeqSearch(f *data.FriendList, friendName string) (friend *data.Friend, i int) {

	i = 0
	currentFriend := f.Head
	for {
		if currentFriend.Name == friendName {
			return currentFriend, i
		} else if currentFriend.Next == nil {
			fmt.Println(errors.New("Friend not found."))
			return currentFriend.Next, -1
		} else {
			currentFriend = currentFriend.Next
			i++
		}
	}
}

// DeleteFriend deletes a friend from the list.
func DeleteFriend(res http.ResponseWriter, req *http.Request) {
	if !userp.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	user := userp.GetUser(res, req)

	friend := req.FormValue("friend")
	_, i := SeqSearch(user.Friends, friend)
	user.Friends.RemoveFriend(i)

	http.Redirect(res, req, "/friends/", http.StatusSeeOther)
}

// EditFriendDetails displays a template that allows the user to modify any and all details for that friend.
func EditFriendDetails(res http.ResponseWriter, req *http.Request) {
	if !userp.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	user := userp.GetUser(res, req)

	friend := req.FormValue("friend")
	friendNode, _ := SeqSearch(user.Friends, friend)

	if req.Method == http.MethodPost {
		newFriendName := req.FormValue("newFriendName")
		newGroup := req.FormValue("newGroup")
		newLastContact := req.FormValue("newLastContact")
		newDesiredFreq := req.FormValue("newDesiredFreq")

		if newFriendName != "" {
			if user.Friends.DoesFriendExist(newFriendName) {
				http.Error(res, "User does not exist.", http.StatusUnauthorized)
				return
			}

			friendNode.Name = newFriendName
		}

		if newGroup != "" {
			friendNode.Group = newGroup
		}

		if newLastContact != "" {
			friendNode.LastContact.Top.Date, _ = time.Parse("2006-01-02", newLastContact)
		}

		if newDesiredFreq != "" {
			friendNode.DesiredFreqOfContact, _ = strconv.Atoi(newDesiredFreq)
		}

		http.Redirect(res, req, "/search/?friend="+friendNode.Name, http.StatusSeeOther)
	}

	data.Tpl.ExecuteTemplate(res, "editfrienddetails.gohtml", user.Friends.MakeSearchStruct(friendNode, user))
}
