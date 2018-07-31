package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/galdor/go-cmdline"
)

func main() {
	opts := cmdline.New()

	opts.AddOption("e", "editor", "application", "editor to open changed files [default: vscode]")
	opts.AddOption("m", "merge-id", "id", "ID for the merge request")
	opts.AddOption("p", "project-id", "id", "ID for the GitLab project containing merge request")
	opts.AddOption("t", "token", "token", "your GitLab personal access token")
	opts.AddFlag("v", "version", "print differ version")

	opts.Parse(os.Args)

	if opts.IsOptionSet("v") {
		fmt.Println(version)
		os.Exit(0)
	}

	editor := "vscode"
	if opts.IsOptionSet("e") {
		editor = opts.OptionValue("e")
	}

	err := validate(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	token := opts.OptionValue("t")
	projectID := opts.OptionValue("p")
	mergeID := opts.OptionValue("m")
	config := NewConfig(token, projectID, mergeID)
	mergeData, err := NewMergeData(&config)
	if err != nil {
		log.Fatal(err)
	}

	err = mergeData.changeBranch()
	if err != nil {
		log.Fatal(err)
	}

	err = mergeData.openFiles(editor)
	if err != nil {
		log.Fatal(err)
	}
}

func validate(opts *cmdline.CmdLine) error {
	if !opts.IsOptionSet("t") {
		return errors.New("token must be supplied")
	}

	if !opts.IsOptionSet("p") {
		return errors.New("project-id must be supplied")
	}

	if !opts.IsOptionSet("m") {
		return errors.New("merge-id must be supplied")
	}

	return nil
}
