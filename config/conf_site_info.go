package config

type SiteInfo struct {
	CreateAt    string `yaml:"create_at" json:"create_at,omitempty"`
	BeiAn       string `yaml:"bei_an" json:"bei_an,omitempty"`
	Title       string `yaml:"title" json:"title,omitempty"`
	QQImage     string `yaml:"qq_image" json:"qq_image,omitempty"`
	Version     string `yaml:"version" json:"version,omitempty"`
	Email       string `yaml:"email" json:"email,omitempty"`
	WechatImage string `yaml:"wechat_image" json:"wechat_image,omitempty"`
	Name        string `yaml:"name" json:"name,omitempty"`
	Job         string `yaml:"job" json:"job,omitempty"`
	Addr        string `yaml:"addr" json:"addr,omitempty"`
	Slogan      string `yaml:"slogan" json:"slogan,omitempty"`
	SloganEn    string `yaml:"slogan_en" json:"slogan_en,omitempty"`
	Web         string `yaml:"web" json:"web,omitempty"`
	BiliBiliUrl string `yaml:"bilibili_url" json:"bili_bili_url,omitempty"`
	GiteeUrl    string `yaml:"gitee_url" json:"gitee_url,omitempty"`
	GithubUrl   string `yaml:"github_url" json:"github_url,omitempty"`
}
