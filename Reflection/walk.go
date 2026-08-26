package reflection

import "reflect"

func walk(x interface{}, function func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		iterateArraySlice(&val, function)
	case reflect.Map:
		iterateMap(&val, function)
	case reflect.Struct:
		iterateStruct(&val, function)
	case reflect.String:
		function(val.String())
	case reflect.Chan:
		iterateChannel(&val, function)
	case reflect.Func:
		iterateFunc(&val, function)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}

func iterateArraySlice(val *reflect.Value, function func(input string)) {
	for i := 0; i < val.Len(); i++ {
		walk(val.Index(i).Interface(), function)
	}
}

func iterateStruct(val *reflect.Value, function func(input string)) {
	for _, field := range val.Fields() {
		walk(field.Interface(), function)
	}
}

func iterateMap(val *reflect.Value, function func(input string)) {
	iter := val.MapRange()
	for iter.Next() {
		walk(iter.Value().Interface(), function)
	}
}

func iterateChannel(val *reflect.Value, function func(input string)) {
	for {
		if v, ok := val.Recv(); ok {
			walk(v.Interface(), function)
		} else {
			break
		}
	}
}

func iterateFunc(val *reflect.Value, function func(input string)) {
	funcVals := val.Call(nil)
	for _, element := range funcVals {
		walk(element.Interface(), function)
	}
}
