package object

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

type DictPair struct {
	Key   Object
	Value Object
}

type Dict struct {
	Pairs  map[HashKey]DictPair
	offset int
}

func (d *Dict) Type() ObjectType { return DICT_OBJ }
func (d *Dict) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}

	for _, pair := range d.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

func (d *Dict) Next() (Object, Object) {
	idx := 0
	dict := make(map[string]DictPair)
	var keys []string
	for _, v := range d.Pairs {
		dict[v.Key.Inspect()] = v
		keys = append(keys, v.Key.Inspect())
	}

	sort.Strings(keys)

	for _, k := range keys {
		if d.offset == idx {
			d.offset += 1
			return dict[k].Key, dict[k].Value
		}
		idx += 1
	}
	return nil, nil
}

func (d *Dict) Reset() {
	d.offset = 0
}
