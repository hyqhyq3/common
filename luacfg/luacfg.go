package luacfg

import (
	"errors"
	"log"
	"reflect"
	"strconv"

	"github.com/yuin/gopher-lua"
)

func Load(path string, val interface{}) (err error) {
	l := lua.NewState()
	err = l.DoFile(path)
	if err != nil {
		return
	}

	err = load(l.CheckTable(-1), val)
	return
}

func load(l lua.LValue, val interface{}) (err error) {

	v := reflect.ValueOf(val)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
		if v.Kind() == reflect.Ptr && v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
	}
	switch v.Kind() {
	case reflect.Struct:
		if l.Type() != lua.LTTable {
			return errors.New("我们需要table类型")
		}
		tbl := l.(*lua.LTable)
		for fieldIndex := 0; fieldIndex < v.NumField(); fieldIndex++ {
			ft := v.Type().Field(fieldIndex)
			fname := ft.Tag.Get("luacfg")
			f := v.Field(fieldIndex)
			if i, err := strconv.Atoi(fname); err == nil {
				load(tbl.RawGet(lua.LNumber(i)), f.Addr().Interface())
			} else {
				load(tbl.RawGet(lua.LString(fname)), f.Addr().Interface())
			}
		}
	case reflect.Int, reflect.Float32, reflect.Float64:
		if l.Type() != lua.LTNumber {
			return errors.New("我们需要number类型")
		}
		v.Set(reflect.ValueOf(l.(lua.LNumber)).Convert(v.Type()))
	case reflect.Slice:
		if l.Type() != lua.LTTable {
			return errors.New("我们需要table类型")
		}
		tbl := l.(*lua.LTable)
		v.Set(reflect.MakeSlice(reflect.SliceOf(v.Type().Elem()), tbl.Len(), tbl.Len()))
		for i := 0; i < tbl.Len(); i++ {
			load(tbl.RawGet(lua.LNumber(i+1)), v.Index(i).Addr().Interface())
		}
	case reflect.String:
		if l.Type() != lua.LTString {
			return errors.New("我们需要string类型")
		}
		v.Set(reflect.ValueOf(string(l.(lua.LString))))
	case reflect.Map:
		if l.Type() != lua.LTTable {
			return errors.New("我们需要table类型")
		}
		tbl := l.(*lua.LTable)
		v.Set(reflect.MakeMap(reflect.MapOf(v.Type().Key(), v.Type().Elem())))

		tbl.ForEach(func(key, val lua.LValue) {
			mv := reflect.New(v.Type().Elem())
			load(val, mv.Interface())

			mk := reflect.ValueOf(key).Convert(v.Type().Key())
			v.SetMapIndex(mk, mv.Elem())
		})
	default:
		log.Fatal("Unkown type ", v.Type())
	}
	return
}
