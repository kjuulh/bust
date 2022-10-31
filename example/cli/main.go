package main

import (
	"log"

	"git.front.kjuulh.io/kjuulh/bust/pkg/cli"
)

func main() {
	if err := cli.NewCli().Execute(); err != nil {
		log.Fatal(err)
	}
}
