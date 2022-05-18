// Package data handles all aspects of data, including specialized data structures, template initialization, and log configuration.
package data

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// UserProfile is the struct that keeps all information for a user.
type UserProfile struct {
	ProfileName string
	Password    []byte
	Groups      []string
	Friends     *FriendList
}

// Users map contains all users currently in the app.
var Users = map[string]*UserProfile{
	"victor": {
		ProfileName: "victor",
		Password:    nil,
		Groups:      []string{"USA", "SG"},
		Friends:     &FriendList{nil, 0},
	},
	"tokey": {
		ProfileName: "tokey",
		Password:    nil,
		Groups:      []string{"College", "Work"},
		Friends:     &FriendList{nil, 0},
	},
}

var (
	Tpl         *template.Template
	MapSessions = map[string]string{}

	Info  *log.Logger
	Error *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Lmicroseconds|log.Llongfile)

	Tpl = template.Must(template.ParseGlob("templates/*"))

	err := godotenv.Load(".env")
	if err != nil {
		Error.Println("Error loading .env file")
	}

	password := os.Getenv("PASSWORD")

	var pw, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	Users["victor"].Password = pw
	Users["tokey"].Password = pw

	jimmyStack := Stack{nil, 0}
	jimmyStack.Push(time.Date(2022, 1, 11, 11, 0, 0, 0, time.UTC))
	jimmyStack.Push(time.Date(2022, 1, 31, 11, 0, 0, 0, time.UTC))
	jimmyStack.Push(time.Date(2022, 2, 6, 11, 0, 0, 0, time.UTC))
	jimmyStack.Push(time.Date(2022, 4, 19, 11, 0, 0, 0, time.UTC))

	benStack := Stack{nil, 0}
	benStack.Push(time.Date(2022, 1, 31, 11, 0, 0, 0, time.UTC))
	benStack.Push(time.Date(2022, 2, 6, 11, 0, 0, 0, time.UTC))

	angelStack := Stack{nil, 0}
	angelStack.Push(time.Date(2022, 4, 9, 11, 0, 0, 0, time.UTC))
	angelStack.Push(time.Date(2022, 4, 13, 11, 0, 0, 0, time.UTC))

	arjunStack := Stack{nil, 0}
	arjunStack.Push(time.Date(2022, 3, 29, 11, 0, 0, 0, time.UTC))
	arjunStack.Push(time.Date(2022, 4, 14, 11, 0, 0, 0, time.UTC))

	ignacioStack := Stack{nil, 0}
	ignacioStack.Push(time.Date(2022, 4, 15, 11, 0, 0, 0, time.UTC))
	ignacioStack.Push(time.Date(2022, 4, 16, 11, 0, 0, 0, time.UTC))

	Users["victor"].Friends.AddFriend("Jimmy", "USA", &jimmyStack, 30)
	Users["victor"].Friends.AddFriend("Ben", "USA", &benStack, 60)
	Users["victor"].Friends.AddFriend("Angel", "SG", &angelStack, 60)
	Users["victor"].Friends.AddFriend("Arjun", "SG", &arjunStack, 14)
	Users["victor"].Friends.AddFriend("Ignacio", "SG", &ignacioStack, 7)

	brianStack := Stack{nil, 0}
	brianStack.Push(time.Date(2022, 2, 2, 11, 0, 0, 0, time.UTC))
	brianStack.Push(time.Date(2022, 4, 17, 11, 0, 0, 0, time.UTC))

	thanhStack := Stack{nil, 0}
	thanhStack.Push(time.Date(2022, 1, 8, 11, 0, 0, 0, time.UTC))
	thanhStack.Push(time.Date(2022, 2, 19, 11, 0, 0, 0, time.UTC))

	lyndaStack := Stack{nil, 0}
	lyndaStack.Push(time.Date(2022, 4, 1, 11, 0, 0, 0, time.UTC))
	lyndaStack.Push(time.Date(2022, 4, 2, 11, 0, 0, 0, time.UTC))

	louieStack := Stack{nil, 0}
	louieStack.Push(time.Date(2022, 3, 25, 11, 0, 0, 0, time.UTC))
	louieStack.Push(time.Date(2022, 3, 30, 11, 0, 0, 0, time.UTC))

	Users["tokey"].Friends.AddFriend("Brian", "College", &brianStack, 180)
	Users["tokey"].Friends.AddFriend("Thanh", "College", &thanhStack, 365)
	Users["tokey"].Friends.AddFriend("Lynda", "Work", &lyndaStack, 90)
	Users["tokey"].Friends.AddFriend("Louie", "Work", &louieStack, 270)

	Info.Println("All data initialized successfully.")
}
