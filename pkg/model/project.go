package models

type Project struct {
	ID               int64  `json:"id"`
	GitID            int64  `json:"git_project_id"`
	GitRepoURI       string `json:"git_repo_uri"`
	Level            int64  `json:"level"`
	Desc             string `json:"desc"`
	ProtectMasterSID int64  `json:"protect_sid"`
	Admin            string `json:"admins"`
	Reviewers        string `json:"reviewers"`
	MasterCommitID   string `json:"master_lastest_commit_id"`
}

type DVersion struct {
	Branch    string `json:"branch"`
	ProjectID int64  `json:"git_project_id"`
	Status    int64  `json:"status"`
	CommitID  string `json:"commit_id"`
}
