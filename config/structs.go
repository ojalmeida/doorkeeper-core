package config

type Configuration struct {
	Timeout struct {
		Read       int64 `json:"read,omitempty"`
		Write      int64 `json:"write,omitempty"`
		Processing int64 `json:"processing,omitempty"`
	} `json:"timeout"`

	Logging struct {
		Info struct {
			Path   string `json:"path,omitempty"`
			Prefix string `json:"prefix,omitempty"`
		} `json:"info"`

		Warn struct {
			Path   string `json:"path,omitempty"`
			Prefix string `json:"prefix,omitempty"`
		} `json:"warn"`

		Error struct {
			Path   string `json:"path,omitempty"`
			Prefix string `json:"prefix,omitempty"`
		} `json:"error"`
	} `json:"logging"`
}
