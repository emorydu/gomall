package dal

import (
	"github.com/emory/gomall/app/user/biz/dal/mysql"
	"github.com/emory/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
