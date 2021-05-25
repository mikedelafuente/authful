package config

type ServerConfig struct {
	WebServer WebServerConfig `json:"web_server"`
	Services  ServicesConfig  `json:"services"`
	Security  SecurityConfig  `json:"security"`
}

type SecurityConfig struct {
	JwtKey string `json:"jwt_key"`
}
type WebServerConfig struct {
	Schema  string `json:"schema"`
	Address string `json:"server_address"`
	Port    string `json:"server_port"`
}

type ServicesConfig struct {
	UserServerUri      string `json:"user_server_uri"`
	DeveloperServerUri string `json:"developer_server_uri"`
}
