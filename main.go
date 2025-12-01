/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/dream-horizon-org/odin/cmd"
	_ "github.com/dream-horizon-org/odin/cmd/configure"
	_ "github.com/dream-horizon-org/odin/cmd/create"
	_ "github.com/dream-horizon-org/odin/cmd/delete"
	_ "github.com/dream-horizon-org/odin/cmd/deploy"
	_ "github.com/dream-horizon-org/odin/cmd/describe"
	_ "github.com/dream-horizon-org/odin/cmd/list"
	_ "github.com/dream-horizon-org/odin/cmd/operate"
	_ "github.com/dream-horizon-org/odin/cmd/set"
	_ "github.com/dream-horizon-org/odin/cmd/status"
	_ "github.com/dream-horizon-org/odin/cmd/undeploy"
	_ "github.com/dream-horizon-org/odin/internal/ui"
)

func main() {
	cmd.Execute()
}
