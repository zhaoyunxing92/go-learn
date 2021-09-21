package parsers

import (
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/file"
	"github.com/zhaoyunxing/koanf/domain"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	k := koanf.New("=")

	if err := k.Load(file.Provider(absolutePath("../conf/application.properties")), Parser()); err != nil {
		panic(err)
	}
	var conf = domain.DubboConfig{}
	if err := k.UnmarshalWithConf("dubbo", &conf, koanf.UnmarshalConf{Tag: "yaml"}); err != nil {
		fmt.Println(err)
	}

	fmt.Println(conf.Application.Name)
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
