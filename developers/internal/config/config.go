package config

type DeveloperServerConfig struct {
	WebServer      WebServerConfig      `json:"web_server"`
	DatabaseServer DatabaseServerConfig `json:"database_server"`
	Security       SecurityConfig       `json:"security"`
	LogLevel       string               `json:"log_level"`
}

type SecurityConfig struct {
	JwtKey             string `json:"jwt_key"`
	PasswordCostFactor int    `json:"password_cost_factor"`
}

type WebServerConfig struct {
	Port               string   `json:"port"`
	Host               string   `json:"host"`
	CORSOriginAllowed  []string `json:"cors_origin_allowed"`
	CORSAllowedHeaders []string `json:"cors_allowed_headers"`
	CORSAllowedMethods []string `json:"cors_allowed_methods"`
}

type DatabaseServerConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}
