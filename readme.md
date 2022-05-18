Live: https://guarded-wildwood-57389.herokuapp.com/

### Security
**Validation**  
- Username - To avoid confusion, I have decided that usernames should be all lowercase and contain no spaces. If the user tries to create a username on the signup or edit user pages with spaces, they will receive an error. Also all usernames will now automatically be converted to lowercase before further being processed by the app. This functionality has also been added to the login page.

		// Check if username contains spaces. Username should not contain spaces.
		if strings.ContainsAny(username, " ") {
			http.Error(res, "Username cannot contain spaces.", http.StatusForbidden)
		}

		// Username should not be case sensitive, so before working with input, convert to lower case.
		username = strings.ToLower(username)

- Friend Search - Before being allowed to search for a friend, app will check to see if a valid friends list exists. If not, app will log and send error to user.

		if user.Friends.Head != nil {
			// Allow search
		} else {
			data.Error.Println("User tried to search for a friend in an empty friend list.")
			http.Error(res, "No friends to search for.", http.StatusUnauthorized)
			return
		}

- Add Friend - If user tries to add friend, edit group, or delete group before a valid group exists, log and send error to user.

		if  len(user.Groups) <  1 {
			data.Error.Printf("User %v tried to add a friend without adding a group first.\n", user.ProfileName)
			http.Error(res, "Cannot add friend until a group is added first.", http.StatusUnauthorized)
			return
		}

- `DoesFriendExist` - Before checking if friend exists, check if friends list exists.

		if f.Head ==  nil {
			return  false
		} else {
			// check if friend exists in friends list
		}

- Add/Edit Date of Last Contact - The new date of last contact cannot be earlier than the date of last contact that preceded it.

		if date.Before(user.Friends.Head.LastContact.Top.Date) {
			data.Error.Printf("User %v tried to enter a new date of last contact that is earlier than the current date of last contact.\n", user.ProfileName)
			http.Error(res, "New date of last contact must have happened earlier than current date of last contact.", http.StatusUnauthorized)
			return
		}

**Sanitization**  
- html/template - I switched from using text/template to html/template to escape any possible html tags that are input by the user.
- Gorilla Mux - I switched from ServerMux to Gorilla Mux. ServerMux does not change URL request path for CONNECT requests, making the app vulnerable to path traversal attacks. Gorilla Mux limits the allowed request methods.

**Logging**  
- Log Configuration - I set the log configuration in the `init` function of the main.go file. I decided on creating two custom configurations: one for info about what's going on with the app and another to highlight errors that occurred. They are set with the following flags:

		Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Lmicroseconds|log.Llongfile)

- Logging events/errors in package userp

		data.Error.Printf("Unsuccessful login. User entered username %v which does not exist.\n", username)
		data.Error.Printf("Unsuccessful login. Username %v and password do not match.\n", username)
		data.Info.Printf("user %v logged in.\n", username)
		data.Error.Println("Unsuccessful signup. Either username or password fields were left blank.")
		data.Error.Printf("Unsuccessful signup. User tried to enter username with spaces: %v\n", username)
		data.Error.Printf("Unsuccessful signup. Username already taken: %v\n", username)
		data.Info.Printf("user %v signed up.\n", username)
		data.Error.Println("GetUser function was unable to find cookie and redirected user to login.")
		data.Error.Println("AlreadyLoggedIn function was unable to find cookie.")
		data.Info.Printf("Logout function successfully logged out user %v and deleted cookie.\n", data.MapSessions[cookie.Value])

- Logging events in package data

		Info.Println("All data initialized successfully.")

- Logging events and errors in package friend

		data.Error.Printf("User %v submitted a blank search.\n", user.ProfileName)
		data.Error.Printf("User %v searched for a friend who does not exist.\n", user.ProfileName)
		data.Error.Printf("User %v tried to add a friend without filling out all required fields.\n", user.ProfileName)
		data.Error.Printf("User %v tried to add a friend who already exists.\n", user.ProfileName)
		data.Info.Printf("User %v successfully added a friend %v.\n", user.ProfileName, user.Friends.Head.Name)
		data.Error.Printf("User %v tried to add a group without filling out all required fields.\n", user.ProfileName)
		data.Info.Printf("User %v successfully added a group %v.\n", user.ProfileName, group)
		data.Error.Printf("User %v tried to add a group that already exists.\n", user.ProfileName)
		data.Error.Printf("User %v tried to edit a group without filling out all required fields.\n", user.ProfileName)
		data.Info.Printf("User %v successfully edited all friends with a new group %v.\n", user.ProfileName, group)
		data.Info.Printf("User %v successfully edited a group %v in the group slice.\n", user.ProfileName, group)
		data.Error.Printf("User %v tried to delete a group without filling out all required fields.\n", user.ProfileName)
		data.Info.Printf("User %v successfully deleted all friends associated with group %v.\n", user.ProfileName, group)
		data.Info.Printf("User %v successfully deleted a group %v in the group slice.\n", user.ProfileName, group)

- Logging events and errors in package account

		data.Error.Printf("User %v tried to edit user without filling out all required fields.\n", user.ProfileName)
		data.Info.Printf("User %v successfully changed their username.\n", user.ProfileName)
		data.Error.Printf("Unsuccessful edit user by %v. New username already exists.\n", username)
		data.Info.Printf("User %v successfully deleted their user profile and logged out.\n", user.ProfileName)

- Logging events and errors in package search

		data.Error.Printf("User %v submitted a blank search.\n", user.ProfileName)
		data.Info.Printf("User %v successfully added a date of last contact to one of their friends.\n", user.ProfileName)
		data.Error.Println("Unsuccessful search. Friend not found.")
		data.Info.Println("A user successfully deleted a friend.")
		data.Error.Printf("Unsuccessful EditFriendDetails by %v. User chose a friend name that already exists.\n", user.ProfileName)
		data.Info.Printf("User %v successfully submitted a form to edit friend details.\n", user.ProfileName)

**Session Management** 
- Cookies - `HttpOnly` has been set to true to prevent XSS attacks.

**Authentication & Password Management**  
- Environment Variables - A .env file has been created to securely store the password for the two sample user profiles that come with the app. In the `init` function of the data.go file, the .env file is loaded, the password is retrieved, and the hash of the password is stored for the two user profiles. The app logs a fatal error if there is an error loading the .env file.

		err := godotenv.Load(".env")
		if err != nil {
			Error.Println("Error loading .env file")
		}

		password := os.Getenv("PASSWORD")

		var  pw, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

		users["victor"].password = pw
		users["tokey"].password = pw

**HTTP/TLS**  
- HTTP/TLS - Using a self-signed SSL certificate generated by OpenSSL, I switched to a TLS server to encrypt communication between server and client.

		http.ListenAndServeTLS(":5221", "cert.pem", "key.pem", nil)

### Idiomatic Go
- Organization - To facilitate organization, the app has been split into six packages:
	- main - This package contains the server and routing.
	- userp - This package contains the login, signup, and session management.
	- data - This package contains the specialized data structures, their associated methods, log configuration, and variable initialization.
	- friend - This package contains the friend menu and all associated functions.
	- account - This package contains the account management menu and all associated functions.
	- search - This package contains the search menu and all associated functions.
- Documentation - All packages and exported functions have been given comment descriptions with full sentences in accordance with convention.
