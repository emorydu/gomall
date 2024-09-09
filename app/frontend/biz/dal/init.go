package dal

import (
	"github.com/emory/gomall/app/frontend/biz/dal/mysql"
	"github.com/emory/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
