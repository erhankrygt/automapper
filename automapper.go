package automapper

import (
	"fmt"
	"reflect"
)

// Mapper returns destination. Generated from provided source. Will map destination properties with source properties, include values
// Entity to Entity
func Mapper(s interface{}, d interface{}, diff map[string]string) {
	source := registry(s)

	dest := reflect.Indirect(reflect.ValueOf(d))
	if dest.Kind() != reflect.Struct {
		panic("could not convert to struct")
	}

	elem := reflect.ValueOf(d).Elem()
	for _, el := range source {
		for i := 0; i < dest.NumField(); i++ {
			k := dest.Type().Field(i).Name

			if el.Key == k {
				f := elem.FieldByName(k)
				setValue(&f, el)
				continue
			}

			diffKey := diff[el.Key]
			if diffKey != "" {
				f := elem.FieldByName(diffKey)
				setValue(&f, el)
				continue
			}
		}
	}
}

func setValue(f *reflect.Value, obj mapperObject) {
	if f.CanSet() {
		kind := f.Kind()
		switch kind {
		case reflect.String:
			f.SetString(fmt.Sprintf("%v", obj.Value))
			break
		case reflect.Int:
			f.SetInt(int64(obj.Value.(int)))
			break
		case reflect.Float32:
			f.SetFloat(float64(obj.Value.(float32)))
			break
		case reflect.Float64:
			f.SetFloat(obj.Value.(float64))
			break
		case reflect.Bool:
			f.SetBool(obj.Value.(bool))
			break
		}
	}
}

func registry(s interface{}) []mapperObject {
	var obj []mapperObject
	st := reflect.Indirect(reflect.ValueOf(s))
	if st.Kind() != reflect.Struct {
		panic("could not convert to struct")
	}

	for i := 0; i < st.NumField(); i++ {
		obj = append(obj, mapperObject{
			Value: st.Field(i).Interface(),
			Key:   st.Type().Field(i).Name,
		})
	}

	return obj
}

type mapperObject struct {
	Key   string
	Value interface{}
}
