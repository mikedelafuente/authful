package config

type UserServerConfig struct {
	WebServer      WebServerConfig      `json:"web_server"`
	DatabaseServer DatabaseServerConfig `json:"database_server"`
	Security       SecurityConfig       `json:"security"`
}

type SecurityConfig struct {
	JwtKey             string `json:"jwt_key"`
	PasswordCostFactor int    `json:"password_cost_factor"`
}
type WebServerConfig struct {
	Port string `json:"port"`
}

type DatabaseServerConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
