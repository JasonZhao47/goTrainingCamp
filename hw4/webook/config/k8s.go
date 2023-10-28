//go:build k8s

package config

var Config = config{
	DB: DBConfig{
		DSN: "root:root@tcp(webook-mysql-jasonzhao47:3308)/mysql",
	},
	Redis: RedisConfig{
		Addr: "webook-redis-jasonzhao47:6380",
	},
}
