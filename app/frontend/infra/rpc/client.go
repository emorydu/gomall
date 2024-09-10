package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/emory/gomall/app/frontend/conf"
	"github.com/emory/gomall/app/frontend/utils"
	"github.com/emory/gomall/rpc_gen/kitex_gen/user/userservice"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	UserClient userservice.Client

	once sync.Once
)

func Init() {
	once.Do(func() {
		iniUserClient()
	})
}

func iniUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	utils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	utils.MustHandleError(err)
}
