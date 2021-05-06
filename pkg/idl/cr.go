package idl

type CRQueryReq struct {
	ID int64 `json:"id"`
}

type CRCreateReq struct {
	ProjectID  int64  `json:"project_id"`
	Branch     string `json:"branch"`
	BaseCommit string `json:"base"`
	HeadCommit string `json:"head"`
	Reviewer   string `json:"reviewer"`
	Cc         string `json:"cc"`
}

type Project struct {
	RepoID        int64  `json:"repo_id"`
	Desc          string `json:"desc"`
	Owner         string `json:"owner"`
	ProtectBranch bool   `json:"protect_branch"`
	Status        int64  `json:"status"`
}

type FlowTemplate struct {
	BuildCheck bool `json:"check_build"`
	CodeLint   bool `json:"code_lint"`
	CR         bool `json:"cr"`
	CloudTest  bool `json:"cloud_test"`
}

type Demand struct {
	BranchName   string       `json:"branch"`
	Creator      string       `json:"creator"`
	CreateAt     string       `json:"create_at"`
	UpdateAt     string       `json:"update_at"`
	Desc         string       `json:"desc"`
	Prd          string       `json:"prd"`
	DiffDetail   string       `json:"diff_detail"`
	TemplateConf FlowTemplate `json:"tpl"`
}

type CRRecord struct {
	DemandID   int64  `json:"demand_id"`
	DetailLink string `json:"detail_link"`
	Status     string `json:"status"`
}

type BuildInfo struct {
	ID          int64  `json:"id"`
	ProjectID   int64  `json:"project_id"`
	BuildShell  string `json:"build_shell"`
	BuildLog    string `json:"build_log"`
	BuildResult string `json:"build_result"`
	StartAt     string `json:"start_at"`
	EndAt       string `json:"end_at"`
}

type Product struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Owners   string `json:"owners"`
	StartAt  string `json:"start_at"`
	UpdateAt string `json:"update_at"`
}
