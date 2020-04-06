package comm

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

var (
	metaByFullName map[string]*MessageMeta
	metaByID map[int]*MessageMeta
	metaByType map[reflect.Type]*MessageMeta
)

func init() {
	metaByFullName = make(map[string]*MessageMeta)
	metaByID = make(map[int]*MessageMeta)
	metaByType = make(map[reflect.Type]*MessageMeta)
}

type context struct {
	name  string
	value interface{}
}

type MessageMeta struct {
	Codec Codec
	Type  reflect.Type
	MsgId int

	ctxGuard sync.Mutex
	ctxList []context
}

func (self *MessageMeta) TypeName() string {
	return self.Type.Name()
}

func (self *MessageMeta) FullName() string {
	var sb strings.Builder

	sb.WriteString(self.Type.PkgPath())
	sb.WriteString(".")
	sb.WriteString(self.Type.Name())

	return sb.String()
}

func (self *MessageMeta) NewType() interface{} {
	if self.Type == nil {
		return nil
	}

	return reflect.New(self.Type).Interface()
}

func (self *MessageMeta) SetContext(name string, value interface{}) {
	self.ctxGuard.Lock()
	defer self.ctxGuard.Unlock()

	for _, c := range self.ctxList {
		if c.name == name {
			return
		}
	}

	self.ctxList = append(self.ctxList, context{
		name : name,
		value : value,
	})
}

func (self *MessageMeta) GetContext(name string) (interface{}, bool) {
	self.ctxGuard.Lock()
	defer self.ctxGuard.Unlock()

	for _, c := range self.ctxList {
		if c.name == name {
			return c.value, true
		}
	}

	return nil, false
}

func (self *MessageMeta) GetContextAsString(name string, defaultVal string) string {
	if v, ok := self.GetContext(name); ok {
		if str, ok := v.(string); ok {
			return str
		}
	}

	return defaultVal
}

func (self *MessageMeta) GetContextAsInt(name string, defaultVal int) int {
	if v, ok := self.GetContext(name); ok {
		if num, ok := v.(int); ok {
			return num
		}
	}

	return defaultVal
}

func RegMessageMeta(meta *MessageMeta) {
	// 统一使用非指针
	if meta.Type.Kind() == reflect.Ptr {
		meta.Type = meta.Type.Elem()
	}

	if meta.MsgId == 0 {
		panic(fmt.Errorf("meta msgid can't be zero"))
	}

	if _, ok := metaByFullName[meta.FullName()]; ok {
		panic(fmt.Errorf("meta %s already exist", meta.FullName()))
	} else {
		metaByFullName[meta.FullName()] = meta
	}

	if _, ok := metaByID[meta.MsgId]; !ok {
		panic(fmt.Errorf("msgid %d already exist", meta.MsgId))
	} else {
		metaByID[meta.MsgId] = meta
	}

	if _, ok := metaByType[meta.Type]; ok {
		panic(fmt.Errorf("meta type %s already exist", meta.FullName()))
	} else {
		metaByType[meta.Type] = meta
	}
}

func MessageMetaByFullName(name string) *MessageMeta {
	if m, ok := metaByFullName[name]; ok {
		return m
	}

	return nil
}

func MessageMetaByID(id int) *MessageMeta {
	if m, ok := metaByID[id]; ok {
		return m
	}

	return nil
}

func MessageMetaByType(t reflect.Type) *MessageMeta {
	if t == nil {
		return nil
	}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if m, ok := metaByType[t]; ok {
		return m
	}

	return nil
}

func MessageMetaByMsg(msg interface{}) *MessageMeta {
	if msg == nil {
		return nil
	}

	return MessageMetaByType(reflect.TypeOf(msg))
}

func MessageToID(msg interface{}) int {
	if msg == nil {
		return 0
	}

	meta := MessageMetaByMsg(msg)
	if meta != nil {
		return meta.MsgId
	} else {
		return 0
	}
}

func MessageToName(msg interface{}) string {
	if msg == nil {
		return ""
	}

	meta := MessageMetaByMsg(msg)
	if meta == nil {
		return ""
	} else {
		return meta.TypeName()
	}
}

func MessageSize(msg interface{}) int {
	if msg == nil {
		return 0
	}

	meta := MessageMetaByMsg(msg)
	if meta == nil {
		return 0
	}

	data, err := meta.Codec.Encode(msg, nil)
	if err != nil {
		return 0
	} else {
		return len(data.([]byte))
	}
}

func MessageToString(msg interface{}) string {
	if msg == nil {
		return ""
	}

	if stringer, ok := msg.(interface{ String() string }); ok {
		return stringer.String()
	} else {
		return ""
	}
}