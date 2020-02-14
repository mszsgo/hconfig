package hconfig

// 应用配置
type App struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`

	Host int64 `json:"host"`
	Port int64 `json:"port"`

	LogLevel string `json:"logLevel"`
}
