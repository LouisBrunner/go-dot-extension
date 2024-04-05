// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForJSONList struct {
  fnStringify gdc.MethodBindPtr
  fnParseString gdc.MethodBindPtr
  fnParse gdc.MethodBindPtr
  fnGetData gdc.MethodBindPtr
  fnSetData gdc.MethodBindPtr
  fnGetParsedText gdc.MethodBindPtr
  fnGetErrorLine gdc.MethodBindPtr
  fnGetErrorMessage gdc.MethodBindPtr
}

var ptrsForJSON ptrsForJSONList

func initJSONPtrs(iface gdc.Interface) {

  className := StringNameFromStr("JSON")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("stringify")
    defer methodName.Destroy()
    ptrsForJSON.fnStringify = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 462733549))
  }
  {
    methodName := StringNameFromStr("parse_string")
    defer methodName.Destroy()
    ptrsForJSON.fnParseString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 309047738))
  }
  {
    methodName := StringNameFromStr("parse")
    defer methodName.Destroy()
    ptrsForJSON.fnParse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 885841341))
  }
  {
    methodName := StringNameFromStr("get_data")
    defer methodName.Destroy()
    ptrsForJSON.fnGetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1214101251))
  }
  {
    methodName := StringNameFromStr("set_data")
    defer methodName.Destroy()
    ptrsForJSON.fnSetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1114965689))
  }
  {
    methodName := StringNameFromStr("get_parsed_text")
    defer methodName.Destroy()
    ptrsForJSON.fnGetParsedText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_error_line")
    defer methodName.Destroy()
    ptrsForJSON.fnGetErrorLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_error_message")
    defer methodName.Destroy()
    ptrsForJSON.fnGetErrorMessage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
}

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
  cargs := []gdc.ConstTypePtr{data.AsCTypePtr(), indent.AsCTypePtr(), gdc.ConstTypePtr(&sort_keys) , gdc.ConstTypePtr(&full_precision) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&sort_keys)
  pinner.Pin(&full_precision)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJSON.fnStringify), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  JSONParseString(json_string String, ) Variant {
  cargs := []gdc.ConstTypePtr{json_string.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJSON.fnParseString), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *JSON) Parse(json_text String, keep_text bool, ) Error {
  cargs := []gdc.ConstTypePtr{json_text.AsCTypePtr(), gdc.ConstTypePtr(&keep_text) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&keep_text)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJSON.fnParse), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *JSON) GetData() Variant {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJSON.fnGetData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *JSON) SetData(data Variant, )  {
  cargs := []gdc.ConstTypePtr{data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJSON.fnSetData), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *JSON) GetParsedText() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJSON.fnGetParsedText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *JSON) GetErrorLine() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJSON.fnGetErrorLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *JSON) GetErrorMessage() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJSON.fnGetErrorMessage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
