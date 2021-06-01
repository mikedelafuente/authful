package config

type ServerConfig struct {
	WebServer WebServerConfig `json:"web_server"`
	Providers ProvidersConfig `json:"providers"`
	Security  SecurityConfig  `json:"security"`
	LogLevel  string          `json:"log_level"`
}

type SecurityConfig struct {
	JwtKey string `json:"jwt_key"`
}

type WebServerConfig struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

type ProvidersConfig struct {
	UserServerUri      string `json:"user_server_uri"`
	DeveloperServerUri string `json:"developer_server_uri"`
}
