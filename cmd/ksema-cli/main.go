package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"golang.org/x/term"

	"github.com/carakacloud/ksema"
	"github.com/daniel/ksema-cli/internal/commands"
	"github.com/daniel/ksema-cli/internal/config"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: ksema-cli <command> [arguments]")
		fmt.Println("Use 'help' to see available commands.")
		os.Exit(1)
	}

	command := args[1]
	commandArgs := args[2:]

	if command == "help" {
		handler := commands.NewCommandHandler(nil)
		if err := handler.Execute(command, commandArgs); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		return
	}

	// Load config and create Ksema connection for commands that need it
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Error: Missing required environment variables.")
		fmt.Println("Please set the following environment variables:")
		fmt.Println("  - KSEMA_HOST: The Ksema server host")
		fmt.Println("  - KSEMA_API_KEY: Your Ksema API key")
		fmt.Println("\nYou can set these in a .env file or in your environment.")
		os.Exit(1)
	}

	fmt.Print("Enter KSEMA PIN: ")
	bytePIN, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal("Error reading PIN:", err)
	}
	fmt.Println() // Add a newline after PIN input
	cfg.KsemaPIN = string(bytePIN)

	user, err := ksema.New(cfg.KsemaServerIP, cfg.KsemaAPIKey, cfg.KsemaPIN)
	if err != nil {
		log.Fatal("Error creating Ksema object:", err)
	}

	if err := user.Ping(); err != nil {
		log.Fatal("Failed to ping server:", err)
	}

	handler := commands.NewCommandHandler(user)
	if err := handler.Execute(command, commandArgs); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
