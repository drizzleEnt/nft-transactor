package main

import (
	"github.com/drizzleent/nft-transactor/app"
)

func main() {
	a := app.NewApp()
	defer a.Close()
	a.Run()
}
