package main

import (
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/zhaoyunxing/viper/config/root"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	k := koanf.New(".")
	getConfigPath("$HOME/conf/application.yaml")
	if err := k.Load(file.Provider("../conf"), yaml.Parser()); err != nil {
		panic(err)
	}

	var conf root.Config
	if err := k.Unmarshal("dubbo", &conf); err != nil {
		fmt.Println(err)
	}
	fmt.Println(conf)
}

func getConfigPath(in string) {
	//configPaths:=make([]string,)
	if in != "" {
		absin := absPathify(in)
		fmt.Println(absin)
		//if !stringInSlice(absin, v.configPaths) {
		//	v.configPaths = append(v.configPaths, absin)
		//}
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func absPathify(inPath string) string {

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
