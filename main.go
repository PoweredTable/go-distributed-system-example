package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
	"github.com/spf13/cobra"
	"go-client-server/server"
	"go-client-server/client"
	"go-client-server/common" // Make sure to import common if needed
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "App with client and server commands",
	}

	// Server command
	var serverCmd = &cobra.Command{
		Use:   "server [address] [port]",
		Short: "Start the server",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			address := args[0]
			port, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid port number")
				os.Exit(1)
			}
			s := server.NewTCPServer(fmt.Sprintf("%s:%d", address, port))
			log.Fatal(s.Start())
		},
	}

	// Client command
	var clientCmd = &cobra.Command{
		Use:   "client [file] [routines] [servers...]",
		Short: "Run the client with file, routines, and servers",
		Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			file := args[0]
			routines, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid number of routines")
				os.Exit(1)
			}
			servers := args[2:]

			if !common.FileExists(file) {
				fmt.Println("File does not exist:", file)
				os.Exit(1)
			}

			// Create a new client and process the text
			client := client.NewClient(routines, servers)
			text, err := os.ReadFile(file)
			if err != nil {
				fmt.Println("Error reading file:", err)
				os.Exit(1)
			}
			client.ProcessText(string(text))
		},
	}
	
	rootCmd.AddCommand(serverCmd, clientCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
