package describer

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var exclusionTypeSet = map[string]struct{}{
	"github.com/gofrs/uuid":                      {},
	"github.com/Azure/go-autorest/autorest/date": {},
}

func isGoPackage(path string) bool {
	return path != "" && !strings.Contains(path, "/")
}

// JSONAllFieldsMarshaller is a hack around the issue described here
// https://githubmemory.com/repo/Azure/azure-sdk-for-go/issues/12227
// Azure sdk overrides all the MarshalJSON methods for the struct fields
// to exclude the 'READ-ONLY' fields from the JSON output of the struct.
// By simply wrapping the original struct by JSONAllFieldsMarshaller, all
// the fields will appear in the json output.
type JSONAllFieldsMarshaller struct {
	Value interface{}
}

func (x JSONAllFieldsMarshaller) MarshalJSON() (res []byte, err error) {
	var val = x.Value

	v := reflect.ValueOf(x.Value)
	if !v.IsValid() {
		return json.Marshal(val)
	}
	if _, ok := exclusionTypeSet[v.Type().PkgPath()]; !ok && !isGoPackage(v.Type().PkgPath()) {
		switch v.Kind() {
		case reflect.Slice, reflect.Array:
			val = azSliceMarshaller{Value: v}
		case reflect.Ptr:
			val = azPtrMarshaller{Value: v}
		case reflect.Struct:
			val = azStructMarshaller{Value: v}
		}
	}

	return json.Marshal(val)
}

func (x *JSONAllFieldsMarshaller) UnmarshalJSON(data []byte) (err error) {
	v := reflect.ValueOf(x.Value)
	if !v.IsValid() {
		return nil
	}
	if _, ok := exclusionTypeSet[v.Type().PkgPath()]; !ok && !isGoPackage(v.Type().PkgPath()) {
		switch v.Kind() {
		case reflect.Slice, reflect.Array:
			val := &azSliceMarshaller{Value: v}
			err := json.Unmarshal(data, val)
			if err != nil {
				return err
			}
			newVal := reflect.New(v.Type())
			if !val.Value.Type().AssignableTo(newVal.Elem().Type()) {
				return nil
			}
			newVal.Elem().Set(val.Value)
			x.Value = newVal.Elem().Interface()
		case reflect.Struct:
			val := &azStructMarshaller{Value: v}
			err := json.Unmarshal(data, val)
			if err != nil {
				return err
			}
			newVal := reflect.New(v.Type())
			if !val.Value.Type().AssignableTo(newVal.Elem().Type()) {
				return nil
			}
			newVal.Elem().Set(val.Value)
			x.Value = newVal.Elem().Interface()
		case reflect.Ptr:
			val := &azPtrMarshaller{Value: v}
			err := json.Unmarshal(data, val)
			if err != nil {
				return err
			}

			newVal := reflect.New(v.Type())
			if !val.Value.Type().AssignableTo(newVal.Elem().Type()) {
				return nil
			}
			newVal.Elem().Set(val.Value)
			x.Value = newVal.Elem().Interface()
		default:
			val := reflect.New(v.Type())
			err := json.Unmarshal(data, val.Interface())
			if err != nil {
				return err
			}

			newVal := reflect.New(v.Type())
			if !val.Elem().Type().AssignableTo(newVal.Elem().Type()) {
				return nil
			}
			newVal.Elem().Set(val.Elem())
			x.Value = newVal.Elem().Interface()
		}
		return nil
	}

	val := reflect.New(v.Type())
	err = json.Unmarshal(data, val.Interface())
	if err != nil {
		return err
	}

	newVal := reflect.New(v.Type())
	if !val.Elem().Type().AssignableTo(newVal.Elem().Type()) {
		return nil
	}
	newVal.Elem().Set(val.Elem())
	x.Value = newVal.Elem().Interface()

	return nil
}

type azStructMarshaller struct {
	reflect.Value
}

func (x azStructMarshaller) MarshalJSON() ([]byte, error) {
	v := x.Value
	m := make(map[string]interface{})
	num := v.Type().NumField()
	for i := 0; i < num; i++ {
		field := v.Type().Field(i)
		if !field.IsExported() {
			continue
		}

		jsonTag := field.Tag.Get("json")
		jsonFields := strings.Split(jsonTag, ",")
		jsonField := jsonFields[0]
		if jsonField == "-" {
			continue
		}
		jsonField = field.Name

		jsonOmitEmpty := false
		for _, field := range jsonFields {
			if field == "omitempty" {
				jsonOmitEmpty = true
				break
			}
		}
		if jsonOmitEmpty && isEmptyValue(v.Field(i)) {
			continue
		}
		m[jsonField] = JSONAllFieldsMarshaller{Value: v.Field(i).Interface()}
	}

	return json.Marshal(m)
}

func (x *azStructMarshaller) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", x.Value, err)
	}

	x.Value = reflect.New(x.Value.Type()).Elem()
	for i := 0; i < x.Value.Type().NumField(); i++ {
		field := x.Value.Type().Field(i)
		if !field.IsExported() {
			continue
		}

		jsonTag := field.Tag.Get("json")
		jsonFields := strings.Split(jsonTag, ",")
		jsonField := jsonFields[0]
		if jsonField == "-" {
			continue
		}
		jsonField = field.Name
		if msg, ok := rawMsg[jsonField]; !ok || string(msg) == "null" {
			continue
		}
		var err error
		k := reflect.New(field.Type)
		if k.Elem().Kind() == reflect.Interface {
			continue
		}
		v := JSONAllFieldsMarshaller{Value: reflect.New(field.Type).Elem().Interface()}
		err = json.Unmarshal(rawMsg[jsonField], &v)
		if err != nil {
			return fmt.Errorf("unmarshalling field %s: %v", jsonField, err)
		}

		// check if the type is assignable
		if !reflect.ValueOf(v.Value).Type().AssignableTo(k.Elem().Type()) {
			continue
		}
		k.Elem().Set(reflect.ValueOf(v.Value))
		if !k.Elem().Type().AssignableTo(x.Value.Type().Field(i).Type) {
			continue
		}
		x.Value.Field(i).Set(k.Elem())
	}

	return nil
}

type azPtrMarshaller struct {
	reflect.Value
}

func (x azPtrMarshaller) MarshalJSON() ([]byte, error) {
	val := x.Value
	for val.Type().Kind() == reflect.Ptr {
		if val.IsNil() {
			return []byte("null"), nil
		}
		val = val.Elem()
	}

	if !val.CanInterface() {
		return nil, errors.New("cannot interface ptr marshaller")
	}

	return JSONAllFieldsMarshaller{Value: val.Interface()}.MarshalJSON()
}

func (x *azPtrMarshaller) UnmarshalJSON(data []byte) error {
	v := reflect.New(x.Value.Type().Elem())
	k := JSONAllFieldsMarshaller{Value: v.Elem().Interface()}
	if err := json.Unmarshal(data, &k); err != nil {
		return err
	}

	p := reflect.New(reflect.TypeOf(k.Value))
	if !reflect.ValueOf(k.Value).Type().AssignableTo(p.Elem().Type()) {
		return nil
	}
	p.Elem().Set(reflect.ValueOf(k.Value))

	x.Value = p

	return nil
}

type azSliceMarshaller struct {
	reflect.Value
}

func (x azSliceMarshaller) MarshalJSON() ([]byte, error) {
	num := x.Value.Len()
	if num == 0 {
		return []byte("null"), nil
	}
	list := make([]JSONAllFieldsMarshaller, 0, num)
	for i := 0; i < num; i++ {
		if !x.Value.Index(i).CanInterface() {
			continue
		}

		list = append(list, JSONAllFieldsMarshaller{Value: x.Value.Index(i).Interface()})
	}

	return json.Marshal(list)
}

func (x *azSliceMarshaller) UnmarshalJSON(data []byte) error {
	var list []json.RawMessage
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}

	num := len(list)
	x.Value = reflect.MakeSlice(x.Value.Type(), num, num)
	for i := 0; i < num; i++ {
		v := reflect.New(x.Value.Type().Elem())
		if v.Elem().Kind() == reflect.Interface {
			continue
		}
		k := JSONAllFieldsMarshaller{Value: v.Interface()}
		if err := json.Unmarshal(list[i], &k); err != nil {
			return err
		}
		if !reflect.ValueOf(k.Value).Elem().Type().AssignableTo(x.Value.Type().Elem()) {
			continue
		}
		x.Value.Index(i).Set(reflect.ValueOf(k.Value).Elem())
	}

	return nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
