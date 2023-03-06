package config

type Jwt struct {
	Secret  string `json:"secret,omitempty" yaml:"secret"`   //密钥
	Expires int    `json:"expires,omitempty" yaml:"expires"` //过期时间
	Issuer  string `json:"issuer,omitempty" yaml:"issuer"`   //颁发人
}
