package providers

import "waservice/routes"

type routeService struct {
	Path string
}

func (r *routeService) Handle() {
	api := AppService.Group(r.Path)
	routes.Api(api)
}

func RouteService(p string) *routeService {
	return &routeService{
		Path: p,
	}
}
