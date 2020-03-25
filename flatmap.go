package flatmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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

				iv := reflect.ValueOf(s.Index(i))
				nki := fmt.Sprintf("%v%v_%v", p, k, i)

				if reflect.TypeOf(iv.Interface()).Kind() == reflect.Struct {

					js, err := json.Marshal(s.Index(i).Interface())

					if err == nil {
						m := make(map[string]interface{})

						json.NewDecoder(bytes.NewBuffer(js)).Decode(&m)
						flatmap(nki+"_", m, wm, keep)

					} else {
						log.Panicln(err)
					}

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
