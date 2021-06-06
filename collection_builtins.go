package xex

import (
	"fmt"
	"reflect"
)

func registerCollectionBuiltins() {
	RegisterFunction(
		NewFunction(
			"slice",
			`Makes a new slice containing the passed in values. The type of slice created is determined by the type passed in the first element of values.
			slice can be used to create a list of values to test against - is myproperty x, y or z?: select(slice("x", "y", "z"), .myproperty) > 0`,
			func(values ...interface{}) (out interface{}, err error) {
				logger.Tracef("creating slice of type %s", reflect.TypeOf(values[0]))
				defer func() {
					recv := recover()
					if recv != nil {
						err = fmt.Errorf("error creating slice: %s", recv)
					}
				}()
				slice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(values[0])), 0, 0)
				for _, v := range values {
					slice = reflect.Append(slice, reflect.ValueOf(v))
				}
				logger.Tracef("Created %s of %s: %v", slice.Kind(), slice.Type().Elem(), slice.Interface())
				return slice.Interface(), nil
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"map",
			`Makes a new map containing the passed in mapEntry values.
			The type of the map (key / value) created is determined by the types passed in the first element of values.`,
			func(values ...MapEntry) (out interface{}, err error) {
				logger.Tracef("creating map of type %s:%s", reflect.TypeOf(values[0].key), reflect.TypeOf(values[0].value))
				defer func() {
					recv := recover()
					if recv != nil {
						err = fmt.Errorf("error creating map: %s", recv)
					}
				}()
				mtyp := reflect.MapOf(reflect.TypeOf(values[0].key), reflect.TypeOf(values[0].value))
				mout := reflect.MakeMap(mtyp)
				for _, e := range values {
					mout.SetMapIndex(reflect.ValueOf(e.key), reflect.ValueOf(e.value))
				}
				logger.Tracef("Created %s of %s: %v", mout.Kind(), mout.Type().Elem(), mout.Interface())
				return mout.Interface(), nil
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"select",
			`Returns the elements in the passed in slice / array or map for which expression evaluates to true.
			 If an array is passed in, it is returned as a slice.
			 If coll refers to a map, expression is evaluated on the map value, not the key. 
			 - values should akways be a Noop which will return the Values being evaluated
			 - coll is a *Expression which evaluates to the collection (array, slice or map) being selected
			 - elemKey is a string which will be used in the expression to refer to the elements of elem
			 - expression is the expression to be evaluated on each node`,
			func(values Values, coll *Expression, elemKey string, expression *Expression) (interface{}, error) {
				elem, err := coll.Evaluate(values)
				if err != nil {
					return nil, fmt.Errorf("selection evaluation error: %s", err)
				}

				//We'll take a copy of values because we will add the recordKey entry for the current record (and don't want to overwrite)
				//anything in the original Values graph.
				valcp := make(Values)
				for vk, vv := range values {
					valcp[vk] = vv
				}

				//out is the output result slice
				var out reflect.Value
				switch reflect.TypeOf(elem).Kind() {
				case reflect.Array, reflect.Slice:
					logger.Tracef("selecting array/slice: input is %s of %s", reflect.TypeOf(elem).Kind(), reflect.TypeOf(elem).Elem().Name())
					out = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(elem).Elem()), 0, 0)
					for i := 0; i < reflect.ValueOf(elem).Len(); i++ {
						//Add the current record from the array/slice to valcp & pass it down to the expression for evaluation
						valcp[elemKey] = reflect.ValueOf(elem).Index(i).Interface()
						eval, err := expression.Evaluate(valcp)
						if err != nil {
							return nil, fmt.Errorf("error selecting array/slice: %s", err)
						}
						if e, ok := eval.(bool); ok && e {
							//expression evaluated to true, add the current record to our output slice
							out = reflect.Append(out, reflect.ValueOf(valcp[elemKey]))
						} else if !ok {
							return nil, fmt.Errorf("selector expression must return bool (true/false) not %q", reflect.TypeOf(eval).String())
						}
					}
				case reflect.Map:
					logger.Tracef("select map: input is %s of %s", reflect.TypeOf(elem).Kind(), reflect.TypeOf(elem).Elem().Name())
					out = reflect.MakeMap(reflect.MapOf(reflect.TypeOf(elem).Key(), reflect.TypeOf(elem).Elem()))
					for _, k := range reflect.ValueOf(elem).MapKeys() {
						//Add the current value from the map to valcp & pass it down to the expression for evaluation
						valcp[elemKey] = reflect.ValueOf(elem).MapIndex(k).Interface()
						eval, err := expression.Evaluate(valcp)
						if err != nil {
							return nil, fmt.Errorf("error selecting from map: %s", err)
						}
						if e, ok := eval.(bool); ok && e {
							out.SetMapIndex(reflect.ValueOf(k.Interface()), reflect.ValueOf(valcp[elemKey]))
						} else if !ok {
							return nil, fmt.Errorf("selector expression must return bool (true/false) not %q", reflect.TypeOf(eval).String())
						}
					}
				default:
					return 0, fmt.Errorf("cannot select from %q", reflect.TypeOf(elem).String())
				}
				logger.Tracef("select: response is a %q of %q", reflect.TypeOf(out.Interface()).Kind(), reflect.TypeOf(out.Interface()).Elem().Name())
				return out.Interface(), nil
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"count",
			`Returns the number of elements in the passed in slice / array or map.`,
			func(in interface{}) (int, error) {
				switch reflect.TypeOf(in).Kind() {
				case reflect.Array, reflect.Slice, reflect.Map:
					return reflect.ValueOf(in).Len(), nil
				}
				return 0, fmt.Errorf("cannot count type %s", reflect.TypeOf(in).String())
			},
		),
	)
}

type MapEntry struct {
	key   interface{}
	value interface{}
}
