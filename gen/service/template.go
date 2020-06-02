package service

var serviceCommon = `// Code generated by gen_batch.go; DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package {{ .Package }}

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// List - api list data
type List struct {
	Total int64       ` + "`json:\"total\"`" + `
	Data  interface{} ` + "`json:\"data\"`" + `
}

// Filter - is the param or list
type Filter struct {
	// @inject_tag: form:"offset"
	Offset int64 ` + "`protobuf:\"varint,1,opt,name=offset,proto3\" json:\"offset,omitempty\" form:\"offset\"`" + `
	// @inject_tag: form:"limit"
	Limit int64 ` + "`protobuf:\"varint,2,opt,name=limit,proto3\" json:\"limit,omitempty\" form:\"limit\"`" + `
	// @inject_tag: form:"order"
	Order string ` + "`protobuf:\"bytes,3,opt,name=order,proto3\" json:\"order,omitempty\" form:\"order\"`" + `
	// @inject_tag: form:"param"
	Param string ` + "`protobuf:\"bytes,4,opt,name=param,proto3\" json:\"param,omitempty\" form:\"param\"`" + `
}

func buildParam(req interface{}, out interface{}) error {
	rVal := reflect.ValueOf(req)
	rType := reflect.TypeOf(req)
	if rType.Kind() == reflect.Ptr {
		rVal = rVal.Elem()
		rType = rType.Elem()
	}

	oVal := reflect.ValueOf(out)
	oType := reflect.TypeOf(out)
	if oType.Kind() == reflect.Ptr {
		oVal = oVal.Elem()
		oType = oType.Elem()
	}

	for index := 0; index < rType.NumField(); index++ {
		reqTypeField := rType.Field(index)
		reqValueField := rVal.Field(index)
		reqJSONTagName := strings.Replace(reqTypeField.Tag.Get("json"), ",omitempty", "", -1)

		// out
		for j := 0; j < oType.NumField(); j++ {
			respTypeField := oType.Field(j)
			respValueField := oVal.Field(j)
			respJSONTagName := strings.Replace(respTypeField.Tag.Get("json"), ",omitempty", "", -1)

			if reqJSONTagName != respJSONTagName {
				continue
			}

			// compare value and set
			switch respValueField.Kind() {
			case reflect.String:
				src := reqValueField.String()
				target := respValueField.String()
				if src != target {
					respValueField.SetString(src)
				}
			case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Int16:
				src := reqValueField.Int()
				target := respValueField.Int()
				if src != target {
					respValueField.SetInt(src)
				}
			case reflect.Float32, reflect.Float64:
				src := reqValueField.Float()
				target := respValueField.Float()
				if src != target {
					respValueField.SetFloat(src)
				}
			case reflect.Bool:
				src := reqValueField.Bool()
				target := respValueField.Bool()
				if src != target {
					respValueField.SetBool(src)
				}
			default:
				fmt.Printf("unknow type")
			}
		}
	}
	return nil
}

func buildParamWithFields(req interface{}, fields []string) map[string]interface{} {
	rVal := reflect.ValueOf(req)
	rType := reflect.TypeOf(req)
	if rType.Kind() == reflect.Ptr {
		rVal = rVal.Elem()
		rType = rType.Elem()
	}

	resp := make(map[string]interface{})

	for index := 0; index < rType.NumField(); index++ {
		reqTypeField := rType.Field(index)
		reqValueField := rVal.Field(index)
		reqJSONTagName := strings.Replace(reqTypeField.Tag.Get("json"), ",omitempty", "", -1)

		// out
		for j := 0; j < len(fields); j++ {
			fieldTagName := fields[j]

			if reqJSONTagName != fieldTagName {
				continue
			}

			// compare value and set
			switch reqValueField.Kind() {
			case reflect.String:
				src := reqValueField.String()
				resp[reqTypeField.Name] = src
			case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Int16:
				src := reqValueField.Int()
				resp[reqTypeField.Name] = src
			case reflect.Float32, reflect.Float64:
				src := reqValueField.Float()
				resp[reqTypeField.Name] = src
			case reflect.Bool:
				src := reqValueField.Bool()
				resp[reqTypeField.Name] = src
			default:
				fmt.Printf("unknow type")
			}
		}
	}
	return resp
}

// checkAppName - check app name format
func checkAppName(value string) error {
	if strings.Contains(value, " ") {
		return fmt.Errorf("invalid param, param can not contain space")
	}

	if len(value) > 32 {
		return fmt.Errorf("invalid param length")
	}

	// 中文、数字、字母、下划线
	pattern := "^[\u4E00-\u9FA5A-Za-z0-9_]+$"
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	if !reg.MatchString(value) {
		return fmt.Errorf("invalid param")
	}

	return nil
}

// checkAppAlias - check app alias format
func checkAppAlias(value string) error {
	if strings.Contains(value, " ") {
		return fmt.Errorf("invalid param, param can not contain space")
	}

	if len(value) > 16 {
		return fmt.Errorf("invalid param length")
	}

	// 数字、小写字母、中线
	pattern := "^[a-z0-9-]+$"
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	if !reg.MatchString(value) {
		return fmt.Errorf("invalid param")
	}

	return nil
}

// checkAppKeyspace - check app keyspace format
func checkAppKeyspace(value string) error {
	if strings.Contains(value, " ") {
		return fmt.Errorf("invalid param, param can not contain space")
	}

	if len(value) > 16 {
		return fmt.Errorf("invalid param length")
	}

	// 数字、小写字母、下划线
	pattern := "^[a-z0-9_]+$"
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	if !reg.MatchString(value) {
		return fmt.Errorf("invalid param")
	}

	return nil
}

// CheckEmail check Email string
func CheckEmail(email string) error {
	//匹配电子邮箱
	pattern := ` + "`\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*`" + `
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(email) {
		return fmt.Errorf("invalid email formats")
	}
	return nil
}

`

var serviceTemplate = `// Code generated by gen_batch.go; DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package {{ .Package }}

import (
	"context"
	"fmt"
	"{{ .ImportModelPath }}"
	{{ if ne .ImportProtoPath "" -}}
	"{{ .ImportProtoPath }}"
	{{ end }}
)

const(
	{{ .DstName }}Success 			= 0
	{{ .DstName }}Failure 			= 1
	{{ .DstName }}InvalidParams	= 2
	ErrorCreate{{ .DstName }} 		= {{ .Index }}1
	ErrorQuery{{ .DstName }} 		= {{ .Index }}2
	ErrorUpdate{{ .DstName }} 		= {{ .Index }}3
	ErrorDelete{{ .DstName }} 		= {{ .Index }}4
	ErrorList{{ .DstName }} 		= {{ .Index }}5
)

// Add{{ .DstName }} - create a new {{ .DstName | ToSnake }}
func Add{{ .DstName}}(ctx context.Context, req *proto.{{ .ReqName }}) (int64, string, interface{}) {
	{{ if eq .CheckApp true }} 
	// check whether app is exist
	if _, err := models.GetApp(req.AppId); err != nil {
		return {{ .DstName }}InvalidParams, fmt.Sprintf("query by appID err, err:%v", err.Error()), nil
	}
	{{ end }}

	{{ .DstConstruct }}
	id, err := models.Add{{ .DstName }}(val)
	if err != nil {
		return ErrorCreate{{ .DstName }}, err.Error(), nil
	}
	return {{ .DstName }}Success, fmt.Sprintf("create {{ .DstName | ToSnake }} success, id:%v", id), id
}

// Get{{ .DstName }} - query {{ .DstName | ToSnake }}
func Get{{ .DstName }}(ctx context.Context, id int64) (int64, string, interface{}) {
	val, err := models.Get{{ .DstName }}(id)
	if err != nil {
		return ErrorQuery{{ .DstName }}, err.Error(), nil
	}
	return {{ .DstName }}Success, "query {{ .DstName | ToSnake }} success", convert{{ .DstName }}(val)
}

// Update{{ .DstName }} - update {{ .DstName | ToSnake }}
func Update{{ .DstName }}(ctx context.Context, req *proto.{{ .ReqName }}, fields []string) (int64, string, interface{}) {
	id := req.Id
	val, err := models.Get{{ .DstName }}(id)
	if err != nil {
		return ErrorQuery{{ .DstName }}, err.Error(), nil
	}

	if err := buildParam(req, val); err != nil {
		return {{ .DstName }}InvalidParams, err.Error(), nil
	}

	var params map[string]interface{}
	if len(fields) > 0 {
		params = buildParamWithFields(val, fields)
	}

	rowAffected, err := models.Edit{{ .DstName }}(val, params)
	if err != nil {
		return ErrorUpdate{{ .DstName }}, err.Error(), nil
	}
	return {{ .DstName }}Success, fmt.Sprintf("update {{ .DstName | ToSnake }} success, affected count:%v", rowAffected), nil
}

// Delete{{ .DstName }} - delete {{ .DstName | ToSnake }}
func Delete{{ .DstName }}(ctx context.Context, id int64) (int64, string, interface{}) {
	rowAffected, err := models.Delete{{ .DstName }}(id)
	if err != nil {
		return ErrorDelete{{ .DstName }}, err.Error(), nil
	}
	return {{ .DstName }}Success, fmt.Sprintf("delete {{ .DstName | ToSnake }} success, affected count:%v", rowAffected), nil
}

// List{{ .DstName }} - query {{ .DstName | ToSnake }} list
func List{{ .DstName }}(ctx context.Context, param *Filter) (int64, string, interface{}) {
	list, err := models.List{{ .DstName }}(int64(param.Offset), int64(param.Limit), param.Order, param.Params())
	if err != nil {
		return ErrorList{{ .DstName }}, err.Error(), nil
	}

	resp := make([]*proto.{{ .DstName }}, 0)
	for i := 0; i < len(list); i++ {
		resp = append(resp, convert{{ .DstName }}(&list[i]))
	}

	total, err := models.Count{{ .DstName }}(param.Params())
	if err != nil {
		return ErrorList{{ .DstName }}, err.Error(), nil
	}
	return {{ .DstName }}Success, "list {{ .DstName | ToSnake }} success", api.List{Total: total, Data: resp}
}
`