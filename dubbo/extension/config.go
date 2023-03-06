package extension

import (
	"math"
	"reflect"
	"sort"
	"strconv"
	strings "strings"

	"github.com/spf13/viper"
)

type Config interface {
	Prefix() string
}

type Init interface {
	Init(key string)
}

type Order interface {
	Order() int
}

type Definition struct {
	prefix string
	order  int
	config Config
}

var (
	events = make([]Definition, 0, 16)
	keys   = make(map[string]struct{}, 0)
)

func GetDefinitions() []Definition {
	sort.Slice(events, func(i, j int) bool {
		return events[i].order < events[j].order
	})
	return events
}

func Register(config Config) {
	analysis(config, "")
}

func AddDefinitions(def Definition) {
	if _, ok := keys[def.prefix]; !ok {
		events = append(events, def)
		keys[def.prefix] = struct{}{}
	}
}

func Load(path, name string, config Config) (err error) {
	defer func() { load() }()

	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AllSettings()

	if err = viper.ReadInConfig(); err != nil {
		return err
	}
	if err = viper.UnmarshalKey(config.Prefix(), config); err != nil {
		return err
	}
	analysis(config, "")
	return nil
}

func analysis(config Config, key string) {

	tp := reflect.TypeOf(config)
	sv := reflect.ValueOf(config)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
		sv = sv.Elem()
	}
	if !sv.IsValid() {
		return
	}
	prefix := analyzePrefix(config.Prefix(), key)
	def := Definition{prefix: prefix, config: config}
	if _, ok := config.(Init); ok {
		if order, ok := config.(Order); ok {
			def.order = order.Order()
		} else {
			def.order = math.MaxInt
		}
		AddDefinitions(def)
	}
	for i := 0; i < tp.NumField(); i++ {
		field := tp.Field(i)
		kind := field.Type.Kind()
		val := sv.Field(i)
		switch kind {
		case reflect.Ptr:
			if c, ok := val.Interface().(Config); ok {
				analysis(c, "")
			}
		case reflect.Map:
			analyzeMap(val, "")
		case reflect.Slice:
			analyzeSlice(val, "")
		}
	}
}

func analyzePrefix(pre string, key string) string {
	if len(pre) > 0 && len(key) > 0 {
		return strings.Join([]string{pre, key}, ".")
	}
	if len(pre) > 0 {
		return pre
	}
	if len(key) > 0 {
		return key
	}
	return ""
}

// analyzeCollector analyze collector
func analyzeMap(value reflect.Value, key string) {
	for _, k := range value.MapKeys() {
		m := value.MapIndex(k)
		kind := m.Kind()
		suffix := analyzePrefix(k.String(), key)
		switch kind {
		case reflect.Ptr:
			if c, ok := m.Interface().(Config); ok {
				analysis(c, analyzePrefix(key, k.String()))
			}
		case reflect.Map:
			analyzeMap(m, suffix)
		case reflect.Slice:
			analyzeSlice(m, suffix)
		}
	}
}

func analyzeSlice(value reflect.Value, key string) {
	for i := 0; i < value.Len(); i++ {
		s := value.Index(i)
		kind := s.Kind()
		suffix := analyzePrefix(key, strconv.Itoa(i))
		switch kind {
		case reflect.Ptr:
			if c, ok := s.Interface().(Config); ok {
				analysis(c, suffix)
			}
		case reflect.Map:
			analyzeMap(s, strconv.Itoa(i))
		case reflect.Slice:
			analyzeSlice(s, suffix)
		}
	}
}

func load() {
	for _, def := range GetDefinitions() {
		if init, ok := def.config.(Init); ok {
			init.Init(def.prefix)
		}
	}
}
