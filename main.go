package main

import (
	"log"

	"git.front.kjuulh.io/kjuulh/dagger-go/pkg/cli"
)

func main() {
	if err := cli.NewCli().Execute(); err != nil {
		log.Fatal(err)
	}
}
