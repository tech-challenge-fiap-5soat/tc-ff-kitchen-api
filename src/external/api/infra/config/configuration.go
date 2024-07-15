package config

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"

	"github.com/integralist/go-findroot/find"

	"github.com/spf13/viper"
)

var (
	runOnce sync.Once
	config  *Config
)

type Config struct {
	MongoCfg MongoConfig `mapstructure:"mongodb"`
	ApiCfg   Api         `mapstructure:"api"`
}

type MongoConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Pass     string `mapstructure:"pass"`
}

type Api struct {
	Port                 string     `mapstructure:"port"`
	AuthConfig           AuthConfig `mapstructure:"authconfig"`
	AuthorizationBaseUrl string     `mapstructure:"authorizationurl"`
	OrderApiBaseURL      string     `mapstructure:"orderApiBaseURL"`
}

type AuthConfig struct {
	UserPoolId string `mapstructure:"userpoolid"`
	ClientId   string `mapstructure:"clientid"`
	TokenUse   string `mapstructure:"tokenuse"`
}

func init() {
	config = setupConfig()
}

func GetMongoCfg() MongoConfig {
	return config.MongoCfg
}

func GetApiCfg() Api {
	return config.ApiCfg
}

func setupConfig() *Config {
	runOnce.Do(func() {
		var appConfig Config

		root, _ := find.Repo()
		configFilePath := path.Join(root.Path, "/src/external/api/infra/config")

		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()
		viper.SetConfigName("configs")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(configFilePath)
		viper.AddConfigPath("/home/hugofranciscofagundesdesouza/Projects/pos-graduacao/tc-ff-kitchen-api/src/external/api/infra/config")
		viper.AddConfigPath("/app/data/configs")
		err := viper.ReadInConfig()

		//if err = viper.ReadInConfig(); err != nil {
		//	panic(fmt.Errorf("Falha ao carregar as configurações: %w \n", err))
		//}

		if err != nil && !allConfigsAreSet() {
			panic(fmt.Errorf("falha ao carregar as configurações: %w", err))
		}

		for _, key := range viper.AllKeys() {
			value := viper.GetString(key)
			envOrRaw := replaceEnvInConfig([]byte(value))
			viper.Set(key, string(envOrRaw))
		}

		if err == nil {
			err := viper.Unmarshal(&appConfig)
			if err != nil {
				panic(err)
			}
		}

		if allConfigsAreSet() { // load envs from infra
			appConfig.ApiCfg.Port = viper.Get("api.port").(string)
			appConfig.MongoCfg.Host = viper.Get("mongodb.host").(string)
			appConfig.MongoCfg.Database = viper.Get("mongodb.database").(string)
			appConfig.MongoCfg.User = viper.Get("mongodb.user").(string)
			appConfig.MongoCfg.Pass = viper.Get("mongodb.pass").(string)
			appConfig.ApiCfg.AuthConfig.UserPoolId = viper.Get("api.authConfig.userPoolId").(string)
			appConfig.ApiCfg.AuthConfig.ClientId = viper.Get("api.authConfig.clientId").(string)
			appConfig.ApiCfg.AuthConfig.TokenUse = viper.Get("api.authConfig.tokenUse").(string)
			appConfig.ApiCfg.AuthorizationBaseUrl = viper.Get("api.authorizationUrl").(string)

		}

		config = &appConfig
	})

	return config
}

func allConfigsAreSet() bool {
	return viper.Get("mongodb.host") != nil &&
		viper.Get("mongodb.database") != nil &&
		viper.Get("mongodb.user") != nil &&
		viper.Get("api.port") != nil &&
		viper.Get("mongodb.pass") != nil &&
		viper.Get("api.authConfig.userPoolId") != nil &&
		viper.Get("api.authConfig.clientId") != nil &&
		viper.Get("api.authConfig.tokenUse") != nil &&
		viper.Get("api.authorizationUrl") != nil
}

func replaceEnvInConfig(body []byte) []byte {
	search := regexp.MustCompile(`\$\{([^{}]+)\}`)
	replacedBody := search.ReplaceAllFunc(body, func(b []byte) []byte {
		group1 := search.ReplaceAllString(string(b), `$1`)

		envValue := os.Getenv(group1)
		if len(envValue) > 0 {
			return []byte(envValue)
		}
		return []byte("")
	})
	return replacedBody
}
