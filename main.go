package main

import (
	"echostarter_example/router"
	"flag"
	"log"
	"os"

	"github.com/labstack/echo/middleware"
)

var (
	flags = flag.NewFlagSet("echostarter_example", flag.ExitOnError)
)

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])
	args := flags.Args()

	switch args[0] {
	default:
		flags.Usage()
		break
	case "run":
		e := router.NewRouter()
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"*"},
			AllowHeaders: []string{"*"},
		}))
		e.Logger.Fatal(e.Start(":8050"))
		os.Exit(0)
		break
	}
}

// this will be printed if app command is wrong
func usage() {
	usagestring := `
add app usage description here
	`

	log.Print(usagestring)
}
