package wss

type Routers []Router

var (
	MethodPost   = "POST"
	MethodGet    = "GET"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
)

type Router struct {
	Method  string
	Pattern string
	Action  Action
}
