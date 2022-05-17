package friend

import "GoSchool-Assignment4/data"

func sortControl(f *data.FriendList, sortType string) {
	for i := 1; i < f.Size; i++ {
		// copy
		copy := f.Traverse(i)

		// loop over sorted and shift to right as necessary
		sorted := i
		switch sortType {

		// friend name
		case "friend":
			for sorted > 0 && copy.Name < f.Traverse(sorted-1).Name {
				f.Traverse(sorted - 1).Next = f.Traverse(sorted + 1)
				copy.Next = f.Traverse(sorted - 1)

				if sorted <= 1 {
					f.Head = copy
				} else {
					f.Traverse(sorted - 2).Next = copy
				}

				sorted--
			}

		// group
		case "group":
			for sorted > 0 && copy.Group < f.Traverse(sorted-1).Group {
				f.Traverse(sorted - 1).Next = f.Traverse(sorted + 1)
				copy.Next = f.Traverse(sorted - 1)

				if sorted <= 1 {
					f.Head = copy
				} else {
					f.Traverse(sorted - 2).Next = copy
				}

				sorted--
			}

		// date of last contact
		case "lastcontact":
			for sorted > 0 && copy.LastContact.Top.Date.After(f.Traverse(sorted-1).LastContact.Top.Date) {
				f.Traverse(sorted - 1).Next = f.Traverse(sorted + 1)
				copy.Next = f.Traverse(sorted - 1)

				if sorted <= 1 {
					f.Head = copy
				} else {
					f.Traverse(sorted - 2).Next = copy
				}

				sorted--
			}

		// desired frequency of contact
		case "desiredfreq":
			for sorted > 0 && copy.DesiredFreqOfContact < f.Traverse(sorted-1).DesiredFreqOfContact {
				f.Traverse(sorted - 1).Next = f.Traverse(sorted + 1)
				copy.Next = f.Traverse(sorted - 1)

				if sorted <= 1 {
					f.Head = copy
				} else {
					f.Traverse(sorted - 2).Next = copy
				}

				sorted--
			}

		// recommended date of next contact
		case "recdate":
			for sorted > 0 && copy.LastContact.Top.Date.AddDate(0, 0, copy.DesiredFreqOfContact).Before(f.Traverse(sorted-1).LastContact.Top.Date.AddDate(0, 0, f.Traverse(sorted-1).DesiredFreqOfContact)) {
				f.Traverse(sorted - 1).Next = f.Traverse(sorted + 1)
				copy.Next = f.Traverse(sorted - 1)

				if sorted <= 1 {
					f.Head = copy
				} else {
					f.Traverse(sorted - 2).Next = copy
				}

				sorted--
			}

		default:
			return
		}
	}

}
