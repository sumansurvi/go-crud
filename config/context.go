package config

type AppContext struct {
	Config *Config
}

func NewAppContext(config *Config) *AppContext {
	return &AppContext{
		Config: config,
	}
}
