package domain

type Package struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Cmd         string `json:"cmd"`
	Repo        string `json:"repo"`
	InstallCmd  string `json:"install_cmd"`
}
