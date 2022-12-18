package model

var (
	RootResource = Resource{resourceId{"", "root"}, nil}
)

const (
	AggregateRelationship   = "Aggregation"
	CompositionRelationship = "Composition"
)

type resourceId struct {
	id   string
	kind string
}

type Resource struct {
	id         resourceId
	Attributes []Attribute
}

func NewResource(id, kind string) Resource {
	return Resource{
		id: resourceId{
			id:   id,
			kind: kind,
		},
	}
}

func (r Resource) Id() string {
	return r.id.id
}

func (r Resource) SetId(id string) {
	r.id.id = id
}

func (r Resource) Kind() string {
	return r.id.kind
}

func (r Resource) SetKind(kind string) {
	r.id.kind = kind
}

func (r Resource) Name() string {
	return r.id.kind + r.id.id
}
