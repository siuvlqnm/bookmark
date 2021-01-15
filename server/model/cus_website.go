package model

type CusWebsite struct {
	Protocol    string `json:"protocol"`
	Domain      string `json:"domain"`
	Port        int    `json:"port"`
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
