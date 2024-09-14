package mpwx

import (
	"fmt"
	"os"
)

type MpwxClient struct {
	appid         string
	secret        string
	plainTemplate string
	adminUser     string
}

func NewFromEnv(envPrefix string) MpwxClient {
	appid := os.Getenv(fmt.Sprintf("%s_APPID", envPrefix))
	secret := os.Getenv(fmt.Sprintf("%s_SECRET", envPrefix))
	plainTemplate := os.Getenv(fmt.Sprintf("%s_TEMPLATE_PLAIN", envPrefix))
	adminUser := os.Getenv(fmt.Sprintf("%s_USER_ADMIN", envPrefix))
	return MpwxClient{appid, secret, plainTemplate, adminUser}
}
