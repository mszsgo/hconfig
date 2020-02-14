package hconfig

type MongoConfig struct {
	ConnectionString string `json:"connectionString"`
}

// hmgdb接收的参数为mongodb连接字符串
func GetMongoConnectionString() string {
	// 获取DB连接
	var mc *MongoConfig
	NowConfig(nil, "mongo", &mc)
	return mc.ConnectionString
}
