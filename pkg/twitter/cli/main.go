package main

import (
	twitter "github.com/FriendlyUser/finfiber/pkg/twitter/util"
	"log"
)


func main() {
	body, _ := twitter.MakeRequest()
	log.Println(string([]byte(body)))
}
