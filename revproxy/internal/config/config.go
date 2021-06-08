package config

type ServerConfig struct {
	WebServer WebServerConfig `json:"web_server"`
	Security  SecurityConfig  `json:"security"`
	LogLevel  string          `json:"log_level"`
}

type ProxyConfig struct {
	ProxyMaps []ProxyMap `json:"proxy_mappings"`
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

type ProxyMap struct {
	Name           string      `json:"name"`
	ServiceBaseUrl string      `json:"service_base_url"`
	Paths          []ProxyPath `json:"paths"`
}

type ProxyPath struct {
	Path     string `json:"path"`
	IsSecure bool   `json:"is_secure"`
}
