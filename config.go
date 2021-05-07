package ezconfig

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func LoadConfig(envPrefix string, cfg interface{}, paths []string) interface{} {
	viper.AddConfigPath(".")
	if (paths != nil) {
		for _, p := range paths {
			viper.AddConfigPath(p)
		}
	}
	viper.SetEnvPrefix("envPrefix")
	//viper.SetEnvKeyReplacer(strings.NewReplacer())
	viper.AutomaticEnv()

	// 우선 순위는 default example_config, ${{ envPrefix}}_ENVIRONMENT example_config 순이다.
	for _, environment := range []string{"default", os.Getenv(envPrefix + "_" + "ENVIRONMENT")} {
		if environment != "" {
			viper.SetConfigName(environment)
			err := viper.ReadInConfig()
			if err != nil {
				log.Println("[Warn] No config file found for " + environment +".")
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