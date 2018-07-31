package main

type Config struct {
	MergeID   string
	ProjectID string
	Token     string
}

func NewConfig(token, projectID, mergeID string) Config {
	return Config{
		MergeID:   mergeID,
		ProjectID: projectID,
		Token:     token,
	}
}
