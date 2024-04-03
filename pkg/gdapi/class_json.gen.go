// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type JSON struct {
  Resource
}

func (me *JSON) BaseClass() string {
  return "JSON"
}

func NewJSON() *JSON {
  str := StringNameFromStr("JSON") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &JSON{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *JSON) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *JSON) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *JSON) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  JSONStringify(data Variant, indent String, sort_keys bool, full_precision bool, ) String {
  classNameV := StringNameFromStr("JSON")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stringify")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 462733549) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(data.AsCTypePtr()), gdc.ConstTypePtr(indent.AsCTypePtr()), gdc.ConstTypePtr(&sort_keys), gdc.ConstTypePtr(&full_precision), }
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  JSONParseString(json_string String, ) Variant {
  classNameV := StringNameFromStr("JSON")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("parse_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 309047738) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(json_string.AsCTypePtr()), }
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *JSON) Parse(json_text String, keep_text bool, ) Error {
  classNameV := StringNameFromStr("JSON")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("parse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 885841341) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(json_text.AsCTypePtr()), gdc.ConstTypePtr(&keep_text), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *JSON) GetData() Variant {
  classNameV := StringNameFromStr("JSON")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1214101251) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *JSON) SetData(data Variant, )  {
  classNameV := StringNameFromStr("JSON")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1114965689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(data.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *JSON) GetParsedText() String {
  classNameV := StringNameFromStr("JSON")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parsed_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *JSON) GetErrorLine() int64 {
  classNameV := StringNameFromStr("JSON")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_error_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *JSON) GetErrorMessage() String {
  classNameV := StringNameFromStr("JSON")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_error_message")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
