package providers

import "waservice/databases"

func RedisHandle() {
	databases.RedisRun()
	databases.MysqlRun()
}
