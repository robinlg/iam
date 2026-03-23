package reflect

import "reflect"

func GetObjFieldsMap(obj interface{}, fields []string) map[string]interface{} {
	ret := make(map[string]interface{})

	modelReflect := reflect.ValueOf(obj)
	if modelReflect.Kind() == reflect.Ptr {
		modelReflect = modelReflect.Elem()
	}

	modelRefType := modelReflect.Type()
	fieldsCount := modelReflect.NumField()
	var fieldData interface{}
	for i := 0; i < fieldsCount; i++ {
		field := modelReflect.Field(i)
		if len(fields) != 0 && !findString(fields, modelRefType.Field(i).Name) {
			continue
		}

		switch field.Kind() {
		case reflect.Struct:
			fallthrough
		case reflect.Ptr:
			fieldData = GetObjFieldsMap(field.Interface(), []string{})
		default:
			fieldData = field.Interface()
		}

		ret[modelRefType.Field(i).Name] = fieldData
	}

	return ret
}

func CopyObj(from interface{}, to interface{}, fields []string) (changed bool, err error) {
	fromMap := GetObjFieldsMap(from, fields)
	toMap := GetObjFieldsMap(to, fields)
	if reflect.DeepEqual(fromMap, toMap) {
		return false, nil
	}

	t := reflect.ValueOf(to).Elem()
	for k, v := range fromMap {
		val := t.FieldByName(k)
		val.Set(reflect.ValueOf(v))
	}
	return true, nil
}

// findString return true if target in slice, return false if not.
func findString(slice []string, target string) bool {
	for _, str := range slice {
		if str == target {
			return true
		}
	}
	return false
}
