package detector

import (
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-azure/global/maps"
	"go/ast"
	"reflect"
	"strings"
)

func getInitValue(t reflect.Type) reflect.Value {
	switch t.Kind() {
	case reflect.String:
		return reflect.ValueOf("string zero value")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(1)
	case reflect.Bool:
		return reflect.ValueOf(true)
	case reflect.Float64, reflect.Float32:
		return reflect.ValueOf(1.0)
	case reflect.Slice:
		slice := reflect.MakeSlice(t, 1, 1)
		switch t.Elem().Kind() {
		case reflect.String:
			slice.Index(0).SetString(getInitValue(t.Elem()).String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice.Index(0).SetInt(getInitValue(t.Elem()).Int())
		case reflect.Bool:
			slice.Index(0).SetBool(getInitValue(t.Elem()).Bool())
		case reflect.Float64, reflect.Float32:
			slice.Index(0).SetFloat(getInitValue(t.Elem()).Float())
		case reflect.Pointer:
			slice.Index(0).Set(reflect.New(t.Elem().Elem()))
			switch t.Elem().Elem().Kind() {
			case reflect.String:
				slice.Index(0).Elem().SetString(getInitValue(t.Elem().Elem()).String())
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				slice.Index(0).Elem().SetInt(getInitValue(t.Elem().Elem()).Int())
			case reflect.Bool:
				slice.Index(0).Elem().SetBool(getInitValue(t.Elem().Elem()).Bool())
			case reflect.Float64, reflect.Float32:
				slice.Index(0).Elem().SetFloat(getInitValue(t.Elem().Elem()).Float())
			default:
				slice.Index(0).Elem().Set(getInitValue(t.Elem().Elem()))
			}
		default:
			slice.Index(0).Set(getInitValue(t.Elem()))
		}
	case reflect.Struct:
		return reflect.New(t).Elem()
	case reflect.Ptr:
		derefVal := getInitValue(t.Elem())
		val := reflect.New(derefVal.Type())
		switch derefVal.Kind() {
		case reflect.String:
			val.Elem().SetString(derefVal.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val.Elem().SetInt(derefVal.Int())
		case reflect.Bool:
			val.Elem().SetBool(derefVal.Bool())
		case reflect.Float64, reflect.Float32:
			val.Elem().SetFloat(derefVal.Float())
		default:
			val.Elem().Set(derefVal)
		}
		return val
	}
	return reflect.New(t).Elem()
}

func initValueInPath(value reflect.Value, path []string) reflect.Value {
	if len(path) == 0 {
		return getInitValue(value.Type())
	}
	if value.Kind() == reflect.Struct {
		field := value.FieldByName(path[0])
		if field.IsValid() {
			if field.Kind() == reflect.Struct {
				field.Set(initValueInPath(field, path[1:]))
			} else if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Struct {
				// init the pointer to a struct
				field.Set(reflect.New(field.Type().Elem()))
				field.Elem().Set(initValueInPath(field.Elem(), path[1:]))
			} else if field.Kind() == reflect.Ptr {
				field.Set(reflect.New(field.Type().Elem()))
				switch field.Type().Elem().Kind() {
				case reflect.String:
					field.Elem().SetString(getInitValue(field.Type().Elem()).String())
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					field.Elem().SetInt(getInitValue(field.Type().Elem()).Int())
				case reflect.Bool:
					field.Elem().SetBool(getInitValue(field.Type().Elem()).Bool())
				case reflect.Float64, reflect.Float32:
					field.Elem().SetFloat(getInitValue(field.Type().Elem()).Float())
				default:
					field.Elem().Set(getInitValue(field.Type().Elem()))
				}
			} else {
				switch field.Kind() {
				case reflect.String:
					field.SetString(getInitValue(field.Type()).String())
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					field.SetInt(getInitValue(field.Type()).Int())
				case reflect.Bool:
					field.SetBool(getInitValue(field.Type()).Bool())
				case reflect.Float64, reflect.Float32:
					field.SetFloat(getInitValue(field.Type()).Float())
				default:
					field.Set(getInitValue(field.Type()))
				}
			}
			value.FieldByName(path[0]).Set(field)
		}
		return value
	} else {
		return getInitValue(value.Type())
	}
}

func Flatten(m map[string]any) map[string]any {
	o := make(map[string]any)
	for k, v := range m {
		switch child := v.(type) {
		case map[string]any:
			nm := Flatten(child)
			for nk, nv := range nm {
				o[k+"."+nk] = nv
			}
		default:
			o[k] = v
		}
	}
	return o
}

func DetectFilters(resourceType string, tableNode *ast.File) (getFilters map[string]string, listFilters map[string]string) {
	defer func() {

	}()
	getFilters = make(map[string]string)
	listFilters = make(map[string]string)

	descModel, ok := maps.ResourceTypeToDescription[resourceType]
	if !ok {
		return getFilters, listFilters
	}

	ast.Inspect(tableNode, func(tnode ast.Node) bool {
		if c, ok := tnode.(*ast.CompositeLit); ok {
			var columnName, transformer string
			for _, arg := range c.Elts {
				if kv, ok := arg.(*ast.KeyValueExpr); ok {
					if i, ok := kv.Key.(*ast.Ident); ok {
						if i.Name == "Name" {
							if bl, ok := kv.Value.(*ast.BasicLit); ok {
								columnName = strings.Trim(bl.Value, "\"")
							}
						} else if i.Name == "Transform" {
							if cl, ok := kv.Value.(*ast.CallExpr); ok {
								transformer = ExtractTransformer(cl)
							}
						}
					}
				}
			}

			// We only want to detect filters for the description columns
			if !strings.HasPrefix(transformer, "Description.") {
				return true
			}
			// make a null instance of the description model
			descModelReflect := reflect.New(reflect.TypeOf(descModel)).Elem()
			valuePath := strings.Split(transformer, ".")
			// fill the description model with a non zero value to detect json path later
			descModelReflect = initValueInPath(descModelReflect, valuePath)
			jsonModel, err := json.Marshal(descModelReflect.Interface())
			if err != nil {
				fmt.Println("Error marshalling description model")
				return true
			}
			intermediateMap := make(map[string]any)
			err = json.Unmarshal(jsonModel, &intermediateMap)
			if err != nil {
				fmt.Println("Error unmarshalling description model")
				return true
			}
			flattenedMap := Flatten(intermediateMap)
			for k, v := range flattenedMap {
				if v == nil || reflect.ValueOf(v).IsZero() {
					continue
				}
				valV := reflect.ValueOf(v)
				initV := getInitValue(valV.Type())
				if reflect.DeepEqual(valV.Interface(), initV.Interface()) {
					getFilters[columnName] = k
					listFilters[columnName] = k
				}
			}
			return true
		}
		return true
	})

	return getFilters, listFilters
}

func ExtractTransformer(cl *ast.CallExpr) string {
	if sl, ok := cl.Fun.(*ast.SelectorExpr); ok {
		if sl.Sel.Name == "Transform" {
			return ""
		}
		if call, ok := sl.X.(*ast.CallExpr); ok {
			return ExtractTransformer(call)
		}
		if sl.Sel.Name == "FromField" {
			for _, arg := range cl.Args {
				if bl, ok := arg.(*ast.BasicLit); ok {
					return strings.Trim(bl.Value, "\"")
				}
			}
		}
	}
	return ""
}
