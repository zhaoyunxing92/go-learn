package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"strings"
)

func main() {

	path := "/Users/docker/code/github/go/go-learn/scan/service"

	fset := token.NewFileSet()
	packages, _ := parser.ParseDir(fset, path, func(fi os.FileInfo) bool {
		return true
	}, parser.ParseComments)

	for _, pk := range packages {
		//全部文件
		for _, f := range pk.Files {
			//parser.(idx)
			ast.Inspect(f, collectStructs)
			//_, _ = parser.ParseExpr(idx)
			//fmt.Println(idx)
			// 获取注释
			for _, comment := range f.Comments {
				fmt.Println(comment.Text())
			}

			// 获取结构体
			for _, obj := range f.Scope.Objects {
				//of := reflect.TypeOf(obj)
				//elem := of.Elem()
				//fmt.Println(elem)
				//vl = reflect.New(obj)
				//fmt.Println(vl)
				fmt.Println(obj.Name)
				fmt.Println(obj.Type)

			}
		}
	}
}

func collectStructs(x ast.Node) bool {
	ts, ok := x.(*ast.TypeSpec)
	if !ok || ts.Type == nil {
		return true
	}

	// 获取结构体名称
	structName := ts.Name.Name

	fmt.Println(structName)

	s, ok := ts.Type.(*ast.StructType)
	if !ok {
		return true
	}

	for _, field := range s.Fields.List {
		if field.Tag == nil {
			continue
		}
		tag := field.Tag.Value
		tag = strings.Trim(tag, "`")
		//tags, err := structtag.Parse(string(tag))
		//if err != nil {
		//	return true
		//}
		//_, err = tags.Get(tagName)
		//if err == nil {
		//	structMap[structName] = append(structMap[structName], field)
		//}
	}
	return false
}

//getFieldsOfStructAddr 获取结构体实例字段的地址
func getFieldsOfStructAddr(s interface{}, filedNames []string) (res []interface{}, err error) {
	v := reflect.ValueOf(s)
	//如果传入指针则拿到指针指向的值
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	//参数必须是结构体
	for v.Kind() != reflect.Struct {
		//防止panic，返回自定义错误
		return
	}
	//v必须是一个实例
	if !v.CanAddr() {
		//防止panic，返回自定义错误
		return
	}
	for _, filedName := range filedNames {
		fieldValue := v.FieldByName(filedName)
		if !fieldValue.CanAddr() {
			//防止panic，返回自定义错误
			return
		}
		res = append(res, fieldValue.Addr().Interface())
	}

	return
}

//getTagAndFiledNameOfStruct 获取结构体字段的变量名和tag名
func getTagAndFiledNameOfStruct(s interface{}, tag string) (tagsRes []string, filedNames []string, err error) {
	t := reflect.TypeOf(s)
	//如果传入指针则拿到指针指向的值
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	//传入的必须是结构体
	if t.Kind() != reflect.Struct {
		return
	}
	//获取结构体字段数量
	n := t.NumField()
	for i := 0; i < n; i++ {
		field := t.Field(i)
		tagsRes = append(tagsRes, field.Tag.Get(tag))
		filedNames = append(filedNames, field.Name)
	}
	return
}
