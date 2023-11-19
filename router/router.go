package router

import "github.com/gin-gonic/gin"

type Router interface {
	Route(r *gin.Engine)
}

var routers []Router

type RegisterRouter struct {
}

func (rg *RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

func InitRouter(r *gin.Engine) {
	for _, router := range routers {
		router.Route(r)
	}
}

func Register(ro Router) {
	routers = append(routers, ro)
}
