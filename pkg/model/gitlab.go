package models

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dup2X/gopkg/logger"
	"github.com/xanzy/go-gitlab"
)

var (
	gitRootURI  string
	gitAPIToken string
)

func GetUser(ctx context.Context) {
	c, err := gitlab.NewClient(gitAPIToken, gitlab.WithBaseURL("http://git.dangdang.com/api/v4"))
	if err != nil {
		logger.Warnf(ctx, logger.DLTagUndefined, "")
		return
	}
	name := "yangdongchen"
	users, _, err := c.Users.ListUsers(&gitlab.ListUsersOptions{Search: &name})
	if err != nil {
		logger.Warnf(ctx, logger.DLTagUndefined, "")
		return
	}
	logger.Infof(ctx, logger.DLTagUndefined, "users:%v", users)
}

func GetBranchInfo(ctx context.Context, projID int64, branchName string) (*GitBranch, error) {
	uri := fmt.Sprintf("%s/projects/%d/repository/branches/%s", gitRootURI, projID, branchName)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("PRIVATE-TOKEN", gitAPIToken)
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var br = &GitBranch{}
	err = json.Unmarshal(data, br)
	if err != nil {
		return nil, err
	}
	return br, nil
}

func CreateBranch(ctx context.Context, projID int64, branchName, parentBranch string) (*GitBranch, error) {
	uri := fmt.Sprintf("%s/projects/%d/repository/branches?branch=%s&ref=%s", gitRootURI, projID, branchName, parentBranch)
	req, err := http.NewRequest("POST", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("PRIVATE-TOKEN", gitAPIToken)
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var br = &GitBranch{}
	err = json.Unmarshal(data, br)
	if err != nil {
		return nil, err
	}
	return br, nil
}

func GetUserInfo() {

}

func GetGitRepo() {

}

type GitBranch struct {
	Name   string `json:"name"`
	Commit struct {
		ID string `json:"id"`
	} `json:"commit"`
}
