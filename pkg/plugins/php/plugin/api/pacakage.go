package api

type Stability string

const (
	Dev    Stability = "dev"
	Alpha  Stability = "alpha"
	Beta   Stability = "beta"
	RC     Stability = "RC"
	Stable Stability = "stable"
)

type Packages struct {
	Packages map[string]Package
}

type Package map[string]PackageVersion

type PackageVersion struct {
	Name              string        `json:"name"`
	Description       string        `json:"description,omitempty"`
	Keywords          []string      `json:"keywords,omitempty"`
	Homepage          string        `json:"homepage,omitempty"`
	License           string        `json:"license,omitempty"`
	Authors           []Author      `json:"authors,omitempty"`
	Support           Support       `json:"support,omitempty"`
	Time              string        `json:"time,omitempty"`
	Type              string        `json:"type,omitempty"`
	Version           string        `json:"version"`
	VersionNormalized string        `json:"version_normalized"`
	Source            PackageSource `json:"source"`
	Dist              PackageSource `json:"dist"`
	Require           PackageList   `json:"require,omitempty"`
	RequireDev        PackageList   `json:"require-dev,omitempty"`
	Conflict          PackageList   `json:"conflict,omitempty"`
	Replace           PackageList   `json:"replace,omitempty"`
	Provider          PackageList   `json:"provider,omitempty"`
	Suggest           PackageList   `json:"suggest,omitempty"`
	AutoLoad          AutoLoad      `json:"autoload"`
	AutoloadDev       AutoLoad      `json:"autoload-dev,omitempty"`
	MinimumStability  Stability     `json:"minimum-stability,omitempty"`
	PreferStable      bool          `json:"prefer-stable,omitempty"`
}

type Author struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Homepage string `json:"homepage,omitempty"`
	Role     string `json:"role,omitempty"`
}

type Support struct {
	Email  string `json:"email,omitempty"`
	Issue  string `json:"issue,omitempty"`
	Forum  string `json:"forum,omitempty"`
	Wiki   string `json:"wiki,omitempty"`
	IRC    string `json:"irc,omitempty"`
	Source string `json:"source,omitempty"`
	Docs   string `json:"docs,omitempty"`
	RSS    string `json:"rss,omitempty"`
}

type PackageList map[string]string

type AutoLoad struct {
	Psr0                map[string]string `json:"psr-0,omitempty"`
	Psr4                map[string]string `json:"psr-4,omitempty"`
	ClassMap            []string          `json:"classmap,omitempty"`
	Files               []string          `json:"files,omitempty"`
	ExcludeFromClassMap []string          `json:"exclude-from-classmap"`
}

type PackageSource struct {
	Type      string `json:"type"`
	Url       string `json:"url"`
	Reference string `json:"reference"`
	ShaSum    string `json:"shasum"`
}
