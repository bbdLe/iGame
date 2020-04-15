package peer

import (
	"reflect"
	"sync"
)

type ctx struct {
	name interface{}
	value interface{}
}

type CoreContextSet struct {
	ctxes []*ctx
	ctxesGuard sync.Mutex
}

func (self *CoreContextSet) SetContext(key interface{}, value interface{}) {
	self.ctxesGuard.Lock()
	defer self.ctxesGuard.Unlock()

	for _, c := range self.ctxes {
		if c.name == key {
			c.value = value
			return
		}
	}

	self.ctxes = append(self.ctxes, &ctx{
		name : key,
		value : value,
	})
}

func (self *CoreContextSet) GetContext(key interface{}) (interface{}, bool) {
	self.ctxesGuard.Lock()
	defer self.ctxesGuard.Unlock()

	for _, c := range self.ctxes {
		if c.name == key {
			return c.value, true
		}
	}

	return nil, false
}

func (self *CoreContextSet) FetchContext(key interface{}, valuePtr interface{}) bool {
	val, exist := self.GetContext(key)
	if !exist {
		return exist
	}

	switch rawValue := valuePtr.(type) {
	case *string: {
		*rawValue = val.(string)
	}
	case *int: {
		*rawValue = val.(int)
	}
	case *int16: {
		*rawValue = val.(int16)
	}
	case *int32: {
		*rawValue = val.(int32)
	}
	case *int64: {
		*rawValue = val.(int64)
	}
	case *uint: {
		*rawValue = val.(uint)
	}
	case *uint16: {
		*rawValue = val.(uint16)
	}
	case *uint32: {
		*rawValue = val.(uint32)
	}
	case *uint64: {
		*rawValue = val.(uint64)
	}
	case *float32: {
		*rawValue = val.(float32)
	}
	case *float64: {
		*rawValue = val.(float64)
	}
	case *bool: {
		*rawValue = val.(bool)
	}
	case *[]byte: {
		*rawValue = val.([]byte)
	}
	default:
		v := reflect.Indirect(reflect.ValueOf(valuePtr))

		if val == nil {
			v.Set(reflect.Zero(v.Type()))
		} else {
			v.Set(reflect.ValueOf(val))
		}
	}

	return true
}
