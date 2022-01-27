package main

import (
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/zhaoyunxing/koanf/domain"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	k := koanf.New(".")

	if err := k.Load(file.Provider(absolutePath("./koanf/conf/application.yaml")), yaml.Parser()); err != nil {
		panic(err)
	}
	dev := koanf.New(".")
	if err := dev.Load(file.Provider(absolutePath("./koanf/conf/application-dev.yaml")), yaml.Parser()); err != nil {
		panic(err)
	}

	err := k.Merge(dev)
	if err != nil {
		panic(err)
	}

	k.Print()
	fmt.Println("===============")
	dev.Print()

	var conf domain.DubboConfig
	if err := k.UnmarshalWithConf("dubbo", &conf, koanf.UnmarshalConf{Tag: "yaml"}); err != nil {
		fmt.Println(err)
	}
	duration := k.Duration("dubbo.registries.nacos.timeout")
	fmt.Println(conf)
	fmt.Println(duration)
}

// absolutePath 获取绝对路径
func absolutePath(inPath string) string {

	if inPath == "$HOME" || strings.HasPrefix(inPath, "$HOME"+string(os.PathSeparator)) {
		inPath = userHomeDir() + inPath[5:]
	}

	if strings.HasPrefix(inPath, "$") {
		end := strings.Index(inPath, string(os.PathSeparator))

		var value, suffix string
		if end == -1 {
			value = os.Getenv(inPath[1:])
		} else {
			value = os.Getenv(inPath[1:end])
			suffix = inPath[end:]
		}

		inPath = value + suffix
	}

	if filepath.IsAbs(inPath) {
		return filepath.Clean(inPath)
	}

	p, err := filepath.Abs(inPath)
	if err == nil {
		return filepath.Clean(p)
	}

	return ""
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
