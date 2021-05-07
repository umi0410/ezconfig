package ezconfig

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type ExampleConfig struct {
	Environment string
	Api struct{
		Root string
		Auth struct{
			Type string
			Basic struct{
				Username string
				Password string
			}
			Jwt struct{
				Token string
			}
		}
		Foo string
	}
	Foo string
}

func TestDefaultConfig(t *testing.T) {
	t.Log("EZ_ENVIRONMENT: \"\"")
	os.Setenv("EZ_ENVIRONMENT", "")
	cfg := &ExampleConfig{}
	LoadConfig("EZ", cfg, []string{"./example_config"})

	assert.Equal(t, "default", cfg.Environment)
	assert.Equal(t, "basic", cfg.Api.Auth.Type)
	assert.Equal(t, "admin", cfg.Api.Auth.Basic.Username)
	assert.Equal(t, "123123", cfg.Api.Auth.Basic.Password)
}

func TestLocalConfig(t *testing.T) {
	t.Log("EZ_ENVIRONMENT: local")
	os.Setenv("EZ_ENVIRONMENT", "local")
	cfg := &ExampleConfig{}
	LoadConfig("EZ", cfg, []string{"./example_config"})

	// cfg.Environment는 환경변수를 따라 오버라이딩됨.
	assert.Equal(t, "local", cfg.Environment)
	assert.Equal(t, "basic", cfg.Api.Auth.Type)
	// default config의 username을 그대로 사용해야함.
	assert.Equal(t, "admin", cfg.Api.Auth.Basic.Username)
	// password는 override됨.
	assert.Equal(t, "simple local password", cfg.Api.Auth.Basic.Password)
}

func TestDevConfig_UsingBasic(t *testing.T) {
	t.Log("EZ_ENVIRONMENT: dev")
	t.Log("EZ_API_AUTH_TYPE: basic")
	os.Setenv("EZ_ENVIRONMENT", "dev")
	os.Setenv("EZ_API.AUTH.TYPE", "basic")
	cfg := &ExampleConfig{}
	LoadConfig("EZ", cfg, []string{"./example_config"})

	// cfg.Environment는 환경변수를 따라 오버라이딩됨.
	assert.Equal(t, "dev", cfg.Environment)
	assert.Equal(t, "basic", cfg.Api.Auth.Type)
	assert.Equal(t, "admin", cfg.Api.Auth.Basic.Username)
	assert.Equal(t, "complex dev password", cfg.Api.Auth.Basic.Password)
}

func TestDevConfig_UsingJwt(t *testing.T) {
	t.Log("EZ_ENVIRONMENT: dev")
	t.Log("EZ_API_AUTH_TYPE: jwt")
	os.Setenv("EZ_ENVIRONMENT", "dev")
	os.Setenv("EZ_API.AUTH.TYPE", "jwt")
	cfg := &ExampleConfig{}
	LoadConfig("EZ", cfg, []string{"./example_config"})

	// cfg.Environment는 환경변수를 따라 오버라이딩됨.
	assert.Equal(t, "dev", cfg.Environment)
	assert.Equal(t, "jwt", cfg.Api.Auth.Type)
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIifQ.UIZchxQD36xuhacrJF9HQ5SIUxH5HBiv9noESAacsxU", cfg.Api.Auth.Jwt.Token)
}