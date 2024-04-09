package mymod

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"go.uber.org/zap"
)

func init() {
	caddy.RegisterModule(MyMod{})
	httpcaddyfile.RegisterHandlerDirective("my_mod", parseMyModCaddyfile)
}

func parseMyModCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m MyMod
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}

type MyMod struct {
	HeaderValue string      `json:"header_value"`
	logger      *zap.Logger `json:"-"`
}

func (m MyMod) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.my_mod",
		New: func() caddy.Module { return new(MyMod) },
	}
}

// UnmarshalCaddyfile parses the my_mod directive. It enables
// the handler and configures it with this syntax:
//
//	my_mod <header_value>
func (m *MyMod) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	d.Next() // Consume directive name

	if !d.Args(&m.HeaderValue) {
		return d.ArgErr()
	}

	return nil
}

func (m *MyMod) Provision(ctx caddy.Context) error {
	m.logger = ctx.Logger()
	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m MyMod) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	w.Header().Set("X-Hello-World", m.HeaderValue)

	m.logger.Info("added header to response",
		zap.String("header_key", "X-Hello-World"),
		zap.String("header_value", m.HeaderValue),
	)

	next.ServeHTTP(w, r)

	return nil
}

var (
	_ caddy.Module                = (*MyMod)(nil)
	_ caddyhttp.MiddlewareHandler = (*MyMod)(nil)
	_ caddyfile.Unmarshaler       = (*MyMod)(nil)
	_ caddy.Provisioner           = (*MyMod)(nil)
)
