package global

type Config struct {
	MySQL     *MySQLConfig     `json:"mysql"`
	Redis     *RedisConfig     `json:"redis"`
	AccessKey *AccessKeyConfig `json:"accessKey"`
}

type MySQLConfig struct {
	Ip       string `json:"ip"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type RedisConfig struct {
	Ip          string `json:"ip"`
	Port        int    `json:"port"`
	Password    string `json:"password"`
	Database    int    `json:"database"`
	MaxIdle     int    `json:"maxIdle"`      //最初的连接数量
	MaxActive   int    `json:"maxActive"`    //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
	IdleTimeout int    `json:"idleTimeout" ` //连接关闭时间 300秒 （300秒不使用自动关闭）
}

type AccessKeyConfig struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
}
