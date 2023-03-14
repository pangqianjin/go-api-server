package model

type Property struct {
	Port         string `json:"port"`
	DataPath     string `json:"data_path"`
	PublicPath   string `json:"public_path"`
	PublicPrefix string `json:"public_prefix"`
	DebugLevel   string `json:"debug_level"`
}
