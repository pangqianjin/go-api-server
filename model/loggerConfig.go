package model

type LoggerConfig struct {
	Status bool `json:"status"`
	IP     bool `json:"ip"`
	Method bool `json:"method"`
	Path   bool `json:"path"`
	Query  bool `json:"query"`
}
