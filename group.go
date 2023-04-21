package ctp

// Group router group, a Group is associated with a prefix and parent group
type Group struct {
	engine      *Engine
	prefix      string
	parent      *Group
	middlewares []HandlerFunc // 存放中间件
}

func NewGroup(e *Engine, prefix string) *Group {
	return &Group{
		engine: e,
		prefix: prefix,
		parent: nil,
	}
}

func (g *Group) getTotalPath() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getTotalPath() + g.prefix
}

func (g *Group) POST(uri string, handlers ...HandlerFunc) {
	uri = g.getTotalPath() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.engine.POST(uri, allHandlers...)
}

func (g *Group) PUT(uri string, handlers ...HandlerFunc) {
	uri = g.getTotalPath() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.engine.PUT(uri, allHandlers...)
}

func (g *Group) DELETE(uri string, handlers ...HandlerFunc) {
	uri = g.getTotalPath() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.engine.DELETE(uri, allHandlers...)
}

func (g *Group) GET(uri string, handlers ...HandlerFunc) {
	uri = g.getTotalPath() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.engine.GET(uri, allHandlers...)
}

func (g *Group) getMiddlewares() []HandlerFunc {
	if g.parent == nil {
		return g.middlewares
	}

	return append(g.parent.getMiddlewares(), g.middlewares...)
}
func (g *Group) Use(middlewares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}
