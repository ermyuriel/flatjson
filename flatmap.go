package flatmap

import (
	"fmt"
	"reflect"
)

//FlatternMap will flatten a map, nested maps and array values will become main level fields
func Flatten(m map[string]interface{}, keep bool) {
	flatmap("", m, m, keep)
}

func flatmap(p string, m, wm map[string]interface{}, keep bool) {

	for k, v := range m {
		if v == nil {
			continue
		}
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			if !keep {
				delete(wm, k)
			}

			nk := fmt.Sprintf("%v%v_", p, k)
			flatmap(nk, v.(map[string]interface{}), wm, keep)

		case reflect.Slice:
			if !keep {
				delete(wm, k)
			}

			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {

				iv := reflect.ValueOf(s.Index(i)).Interface()
				nki := fmt.Sprintf("%v%v_%v", p, k, i)

				if reflect.TypeOf(s.Index(i)).Kind() == reflect.Map {
					flatmap(nki, iv.(map[string]interface{}), wm, keep)
				} else {
					wm[nki] = fmt.Sprintf("%v", iv)
				}

			}
		default:
			nk := fmt.Sprintf("%v%v", p, k)
			wm[nk] = v
		}

	}
}
