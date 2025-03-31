package main

import (
	"fmt"
	"github.com/dabates/gator/internal/config"
)

func main() {
	c := config.Read()
	c.SetUser("Dave")
	c = config.Read()

	fmt.Println(c.DbUrl)
	fmt.Println(c.CurrentUserName)
}
