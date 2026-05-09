package domain

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Repo        string `json:"repo"`
	InstallCmd  string `json:"install_cmd"`
	UsageHint   string `json:"usage_hint"`
}
