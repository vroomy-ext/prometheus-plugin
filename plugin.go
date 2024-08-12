package plugin

import (
	"log"

	"net/http"

	"github.com/vroomy/httpserve"
	"github.com/vroomy/vroomy"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	var (
		p   Plugin
		err error
	)

	if err = vroomy.Register("prometheus", &p); err != nil {
		log.Fatal(err)
	}
}

type Plugin struct {
	vroomy.BasePlugin

	h http.HandlerFunc
}

// Load will be called by vroomy on initialization
func (p *Plugin) Load(env vroomy.Environment) (err error) {
	p.h = promhttp.Handler().ServeHTTP
	return
}

func (p *Plugin) Metrics(ctx *httpserve.Context) {
	p.h.ServeHTTP(ctx.Writer(), ctx.Request())
}
