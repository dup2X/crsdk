package models

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"

	"github.com/dup2X/crsdk/idl"
	"github.com/dup2X/gopkg/logger"
)

var phToken string = ""

var msgTemplate = `%s-%s

Summary:%s

Test Plan:%s

Reviewers:%s

Subscribers:%s`

func CreateRevision(ctx context.Context, req *idl.CRCreateReq) (int64, error) {
	// TODO cwd
	xxx := fmt.Sprintf("%d_%s_%d.txt", req.ProjectID, req.Branch, time.Now().Unix())
	msgFile := fmt.Sprintf("./.git/arc/%s", xxx)
	now := time.Now()
	data := []byte(fmt.Sprintf(msgTemplate, req.Branch, now.Format("2006-01-02 15:04:05"), "auto", "test plan A", req.Reviewer, req.Cc))
	err := ioutil.WriteFile(msgFile, data, 0666)
	if err != nil {
		return 0, err
	}
	cmd := exec.Command("arc", "diff", req.BaseCommit, "--create", "--allow-untracked", "--message-file", msgFile)
	std := bytes.Buffer{}
	cmd.Stdout = &std
	stderr := bytes.Buffer{}
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return 0, err
	}
	logger.Infof(ctx, logger.DLTagUndefined, "stdout=%s,stderr=%s", std.String(), stderr.String())
	return 1, nil
}
