package config

type DeveloperServerConfig struct {
	WebServer      WebServerConfig      `json:"web_server"`
	DatabaseServer DatabaseServerConfig `json:"database_server"`
	Security       SecurityConfig       `json:"security"`
}

type SecurityConfig struct {
	JwtKey             string `json:"jwt_key"`
	PasswordCostFactor int    `json:"password_cost_factor"`
}
type WebServerConfig struct {
	Schema  string `json:"schema"`
	Address string `json:"server_address"`
	Port    string `json:"server_port"`
}

type DatabaseServerConfig struct {
	Address      string `json:"server_address"`
	Port         string `json:"server_port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
}
