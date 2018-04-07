package api

type ProviderList struct {
	Packages         []string        `json:"packages"`
	Includes         map[string]Hash `json:"includes"`
	ProvidersUrl     string          `json:"providers-url"`
	ProviderIncludes map[string]Hash `json:"provider-includes"`
	NotifyBatch      string          `json:"notify-batch"`
}

type Hash struct {
	Sha256 string `json:"sha256"`
}
