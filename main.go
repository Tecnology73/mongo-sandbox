package main

import (
	"fmt"
	"mongo-sandbox/Models"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// Lookup Admin
	err, user := Models.FindOne[Models.User](bson.M{"email": "a@b.co"})

	if err != nil {
		fmt.Println("[Error] Failed to find the user!\n" + err.Error())
		return
	}

	// Change Admin's name
	user.Name = "George"

	// Persist changes
	if err, _ := Models.Update(user); err != nil {
		fmt.Println("[Error] Failed to save user!\n" + err.Error())
		return
	}

	// Create pleb User
	err, newUser := Models.Create(&Models.User{
		Name:  "Framk",
		Email: "frank@gmail.com",
	})

	// Something went wrong creating the pleb
	if err != nil {
		fmt.Println("[Error] Failed to create new user!\n", err.Error())
		return
	}

	// Output the new User
	fmt.Printf("User: %+v\n", newUser)
}
