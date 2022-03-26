package automapper

import (
	"reflect"
)

// Mapper returns destination. Generated from provided source. Will map destination properties with source properties, include values
// Entity to Entity
func Mapper(source interface{}, destination interface{}) {
	s := reflect.ValueOf(source)
	if s.Kind() != reflect.Struct {
		panic("source' type is not struct")
	}

	d := reflect.Indirect(reflect.ValueOf(destination))
	if d.Kind() != reflect.Struct {
		panic("destination' type is not struct")
	}

	v := reflect.Indirect(s)
	vd := reflect.Indirect(d)
	for j := 0; j < v.NumField(); j++ {
		f := vd.FieldByName(v.Type().Field(j).Name)
		f.Set(v.Field(j))
	}
}

// MapperForList returns destination collection. Generated from provided source. Will map destination properties with source properties, include values
// List to List
func MapperForList(source interface{}, destination interface{}) interface{} {
	s := reflect.ValueOf(source)
	if s.Kind() != reflect.Slice {
		panic("source' type is not slice")
	}

	d := reflect.Indirect(reflect.ValueOf(destination))
	if d.Kind() != reflect.Slice {
		panic("destination' type is not slice")
	}

	for i := 0; i < s.Len(); i++ {
		// append empty item in destination
		ptr := reflect.New(d.Type().Elem()).Interface()
		pv := reflect.ValueOf(ptr).Elem()
		d = reflect.Append(d, pv)
		// ***

		item := s.Index(i)
		destItem := d.Index(i)

		if item.Kind() == reflect.Struct {
			v := reflect.Indirect(item)
			vd := reflect.Indirect(destItem)
			for j := 0; j < v.NumField(); j++ {
				f := vd.FieldByName(v.Type().Field(j).Name)
				f.Set(v.Field(j))
			}
		}
	}

	return d.Interface()
}
