// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	home "github.com/emory/gomall/app/frontend/biz/router/home"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	home.Register(r)
}
