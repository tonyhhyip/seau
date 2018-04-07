package api

type Vendor struct {
	Providers map[string]Hash `json:"providers"`
}
