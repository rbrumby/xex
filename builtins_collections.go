package xex

import (
	"fmt"
	"reflect"
)

func registerCollectionBuiltins() {
	RegisterFunction(
		NewFunction(
			"slice",
			FunctionDocumentation{
				Text: `Makes a new slice containing the passed in values. The type of slice created is determined by the type passed in the first element of values.
				slice can be used to create a list of values to test against - is myproperty x, y or z?: select(slice("x", "y", "z"), .myproperty) > 0`,
				Parameters: map[string]string{
					"values": "variadic - any number of values can be passed to be built into a slice. Types must be compatible with the first value passed.",
				},
			},
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
			FunctionDocumentation{
				Text: `Makes a new map containing the passed in mapEntry values.
				The type of the map (key / value) created is determined by the types passed in the first element of values.`,
				Parameters: map[string]string{
					"values": "variadic - any number of MapEntry's can be passed to be built into a Map. Types must be compatible with the first value passed.",
				},
			},
			func(values ...MapEntry) (out interface{}, err error) {
				logger.Tracef("creating map of type %s:%s", reflect.TypeOf(values[0].Key), reflect.TypeOf(values[0].Value))
				defer func() {
					recv := recover()
					if recv != nil {
						err = fmt.Errorf("error creating map: %s", recv)
					}
				}()
				mtyp := reflect.MapOf(reflect.TypeOf(values[0].Key), reflect.TypeOf(values[0].Value))
				mout := reflect.MakeMap(mtyp)
				for _, e := range values {
					mout.SetMapIndex(reflect.ValueOf(e.Key), reflect.ValueOf(e.Value))
				}
				logger.Tracef("Created %s of %s: %v", mout.Kind(), mout.Type().Elem(), mout.Interface())
				return mout.Interface(), nil
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"entry",
			FunctionDocumentation{
				Text: `Creates a map entry with the passed in key & value.`,
				Parameters: map[string]string{
					"key":   "The map entry key.",
					"value": "The map entry value.",
				},
			},
			func(key interface{}, value interface{}) (entry MapEntry) {
				return MapEntry{key, value}
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"select",
			FunctionDocumentation{
				Text: `Returns the elements in the passed in collection (slice / array or map) for which expression evaluates to true.
				If an array is passed in, it is returned as a slice.
				If coll refers to a map, expression is evaluated on the map value, not the key.
				Example:
				//BookList is a collection. For each "book" in the list, we want to evaluate the equals Expression.
				//We also pass enother evaluated value SelectedAuthor which will be accessible as $0 in our expression.
				select(root.BookList, "book", "equals(book.Author, $0)", root.SelectedAuthor)`,
				Parameters: map[string]string{
					"coll":    "The collection (array, slice or map) to select from.",
					"forEach": "The name by which we will refer to each entry in coll",
					"expr":    "An expression (Node) to apply using to each value in coll. MUST return a bool (true or false).",
					"refs":    "An optional list values () which can be referenced as $0, $1, etc within the expression.",
				},
			},
			func(coll interface{}, forEach string, expr Node, refs ...interface{}) (interface{}, error) {
				values := make(Values)
				//Use the indices of refs to create a map of $n values
				for refIdx, ref := range refs {
					values[fmt.Sprintf("$%d", refIdx)] = ref
				}
				var out reflect.Value
				switch reflect.TypeOf(coll).Kind() {
				case reflect.Array, reflect.Slice:
					logger.Debugf("selecting array/slice: input is %s of %s", reflect.TypeOf(coll).Kind(), reflect.TypeOf(coll).Elem().Name())
					out = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(coll).Elem()), 0, 0)
					for i := 0; i < reflect.ValueOf(coll).Len(); i++ {
						values[forEach] = reflect.ValueOf(coll).Index(i).Interface()
						logger.Debugf("Checking if %v matches expression %v", values[forEach], expr)
						eval, err := expr.Evaluate(values)
						if err != nil {
							return nil, fmt.Errorf("error selecting array/slice: %s", err)
						}
						if e, ok := eval.(bool); ok && e {
							//expression evaluated to true, add the current record to our output slice
							out = reflect.Append(out, reflect.ValueOf(values[forEach]))
						} else if !ok {
							return nil, fmt.Errorf("selector expression must return bool (true/false) not %q", reflect.TypeOf(eval).String())
						}
					}
				case reflect.Map:
					logger.Tracef("selecting map: input is %s of %s", reflect.TypeOf(coll).Kind(), reflect.TypeOf(coll).Elem().Name())
					out = reflect.MakeMap(reflect.MapOf(reflect.TypeOf(coll).Key(), reflect.TypeOf(coll).Elem()))
					for _, k := range reflect.ValueOf(coll).MapKeys() {
						values[forEach] = reflect.ValueOf(coll).MapIndex(k).Interface()
						eval, err := expr.Evaluate(values)
						if err != nil {
							return nil, fmt.Errorf("error selecting from map: %s", err)
						}
						if e, ok := eval.(bool); ok && e {
							out.SetMapIndex(reflect.ValueOf(k.Interface()), reflect.ValueOf(values[forEach]))
						} else if !ok {
							return nil, fmt.Errorf("selector expression must return bool (true/false) not %q", reflect.TypeOf(eval).String())
						}
					}
				default:
					return 0, fmt.Errorf("cannot select from %q", reflect.TypeOf(coll).String())
				}
				logger.Tracef("select: response is a %q of %q", reflect.TypeOf(out.Interface()).Kind(), reflect.TypeOf(out.Interface()).Elem().Name())
				return out.Interface(), nil
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"indexOf",
			FunctionDocumentation{
				Text: `Returns the entry from the passed collection at the requested index.`,
				Parameters: map[string]string{
					"coll":  "The collection (array, slice or map) from which to extract a value.",
					"index": "The index / key to extract from coll",
				},
			},
			func(coll interface{}, index interface{}) (interface{}, error) {
				switch reflect.TypeOf(coll).Kind() {
				case reflect.Array, reflect.Slice:
					idx, ok := index.(int)
					if !ok {
						return nil, fmt.Errorf("index_of: cannot access array / slice using a %s", reflect.TypeOf(coll).String())
					}
					return reflect.ValueOf(coll).Index(idx).Interface(), nil
				case reflect.Map:
					logger.Debugf("indexOf: Accessing map with %s %v", reflect.TypeOf(index), index)
					entry := reflect.ValueOf(coll).MapIndex(reflect.ValueOf(index))
					if !entry.IsValid() {
						return nil, fmt.Errorf("indexOf could not find map entry with key %v of type %s", index, reflect.ValueOf(index))
					}
					return entry.Interface(), nil
				}
				return nil, fmt.Errorf("index_of: expected array/slice, not %s", reflect.TypeOf(coll).Kind())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"count",
			FunctionDocumentation{
				Text: `Returns the number of elements in the passed in slice / array or map.`,
				Parameters: map[string]string{
					"in": "The number of elements in the collection.",
				},
			},
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
	Key   interface{}
	Value interface{}
}

//ValuesNode is a Node which returns the Values object being processed.
type ValuesNode struct{}

func (n ValuesNode) Name() string {
	return "<ValuesNode>"
}

func (n ValuesNode) Evaluate(values Values) (interface{}, error) {
	return values, nil
}

func (n ValuesNode) String() string {
	return n.Name()
}
