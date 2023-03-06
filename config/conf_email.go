package config

type Email struct {
	Host             string `json:"host,omitempty" yaml:"host"`
	Port             int    `json:"port,omitempty" yaml:"port"`
	User             string `json:"user,omitempty" yaml:"user"` //发件人邮箱
	Password         string `json:"password,omitempty" yaml:"password"`
	DefaultFromEmail string `json:"default_from_email,omitempty" yaml:"default_from_email"` //默认的发件人名字
	UseSSL           bool   `json:"use_ssl,omitempty" yaml:"use_ssl"`                       //是否使用ssl
	UserTls          bool   `json:"user_tls,omitempty" yaml:"user_tls"`                     //
}
