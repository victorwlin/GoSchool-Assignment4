package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type userProfile struct {
	profileName string
	password    []byte
	groups      []string
	friends     *friendList
}

var users = map[string]*userProfile{
	"victor": {
		profileName: "victor",
		password:    nil,
		groups:      []string{"USA", "SG"},
		friends:     &friendList{nil, 0},
	},
	"tokey": {
		profileName: "tokey",
		password:    nil,
		groups:      []string{"College", "Work"},
		friends:     &friendList{nil, 0},
	},
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	password := os.Getenv("PASSWORD")

	var pw, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	users["victor"].password = pw
	users["tokey"].password = pw

	jimmyStack := stack{nil, 0}
	jimmyStack.push(time.Date(2022, 1, 11, 11, 0, 0, 0, time.UTC))
	jimmyStack.push(time.Date(2022, 1, 31, 11, 0, 0, 0, time.UTC))
	jimmyStack.push(time.Date(2022, 2, 6, 11, 0, 0, 0, time.UTC))
	jimmyStack.push(time.Date(2022, 4, 19, 11, 0, 0, 0, time.UTC))

	benStack := stack{nil, 0}
	benStack.push(time.Date(2022, 1, 31, 11, 0, 0, 0, time.UTC))
	benStack.push(time.Date(2022, 2, 6, 11, 0, 0, 0, time.UTC))

	angelStack := stack{nil, 0}
	angelStack.push(time.Date(2022, 4, 9, 11, 0, 0, 0, time.UTC))
	angelStack.push(time.Date(2022, 4, 13, 11, 0, 0, 0, time.UTC))

	arjunStack := stack{nil, 0}
	arjunStack.push(time.Date(2022, 3, 29, 11, 0, 0, 0, time.UTC))
	arjunStack.push(time.Date(2022, 4, 14, 11, 0, 0, 0, time.UTC))

	ignacioStack := stack{nil, 0}
	ignacioStack.push(time.Date(2022, 4, 15, 11, 0, 0, 0, time.UTC))
	ignacioStack.push(time.Date(2022, 4, 16, 11, 0, 0, 0, time.UTC))

	users["victor"].friends.addFriend("Jimmy", "USA", &jimmyStack, 30)
	users["victor"].friends.addFriend("Ben", "USA", &benStack, 60)
	users["victor"].friends.addFriend("Angel", "SG", &angelStack, 60)
	users["victor"].friends.addFriend("Arjun", "SG", &arjunStack, 14)
	users["victor"].friends.addFriend("Ignacio", "SG", &ignacioStack, 7)

	brianStack := stack{nil, 0}
	brianStack.push(time.Date(2022, 2, 2, 11, 0, 0, 0, time.UTC))
	brianStack.push(time.Date(2022, 4, 17, 11, 0, 0, 0, time.UTC))

	thanhStack := stack{nil, 0}
	thanhStack.push(time.Date(2022, 1, 8, 11, 0, 0, 0, time.UTC))
	thanhStack.push(time.Date(2022, 2, 19, 11, 0, 0, 0, time.UTC))

	lyndaStack := stack{nil, 0}
	lyndaStack.push(time.Date(2022, 4, 1, 11, 0, 0, 0, time.UTC))
	lyndaStack.push(time.Date(2022, 4, 2, 11, 0, 0, 0, time.UTC))

	louieStack := stack{nil, 0}
	louieStack.push(time.Date(2022, 3, 25, 11, 0, 0, 0, time.UTC))
	louieStack.push(time.Date(2022, 3, 30, 11, 0, 0, 0, time.UTC))

	users["tokey"].friends.addFriend("Brian", "College", &brianStack, 180)
	users["tokey"].friends.addFriend("Thanh", "College", &thanhStack, 365)
	users["tokey"].friends.addFriend("Lynda", "Work", &lyndaStack, 90)
	users["tokey"].friends.addFriend("Louie", "Work", &louieStack, 270)
}
