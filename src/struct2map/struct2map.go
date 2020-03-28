package struct2map

import (
	"reflect"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Struct2Map 结构转成Map
func Struct2Map(s, m interface{}) {
	var elem reflect.Value
	if reflect.TypeOf(s).Kind() == reflect.Ptr {
		elem = reflect.ValueOf(s).Elem()
	} else {
		elem = reflect.ValueOf(s)
	}

	relType := elem.Type() // 真实类型

	if elem.Kind() != reflect.Struct {
		panic("error just allow struct")
	}

	for i := 0; i < relType.NumField(); i++ {
		var canEmpty, isObjID bool
		tagsName := relType.Field(i).Name
		curElem := elem.Field(i)
		if tags, ok := relType.Field(i).Tag.Lookup("bson"); ok {
			tagsArr := strings.Split(tags, ",")
			for _, v := range tagsArr {
				if v == "omitempty" {
					canEmpty = true
				} else {
					tagsName = v
				}
			}
		}

		if oid, ok := relType.Field(i).Tag.Lookup("oid"); ok {
			if can, err := strconv.ParseBool(oid); can && err == nil {
				isObjID = true
			}
		}

		SetValue(curElem, m, tagsName, canEmpty, false, isObjID)
	}
}

// SetValue 设置值
func SetValue(v reflect.Value, m, k interface{}, canEmpty, isArr, isObjID bool) {
	var value interface{}
	switch v.Kind() {
	case reflect.Struct:
		innerMap := make(map[interface{}]interface{})
		Struct2Map(v.Interface(), &innerMap)
		value = innerMap
	case reflect.Map:
		innerMap := make(map[interface{}]interface{})
		for _, idx := range v.MapKeys() {
			SetValue(v.MapIndex(idx), &innerMap, idx.Interface(), false, false, false)
		}
		value = innerMap
	case reflect.Slice:
		innerSlice := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			SetValue(v.Index(i), &innerSlice, i, false, true, false)
		}
		value = innerSlice
	case reflect.Ptr:
		if v.IsNil() {
			if canEmpty {
				return
			}
			value = nil
		} else {
			SetValue(v.Elem(), m, k, canEmpty, isArr, isObjID)
			return
		}
	default:
		value = v.Interface()
	}

	if isObjID {
		var err error
		value, err = primitive.ObjectIDFromHex(value.(string))
		if err != nil {
			panic(err)
		}
	}

	if isArr {
		(*m.(*[]interface{}))[k.(int)] = value
	} else {
		(*m.(*map[interface{}]interface{}))[k] = value
	}
}
