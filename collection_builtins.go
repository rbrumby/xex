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
					err := forEachSlice(
						input,
						func(elem interface{}) error {
							eval, err := exp.Evaluate(elem)
							if err != nil {
								return fmt.Errorf("select: %s", err)
							}
							if e, ok := eval.(bool); ok && e {
								res = reflect.Append(res, reflect.ValueOf(elem))
							} else if !ok {
								return fmt.Errorf("Expression passed to select must return bool not %s", reflect.TypeOf(eval).String())
							}
							return nil
						},
					)
					if err != nil {
						return nil, err
					}
				case reflect.Map:
					logger.Tracef("select map: input is %s of %s", reflect.TypeOf(input).Kind(), reflect.TypeOf(input).Elem().Name())
					res = reflect.MakeMap(reflect.MapOf(reflect.TypeOf(input).Key(), reflect.TypeOf(input).Elem()))
					err := forEachMap(
						input,
						func(key, val interface{}) error {
							eval, err := exp.Evaluate(val)
							if err != nil {
								return fmt.Errorf("select: %s", err)
							}
							if e, ok := eval.(bool); ok && e {
								res.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
							}
							return nil
						},
					)
					if err != nil {
						return nil, err
					}
				default:
					return nil, errors.New("select: can only select an array, slice or map")
				}
				logger.Tracef("select: response is a %s of %s", reflect.TypeOf(res.Interface()).Kind(), reflect.TypeOf(res.Interface()).Elem().Name())
				return res.Interface(), nil
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"count",
			`Returns the number of elements in the passed in slice / array or map which match the provided expression.
			To count all elements pass an expression which always returns true.`,
			func(in interface{}, exp *Expression) (int, error) {
				var res int
				switch reflect.TypeOf(in).Kind() {
				case reflect.Array, reflect.Slice:
					forEachSlice(
						in,
						func(elem interface{}) error {
							eval, err := exp.Evaluate(elem)
							if err != nil {
								return fmt.Errorf("count: %s", err)
							}
							if e, ok := eval.(bool); ok && e {
								res++
							}
							return nil
						},
					)
				case reflect.Map:
					forEachMap(
						in,
						func(key, val interface{}) error {
							eval, err := exp.Evaluate(val)
							if err != nil {
								return fmt.Errorf("count: %s", err)
							}
							if e, ok := eval.(bool); ok && e {
								res++
							}
							return nil
						},
					)
				default:
					return 0, fmt.Errorf("cannot count type %s", reflect.TypeOf(in).String())
				}
				return res, nil
			},
		),
	)

}

//forEachSlice is reused in collection functions so they don't have to worry about looping or repeating reflect calls.
//inSlice can be a Slice or Array of any type.
func forEachSlice(inSlice interface{}, call func(elem interface{}) error) error {
	switch reflect.TypeOf(inSlice).Kind() {
	case reflect.Slice,
		reflect.Array:
		for i := 0; i < reflect.ValueOf(inSlice).Len(); i++ {
			if err := call(reflect.ValueOf(inSlice).Index(i).Interface()); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("forEachSlice received %s not array or slice", reflect.TypeOf(inSlice).Kind())
	}
	return nil
}

//forEachMap is reused in collection functions so they don't have to worry about looping or repeating reflect calls.
//inMap can be a map of nay key / value types.
func forEachMap(inMap interface{}, call func(key, val interface{}) error) error {
	switch reflect.TypeOf(inMap).Kind() {
	case reflect.Map:
		for _, k := range reflect.ValueOf(inMap).MapKeys() {
			v := reflect.ValueOf(inMap).MapIndex(k)
			if err := call(k.Interface(), v.Interface()); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("forEachMap received %s not map", reflect.TypeOf(inMap).Kind())
	}
	return nil
}
