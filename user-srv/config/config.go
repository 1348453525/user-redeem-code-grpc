package config

type Config struct {
	Server Server `mapstructure:"server" json:"server"`
	MySQL  MySQL  `mapstructure:"mysql" json:"mysql"`
	Redis  Redis  `mapstructure:"redis" json:"redis"`
	Consul Consul `mapstructure:"consul" json:"consul"`
	Nacos  Nacos  `mapstructure:"nacos" json:"nacos"`
}

type Server struct {
	Name    string   `mapstructure:"name" json:"name"`
	Tags    []string `mapstructure:"tags" json:"tags"`
	Version string   `mapstructure:"version" json:"version"`
	Mode    string   `mapstructure:"mode" json:"mode"`
	Addr    string   `mapstructure:"addr" json:"addr"`
	Port    int      `mapstructure:"port" json:"port"`
}

type MySQL struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         string `mapstructure:"port" json:"port"`
	User         string `mapstructure:"user" json:"user"`
	Password     string `mapstructure:"password" json:"password"`
	DB           string `mapstructure:"db" json:"db"`
	Charset      string `mapstructure:"charset" json:"charset"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns"`
}

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr"`
	Password string `mapstructure:"password" json:"password"`
}

type Consul struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type Nacos struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	Namespace string `mapstructure:"namespace" json:"namespace"`
	User      string `mapstructure:"user" json:"user"`
	Password  string `mapstructure:"password" json:"password"`
	DataId    string `mapstructure:"data_id" json:"data_id"`
	Group     string `mapstructure:"group" json:"group"`
}
