package reflection

import "reflect"

// any == interface{} and I felt like using any
// the fn here is the append() function we're using in
// the test to collect all of the values for comparison
func walk(x any, fn func(string)) {
	val := getValue(x)

	// Alternate solution:
	// numberOfValues := 0
	// var getField func(int) reflect.Value

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		// If it's just a string we can grab that value straight
		fn(val.String())
	case reflect.Struct:
		// If it's a struct we just need to walk through the fields' values
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
		// Alternate solution:
		// numberOfValues = val.NumField()
		// getField = val.Field
	case reflect.Slice, reflect.Array:
		// If it's a slice or an array walk through the values by index
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
		// Alternate solution:
		// numberOfValues = val.Len()
		// getField = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		// Here we're calling the function we want to test, which will store results for us to walk
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}

	// Alternate solution:
	// for i := 0; i < numberOfValues; i++ {
	// 	walk(getField(i).Interface(), fn)
	// }
}

// Gets the value of the interface provided and returns either
// the value it points to or the straight value of x
func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		// Elem is good for getting the value of a pointer or interface
		// In this case, we're retreiving the value of the pointer
		// so we can call Field() on it later without error
		val = val.Elem()
	}

	return val
}
