package repo

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type RepositoryConfig struct {
	RepoName   string `json:"repo_name"`
	ConnPrefix string `json:"conn_prefix"`
	ConnHost   string `json:"conn_host"`
	ConnPort   string `json:"conn_port"`
	// Collections map[string]string `json:"collections"`
}

func InitRepoConfig(cfgPath string) (*RepositoryConfig, error) {
	var config RepositoryConfig

	file, err := os.Open("./configs/repo_config_debug.json")
	if err != nil {
		return nil, err
	}

	buff, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(buff, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
