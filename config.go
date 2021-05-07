package ezconfig

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
)

func LoadConfig(envPrefix string, cfg interface{}) interface{} {
	viper.AddConfigPath(".")
	viper.AddConfigPath(path.Join(os.Getenv(envPrefix + "_PROJECT_ROOT"), "config"))
	viper.SetEnvPrefix("EZ")
	//viper.SetEnvKeyReplacer(strings.NewReplacer())
	viper.AutomaticEnv()

	// 우선 순위는 default config, ${{ envPrefix}}_ENVIRONMENT config 순이다.
	for _, environment := range []string{"default", os.Getenv(envPrefix + "_" + "ENVIRONMENT")} {
		if environment != "" {
			viper.SetConfigName(environment)
			err := viper.ReadInConfig()
			if err != nil {
				log.Println("[Warn] No default config file found.")
			} else{
				err = viper.Unmarshal(cfg)
				if err != nil {
					log.Fatal(err)
				}
			}


		} else{
			// 생략.
			// 그냥 default config만 이용하게 됨.
		}
	}

	return cfg
}