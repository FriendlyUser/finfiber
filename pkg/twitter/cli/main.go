package main

import (
	twitter "github.com/FriendlyUser/finfiber/twitter/util"
	"log"
)


func main() {
	body, _ := twitter.MakeRequest()
	log.Println(string([]byte(body)))
}
