package middleware_sample

import (
	"net/http"

	"github.com/roadrunner-server/api/v2/plugins/config"
	"github.com/roadrunner-server/errors"
	"go.uber.org/zap"
)

const name = "my_middleware_name"

type Plugin struct {
	log *zap.Logger
	cfg *Config
}

func (p *Plugin) Init(cfg config.Configurer, log *zap.Logger) error {
	// check if we need to init this middleware
	if !cfg.Has(name) {
		return errors.E(errors.Disabled)
	}

	// populate configuration
	p.cfg = &Config{}
	err := cfg.UnmarshalKey(name, p.cfg)
	if err != nil {
		return err
	}

	// init default values
	p.cfg.InitDefaults()

	// init logger
	p.log = new(zap.Logger)
	*p.log = *log

	return nil
}

// Middleware is our actual http middleware
func (p *Plugin) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		resp, err := http.Get("https://google.com")
		if err != nil {
			p.log.Error("third-party api call", zap.Error(err))
			return
		}

		p.log.Info("response", zap.Int("http_code", resp.StatusCode))

		next.ServeHTTP(w, r)
	})
}

func (p *Plugin) Name() string {
	return name
}
