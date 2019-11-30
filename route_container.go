package moocro

// RouteContainer maintains all the routes
type RouteContainer interface {
	All() []Route
	Find(path string) Route
	Route(route Route)
}

type mapRouteContainer struct {
	routes map[string]Route
}

// CreateRouteContainer to manage routes
func CreateRouteContainer() RouteContainer {
	return &mapRouteContainer{routes: make(map[string]Route)}
}

// All the routes in the container
func (c *mapRouteContainer) All() []Route {
	v := make([]Route, 0, len(c.routes))

	for _, value := range c.routes {
		v = append(v, value)
	}

	return v
}

// Find the route by a path
func (c *mapRouteContainer) Find(path string) Route {
	for key, value := range c.routes {
		if key == path {
			return value
		}
	}

	return nil
}

// Route adds a route to the container
func (c *mapRouteContainer) Route(route Route) {
	c.routes[route.Path()] = route
}
