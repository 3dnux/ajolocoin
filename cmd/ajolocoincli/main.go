package main

import (
	"os"

	"github.com/ajolocoin/app"
	"github.com/ajolocoin/cmd/myajolocoind/cmd"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/version"
)

func main() {
	// Configura la aplicación
	app.Setup()

	// Añade los comandos específicos de tu aplicación
	rootCmd, _ := cmd.NewRootCmd()

	// Añade el comando de inicio de servidor
	executor := server.NewCommand(rootCmd)

	// Añade la versión del software
	executor.Version = version.Version
	executor.Commit = version.Commit

	// Inicia el servidor
	if err := executor.Execute(); err != nil {
		os.Exit(1)
	}
}
