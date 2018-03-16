package codec

import (
	"bytes"
	"fmt"
	"path"
	"reflect"
)

type MessageMeta struct {
	Type  reflect.Type
	Name  string
	Codec Codec
}

var (
	metaByName = map[string]*MessageMeta{}
	metaByType = map[reflect.Type]*MessageMeta{}
)

func RegisterMessageMeta(codecName string, msg interface{}) {
	msgType := reflect.TypeOf(msg)
	name := msgType.Elem().Name()

	meta := &MessageMeta{
		Type:  msgType,
		Name:  name,
		Codec: CodecManager().GetCodec(codecName),
	}

	if meta.Codec == nil {
		panic("codec not register! " + codecName)
	}

	if _, ok := metaByName[name]; ok {
		panic("duplicate message meta register by name: " + name)
	}

	if _, ok := metaByType[msgType]; ok {
		panic(fmt.Sprintf("duplicate message meta register by type: %s", meta.Name))
	}

	metaByName[name] = meta
	metaByType[msgType] = meta
}

func MessageMetaByName(name string) *MessageMeta {
	if v, ok := metaByName[name]; ok {
		return v
	}

	return nil
}

func MessageMetaByType(t reflect.Type) *MessageMeta {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if v, ok := metaByType[t]; ok {
		return v
	}

	return nil
}

func MessageFullName(rtype reflect.Type) string {
	if rtype == nil {
		panic("empty msg type")
	}

	if rtype.Kind() == reflect.Ptr {
		rtype = rtype.Elem()
	}

	var b bytes.Buffer
	b.WriteString(path.Base(rtype.PkgPath()))
	b.WriteString(".")
	b.WriteString(rtype.Name())

	return b.String()
}

func BuildMessage(name string) interface{} {
	var meta = MessageMetaByName(name)
	if meta != nil {
		return reflect.New(meta.Type).Interface()
	}
	return nil
}
