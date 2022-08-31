package main

import (
	"log"
	"strings"

	"github.com/alecthomas/kong"
)

type GenCli struct {
	Out          string `help:"The file to write the generated keyset to"`
	MasterKeyUri string `help:"The master key URI, if any. GCP KMS and AWS KMS are supported."`
}

type ConvertCli struct {
	In              string `help:"The file to read the original keyset from"`
	MasterKeyUri    string `help:"The master key URI, if any. GCP KMS and AWS KMS are supported."`
	Out             string `help:"The file to write the converted keyset to"`
	NewMasterKeyUri string `help:"The new master key URI, if any. GCP KMS and AWS KMS are supported."`
}

type Cli struct {
	GenerateKeyset GenCli     `cmd:"" help:"Generates a KeySet"`
	ConvertKeyset  ConvertCli `cmd:"" help:"Converts a KeySet"`
}

var CLI Cli

func main() {
	kongCtx := kong.Parse(&CLI,
		kong.Description("TinKey key manager for Tink"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Tree: true,
		}))

	// parsing the command into command + subcommand
	commands := strings.Fields(kongCtx.Command())
	command := commands[0]

	switch command {
	case "generate-keyset":
		HandleGenerateKeySet(CLI.GenerateKeyset)
	case "convert-keyset":
		HandleConvertKeySet(CLI.ConvertKeyset)
	default:
		panic(kongCtx.Command())
	}
}

func HandleGenerateKeySet(cli GenCli) {
	if cli.Out == "" {
		log.Fatal("Missing output parameter")
	}
	handle := GenerateKeyset()
	WriteKeySet(cli.Out, cli.MasterKeyUri, handle)
}

func HandleConvertKeySet(cli ConvertCli) {
	handle := ReadKeySet(cli.In, cli.MasterKeyUri)
	WriteKeySet(cli.Out, cli.NewMasterKeyUri, handle)
}
