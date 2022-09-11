package main

import (
	"fmt"
	"log"

	"github.com/rajiv2205/http_server_golang/internal/database"
)

func main() {
	c := database.NewClient("db.json")
	err := c.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database ensured!")

	// usr, err := c.CreateUser("rajivsharma.2205@gmail.com", "secret#123", "rajiv sharma", 31)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("user created!", usr)

	// updatedUser, err := c.UpdateUser("rajivsharma.2205@gmail.com", "password", "rajiv", 19)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("user updated", updatedUser)

	gotUser, err := c.GetUser("rajivsharma.2205@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user got", gotUser)

	// err = c.DeleteUser("rajivsharma.2205@gmail.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("user deleted")

	// _, err = c.GetUser("rajivsharma.2205@gmail.com")
	// if err == nil {
	// 	log.Fatal("shouldn't be able to get user that was deleted")
	// }
	// fmt.Println("user confirmed deleted")

	post, err := c.CreatePost("rajivsharma.2205@gmail.com", "my cat is way too fat")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("post created", post)

	secondPost, err := c.CreatePost("rajivsharma.2205@gmail.com", "my cat is getting skinny now")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("another post created", secondPost)

	posts, err := c.GetPosts("rajivsharma.2205@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("got posts", posts)

}
