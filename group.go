package ctp

// Group router group, a Group is associated with a prefix and parent group
type Group struct {
	engine *Engine
	prefix string
	parent *Group
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

func (g *Group) POST(uri string, handler HandlerFunc) {
	uri = g.getTotalPath() + uri
	g.engine.POST(uri, handler)
}

func (g *Group) PUT(uri string, handler HandlerFunc) {
	uri = g.getTotalPath() + uri
	g.engine.PUT(uri, handler)
}

func (g *Group) DELETE(uri string, handler HandlerFunc) {
	uri = g.getTotalPath() + uri
	g.engine.DELETE(uri, handler)
}

func (g *Group) GET(uri string, handler HandlerFunc) {
	uri = g.getTotalPath() + uri
	g.engine.GET(uri, handler)
}
