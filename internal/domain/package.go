package domain

type Package struct {
	Name       string `json:"name"`
	Repo       string `json:"repo"`
	InstallCmd string `json:"install_cmd"`
}
