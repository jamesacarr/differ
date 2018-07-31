package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

type FileData struct {
	DeletedFile bool   `json:"deleted_file"`
	NewPath     string `json:"new_path"`
}

type MergeData struct {
	Changes      []FileData `json:"changes"`
	SourceBranch string     `json:"source_branch"`
}

func NewMergeData(config *Config) (*MergeData, error) {
	fmt.Print("===> Getting merge request details\n")
	url := fmt.Sprintf("https://gitlab.geniesolutions.com.au/api/v4/projects/%s/merge_requests/%s/changes?private_token=%s", config.ProjectID, config.MergeID, config.Token)

	client := http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		return nil, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	output := MergeData{}
	jsonErr := json.Unmarshal(body, &output)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &output, nil
}

func (m *MergeData) changeBranch() error {
	fmt.Printf("===> Changing to branch: %s\n", m.SourceBranch)
	cmd := exec.Command("git", "checkout", m.SourceBranch)
	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Print("===> Pulling latest changes\n")
	cmd = exec.Command("git", "pull", "-r")
	return cmd.Run()
}

func (m *MergeData) openFiles(editor string) error {
	editorCmd := editor
	if len(editorMap[editor]) > 0 {
		editorCmd = editorMap[editor]
	}

	files := make([]string, 0)
	for _, fileData := range m.Changes {
		if !fileData.DeletedFile {
			files = append(files, fmt.Sprintf("./%s", fileData.NewPath))
		}
	}

	fmt.Printf("===> Opening files using: %s\n", editor)
	for _, fileName := range files {
		fmt.Printf("     - %s\n", fileName)
	}
	cmd := exec.Command(editorCmd, prepend(files, ".")...)
	return cmd.Run()
}

func prepend(original []string, element string) []string {
	return append([]string{element}, original...)
}
