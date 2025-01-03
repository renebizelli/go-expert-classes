package configs

import (
	"fmt"
	"reflect"
	pkg_utils "renebizelli/api/pkg/utils"
	"slices"
	"strconv"
	"strings"

	"github.com/go-chi/jwtauth"

	"github.com/spf13/viper"
)

type (
	Config struct {
		DB        DBConfig
		WebServer WebServerConfig
		JWT       JWTConfig
	}

	DBConfig struct {
		Driver   string `mapstructure:"DB_DRIVER"`
		Host     string `mapstructure:"DB_HOST"`
		Port     string `mapstructure:"DB_PORT"`
		Name     string `mapstructure:"DB_NAME"`
		User     string `mapstructure:"DB_USER"`
		Password string `mapstructure:"DB_PASSWORD"`
	}

	WebServerConfig struct {
		Port string
	}

	JWTConfig struct {
		Secret    string
		ExpiresIn int64
		Token     *jwtauth.JWTAuth
	}
)

func LoadConfig(path string) Config {

	viper.SetConfigName("app_config")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv() // obtem valores das variáveis de ambiente, sobrepondo valores do .env

	err := viper.ReadInConfig()

	pkg_utils.PanicIfError(err, "Load config file error")

	var config Config

	valueRef := reflect.ValueOf(&config)
	typeRef := reflect.TypeOf(&config)

	skipProperties := []string{"jwt_token"}

	for i := 0; i < typeRef.Elem().NumField(); i++ {

		mainProperty := typeRef.Elem().Field(i)

		for j := 0; j < mainProperty.Type.NumField(); j++ {

			property := mainProperty.Type.Field(j)

			envPropertyName := strings.ToLower(strings.Join([]string{mainProperty.Name, property.Name}, "_"))

			if slices.Contains(skipProperties, envPropertyName) {
				continue
			}

			value := viper.Get(envPropertyName)

			if value != nil {

				refValueProperty := valueRef.Elem().FieldByName(mainProperty.Name).FieldByName(property.Name)

				switch property.Type.Kind() {
				case reflect.String:
					refValueProperty.SetString(value.(string))
				case reflect.Int64:
					v, e := strconv.ParseInt(value.(string), 0, 64)
					pkg_utils.PanicIfError(e, fmt.Sprintf("Invalid value for key: %s, value: %s", envPropertyName, value))
					refValueProperty.SetInt(v)
				}
			} else {
				fmt.Println("**************************************************")
				fmt.Printf("Property %s is nil\n", envPropertyName)
				fmt.Println("**************************************************")
			}
		}
	}

	config.JWT.Token = jwtauth.New("HS256", []byte(config.JWT.Secret), nil)

	return config
}
