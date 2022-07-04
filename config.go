package middleware_sample

type Config struct {
	Say string `mapstructure:"say_smt"`
}

func (c *Config) InitDefaults() {
	if c.Say == "" {
		c.Say = "default value"
	}
}
