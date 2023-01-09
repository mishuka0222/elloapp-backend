package conf

type SmsVerifyCodeConfig struct {
	Name          string
	SendCodeUrl   string
	VerifyCodeUrl string
	Key           string
	Secret        string
	RegionId      string
}

type WebrtcConfig struct {
	Turn     bool
	Stun     bool
	Ip       string
	Ipv6     string
	Port     int
	Username string
	Password string
}
