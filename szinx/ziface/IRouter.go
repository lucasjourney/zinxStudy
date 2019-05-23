package ziface

type IRouter interface {
	PreHandle(r IRequest)
	Handle(r IRequest)
	PostHandle(r IRequest)
}
