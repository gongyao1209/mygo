package config

const env = "local"

func Get(item string) string {
	type confMap map[string]string

	conf := make(map[string]confMap)

	conf["local"] = map[string]string{
		"Port": "9002",
		"Mysql": "root:******@tcp(139.198.5.192:3306)/xiangxin0828?charset=utf8",
		"Redis": "119:6379",
	}

	return  conf[env][item]
}