package api

type Package interface {
	InitRoutes(r *Router)
}

type Config struct {
}

type Router struct {
}

func New(apiPkg Package, cfg *Config) *Router {
	return nil
}
