package xex

import (
	"errors"
	"fmt"
	"reflect"
)

func registerCollectionBuiltins() {
	RegisterFunction(
		NewFunction(
			"select",
			`input must be an array, slice or map from which each element is evaluated against exp
			(in the case of a map, the value is evaluated not the key).
			exp must return a bool. If exp results in true for an element, that element is added to the returned 
			slice or map containing the same type as the input. Note that if an array is passed as input, a slice is returned.
			`,
			func(input interface{}, exp *Expression) (interface{}, error) {
				var res reflect.Value
				switch reflect.TypeOf(input).Kind() {
				case reflect.Array,
					reflect.Slice:
					logger.Tracef("select array/slice: input is %s of %s", reflect.TypeOf(input).Kind(), reflect.TypeOf(input).Elem().Name())
					res = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(input).Elem()), 0, 0)
					for i := 0; i < reflect.ValueOf(input).Len(); i++ {
						eval, err := exp.Evaluate(reflect.ValueOf(input).Index(i).Interface())
						if err != nil {
							return nil, fmt.Errorf("select: %s", err)
						}
						if e, ok := eval.(bool); ok && e {
							res = reflect.Append(res, reflect.ValueOf(input).Index(i))
						} else if !ok {
							return nil, fmt.Errorf("Expression passed to select must return bool not %s", reflect.TypeOf(eval).String())
						}
					}
				case reflect.Map:
					logger.Tracef("select map: input is %s of %s", reflect.TypeOf(input).Kind(), reflect.TypeOf(input).Elem().Name())
					res = reflect.MakeMap(reflect.MapOf(reflect.TypeOf(input).Key(), reflect.TypeOf(input).Elem()))
					// if inmap, ok := input.(map[interface{}]interface{}); !ok {
					// 	return nil, errors.New("select: could not convert input to map")
					// } else {
					for _, k := range reflect.ValueOf(input).MapKeys() {
						v := reflect.ValueOf(input).MapIndex(k)
						eval, err := exp.Evaluate(v.Interface())
						if err != nil {
							return nil, fmt.Errorf("select: %s", err)
						}
						if e, ok := eval.(bool); ok && e {
							res.SetMapIndex(reflect.ValueOf(k.Interface()), reflect.ValueOf(v.Interface()))
						}
					}
				default:
					return nil, errors.New("select: can only select an array, slice or map")
				}
				logger.Tracef("select: response is a %s of %s", reflect.TypeOf(res.Interface()).Kind(), reflect.TypeOf(res.Interface()).Elem().Name())
				return res.Interface(), nil
			},
		),
	)

}
