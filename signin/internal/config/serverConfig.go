package config

type ServerConfig struct {
	WebServer WebServerConfig `json:"web_server"`
	Providers ProvidersConfig `json:"providers"`
	Security  SecurityConfig  `json:"security"`
}

type SecurityConfig struct {
	JwtKey string `json:"jwt_key"`
}
type WebServerConfig struct {
	Schema string `json:"schema"`
	Port   string `json:"server_port"`
}

type ProvidersConfig struct {
	UserServerUri      string `json:"user_server_uri"`
	DeveloperServerUri string `json:"developer_server_uri"`
}
