package net

import "zinxStudy/szinx/ziface"

type BaseRouter struct {

}


func (r *BaseRouter) PreHandle(request ziface.IRequest){}
func (r *BaseRouter) Handle(request ziface.IRequest){}
func (r *BaseRouter) PostHandle(request ziface.IRequest){}
