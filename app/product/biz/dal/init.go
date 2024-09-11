package dal

import (
	"github.com/emory/gomall/app/product/biz/dal/mysql"
	"github.com/emory/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
