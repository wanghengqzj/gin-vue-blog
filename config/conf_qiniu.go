package config

type QiNiu struct {
	AccessKey string  `json:"access_key,omitempty" yaml:"access_key"`
	SecretKey string  `json:"secret_key,omitempty" yaml:"secret_key"`
	Bucket    string  `json:"bucket,omitempty" yaml:"bucket"` //存储桶的名字
	CDN       string  `json:"cdn,omitempty" yaml:"cdn"`       //访问图片的地址的前缀
	Zone      string  `json:"zone,omitempty" yaml:"zone"`     //存储的地区 华东，华北。。。
	Size      float64 `json:"size,omitempty" yaml:"size"`     //存储的大小限制，单位是MB
}
