# ezconfig - Go application에서 configuration 편하게 관리하기

Golang은 마이크로서비스를 개발하기에 아주 좋은 언어입니다. 하지만 각 마이크로서비스를 개발할 때마다 configuration을 설정하는
작업을 위해 비슷한 코드를 매번 작성하는 것은 아주 성가십니다. `ezconfig` 는 그러한 불편을 덜어주고 Go application에 대한 configuration을 설정할 때
Easy하게 작업할 수 있도록 해줍니다.

`ezconfig` 는 Go application에서의 configuration solution으로 가장 유명한 라이브러리 중 하나인 [viper](https://github.com/spf13/viper) 을
바탕으로 작성되어있습니다. viper를 한 번 wrapping하여 좀 더 심플하고 쉽게 configuration할 수 있도록 했습니다.

(*아직은 단순히 개인적으로 마이크로서비스 개발할 때 이용하기 위한 라이브러리로서 개발 중... 따라서 완전하지 않음.*)

# Getting started

```
package main

import "github.com/umi0410/ezconfig"

type SimpleExample struct{
    Foo string
}

func ExampleUsingEnv() {
	cfg := &SimpleExample{}
	os.Setenv("SIMPLE_FOO", "bar")
	ezconfig.Load("SIMPLE", cfg)
	fmt.Println(cfg.Foo) // stdout: bar
}

func ExampleUsingConfigFile() {
	cfg := &SimpleExample{}
	// this example suppose you have config/default.yaml in your working directory
	ezconfig.Load("SIMPLE", cfg)
	fmt.Println(cfg.Foo) // stdout: bar
}
/**
# configuration file should be like this
# config/default.yaml
foo: bar
*/
```

# 사용 방법

`ezconfig` 는 앞서 말씀드렸듯이 [viper](https://github.com/spf13/viper) 을 바탕으로 작성되었기 때문에 어느 정도 viper의 사용법을 인지하고 있으면 더 편리합니다.
간단한 사용 방법을 알려드릴테니 자세한 예시는 예시 코드를 참고해주세요.

* **자신만의 Config struct 정의하기** - `ezconfig.LoadConfig(envPrefix string, cfg interface{}, paths []string)`에서 참조할 수 있는 사용자 정의 config struct 타입의 변수를 cfg 변수에 전달하여 configuration을 적용합니다.
  viper는 `map[string]string`을 지원하기도 하고 `struct` 를 지원하기도 하는 것으로 알고 있습니다.
  LoadConfig 후에는 cfg로 참조하는 값의 필드들에 configuration이 적용됩니다.
  
* **config file을 이용해 설정하기** - viper는 yaml, json을 비롯한 다양한 설정파일을 지원합니다. `{{ paths 속 각 path }}/default.{{ 다양한 file 확장자}}` 의 경로에
  configruation file을 위치시킨 뒤 `ezconfig.LoadConfig(...)` 를 통해 이용하십시오.
  * 기본적으로 default config가 로드 되고 `${{your_env_prefix}}_ENVIRONMENT` 가 의미하는 config가 로드되어 필요한 부분만을 default를 바탕으로 오버라이드합니다. 
  
  * ex) `MYCONFIG_ENVIRONMENT=local` 로 설정 후 `config/local.yaml` 에 설정 기입 후 `ezconfig.LoadConfig("MYCONFIG", cfgPtr, []string{"./config"})`
  
  * ex) `ENVIRONMENT=dev` 로 설정 후 `config/dev.yaml` 에 설정 기입 후 `ezconfig.LoadConfig("", cfgPtr, []string{"./config"})`
  
  * ex) 아무것도 설정 없이 그냥 `config/default.yaml` 에 설정 기입 후 `ezconfig.LoadConfig("", cfgPtr, []string["./config"])`
  
* **environment variable을 이용해 설정하기** - config file보다는 environment variable이 우선순위를 가지므로 config file의 내용을 오버라이드합니다.
  이때 몇 가지 규칙이 있는데 현재는 다음과 같습니다.
  
  * 환경변수 네이밍 룰: `{{ your env prefix}}_UPPER_CASE_KEY.NESTED.KEY.WITH.DOT`
  
  * envPrefix(당신의 앱 설정 관련 환경 변수에 붙일 prefix)를 정의하는 경우 `{{ env prefix}}_` 로 시작하는 환경변수만을 적용합니다.
  
  * config file에서 nested 된 객체의 key로 표현된 경우 `.` 으로 depth를 구분합니다.
  
  * key의 이름은 모두 UPPER_CASE로 표현합니다.


# 사용된 곳

* [khumu-comment](https://github.com/khu-dev/khumu-comment) - 경희대학교 커뮤니티 서비스의 댓글 서비스
