package data

// Friend contains all the information for a friend.
type Friend struct {
	Name                 string
	Group                string
	LastContact          *Stack
	DesiredFreqOfContact int
	Next                 *Friend
}

// FriendList contains the entire linked list with all friends.
type FriendList struct {
	Head *Friend
	Size int
}

// create struct for display
type friendRender struct {
	Name                 string
	Group                string
	LastContact          string
	DesiredFreqOfContact int
	RecDateOfNextContact string
}

// create struct for search display
type searchRender struct {
	Name                 string
	Group                string
	HistoryOfContact     []string
	DesiredFreqOfContact int
	RecDateOfNextContact string
	LastContact          string
	AvailableGroups      []string
}

// AddFriend allows the user to add a friend to their list.
func (f *FriendList) AddFriend(name string, group string, lastContact *Stack, desiredFreqOfContact int) {
	newFriend := &Friend{
		Name:                 name,
		Group:                group,
		LastContact:          lastContact,
		DesiredFreqOfContact: desiredFreqOfContact,
		Next:                 nil,
	}

	if f.Head == nil {
		f.Head = newFriend
	} else {
		currentFriend := f.Head

		// traversal
		for currentFriend.Next != nil {
			currentFriend = currentFriend.Next
		}

		currentFriend.Next = newFriend
	}

	f.Size++
}

// RemoveFriend allows the user to remove a friend from their list.
func (f *FriendList) RemoveFriend(i int) {
	if i == 0 {
		f.Head = f.Head.Next
	} else {
		currentFriend := f.Head
		prevFriend := f.Head
		for index := 0; index < i; index++ {
			prevFriend = currentFriend
			currentFriend = currentFriend.Next
		}
		prevFriend.Next = currentFriend.Next
	}
	f.Size--
}

// Traverse takes an index integer and returns the friend at that point.
func (f *FriendList) Traverse(index int) (friend *Friend) {
	currentFriend := f.Head
	for i := 0; i < index; i++ {
		currentFriend = currentFriend.Next
	}
	return currentFriend
}

// MakeFriendsSlice makes a friends slice for the friends template to render.
func (f *FriendList) MakeFriendsSlice() (friendsSlice []friendRender) {

	currentFriend := f.Head
	for {
		insertFriend := friendRender{}
		if currentFriend == nil {

			return []friendRender{}

		} else {
			insertFriend = friendRender{
				Name:                 currentFriend.Name,
				Group:                currentFriend.Group,
				LastContact:          currentFriend.LastContact.Top.Date.Format("02 Jan 2006"),
				DesiredFreqOfContact: currentFriend.DesiredFreqOfContact,
				RecDateOfNextContact: currentFriend.LastContact.Top.Date.AddDate(0, 0, currentFriend.DesiredFreqOfContact).Format("02 Jan 2006"),
			}
		}

		friendsSlice = append(friendsSlice, insertFriend)

		if currentFriend.Next == nil {
			break
		} else {
			currentFriend = currentFriend.Next
		}
	}

	return friendsSlice
}

// DoesFriendExist checks the list to see if the friend exists.
func (f *FriendList) DoesFriendExist(friendname string) bool {
	currentFriend := f.Head
	for {
		if currentFriend.Name == friendname {
			return true
		}

		if currentFriend.Next == nil {
			break
		} else {
			currentFriend = currentFriend.Next
		}
	}

	return false
}

// MakeSearchStruct returns a struct that allows the search template to display friend information for one friend.
func (f *FriendList) MakeSearchStruct(friend *Friend, user *UserProfile) (searchStruct searchRender) {
	// create slice of dates contacted
	dateHistory := []string{}
	currentDate := friend.LastContact.Top
	for currentDate != nil {
		dateHistory = append(dateHistory, currentDate.Date.Format("02 Jan 2006"))
		currentDate = currentDate.Last
	}

	lastContact := friend.LastContact.Top.Date.Format("02 Jan 2006")

	searchStruct = searchRender{
		Name:                 friend.Name,
		Group:                friend.Group,
		HistoryOfContact:     dateHistory,
		DesiredFreqOfContact: friend.DesiredFreqOfContact,
		RecDateOfNextContact: friend.LastContact.Top.Date.AddDate(0, 0, friend.DesiredFreqOfContact).Format("02 Jan 2006"),
		LastContact:          lastContact,
		AvailableGroups:      user.Groups,
	}

	return searchStruct
}
