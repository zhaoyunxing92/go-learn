package config

import (
	"flag"
	"fmt"
)
import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/zhaoyunxing/viper/config/root"
)

// config config
type config struct {
	// prefix default dubbo
	prefix string
	// config file name default application
	name string
	// config file type default yaml
	genre string
	// config file path default ./conf
	path string
}
type optionFunc func(*config)

func (fn optionFunc) apply(vc *config) {
	fn(vc)
}

type Option interface {
	apply(vc *config)
}

func Load(opts ...Option) *root.Config {
	// pares CommandLine
	parseCommandLine()

	conf := &config{
		viper.GetString("prefix"),
		viper.GetString("name"),
		viper.GetString("genre"),
		viper.GetString("path"),
	}

	for _, opt := range opts {
		opt.apply(conf)
	}

	v := viper.New()
	v.AddConfigPath(conf.path)
	v.SetConfigName(conf.name)
	v.SetConfigType(conf.genre)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	var bc root.Config
	if err := v.UnmarshalKey(conf.prefix, &bc); err != nil {
		fmt.Println(err)
	}
	bc.SetViper(v)

	validate := validator.New()
	uni := translator.New(en.New(), zh.New())
	trans, _ := uni.GetTranslator("zh")
	_ = zh_trans.RegisterDefaultTranslations(validate, trans)

	bc.SetValidate(validate)
	bc.SetTranslator(trans)

	return &bc
}

//parseCommandLine parse command line
func parseCommandLine() {
	flag.String("prefix", "dubbo", "config file prefix key")
	flag.String("name", "application.yaml", "config file name")
	flag.String("genre", "yaml", "config file type")
	flag.String("path", "./conf", "config file path default")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}
}

func WithGenre(genre string) Option {
	return optionFunc(func(conf *config) {
		conf.genre = genre
	})
}

func WithPath(path string) Option {
	return optionFunc(func(conf *config) {
		conf.path = path
	})
}

func WithPrefix(prefix string) Option {
	return optionFunc(func(conf *config) {
		conf.prefix = prefix
	})
}

func WithName(name string) Option {
	return optionFunc(func(conf *config) {
		conf.name = name
	})
}
