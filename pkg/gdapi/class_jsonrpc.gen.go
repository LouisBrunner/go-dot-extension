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

type JSONRPC struct {
  Object
}

func (me *JSONRPC) BaseClass() string {
  return "JSONRPC"
}

func NewJSONRPC() *JSONRPC {
  str := StringNameFromStr("JSONRPC") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &JSONRPC{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type JSONRPCErrorCode int
const (
  JSONRPCErrorCodeParseError JSONRPCErrorCode = -32700
  JSONRPCErrorCodeInvalidRequest JSONRPCErrorCode = -32600
  JSONRPCErrorCodeMethodNotFound JSONRPCErrorCode = -32601
  JSONRPCErrorCodeInvalidParams JSONRPCErrorCode = -32602
  JSONRPCErrorCodeInternalError JSONRPCErrorCode = -32603
)

func (me *JSONRPC) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *JSONRPC) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *JSONRPC) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *JSONRPC) SetScope(scope String, target Object, )  {
  classNameV := StringNameFromStr("JSONRPC")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scope")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2572618360) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scope.AsCTypePtr(), target.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *JSONRPC) ProcessAction(action Variant, recurse bool, ) Variant {
  classNameV := StringNameFromStr("JSONRPC")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("process_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2963479484) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&recurse) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&recurse)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *JSONRPC) ProcessString(action String, ) String {
  classNameV := StringNameFromStr("JSONRPC")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("process_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1703090593) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *JSONRPC) MakeRequest(method String, params Variant, id Variant, ) Dictionary {
  classNameV := StringNameFromStr("JSONRPC")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_request")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3423508980) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{method.AsCTypePtr(), params.AsCTypePtr(), id.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *JSONRPC) MakeResponse(result Variant, id Variant, ) Dictionary {
  classNameV := StringNameFromStr("JSONRPC")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_response")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 5053918) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{result.AsCTypePtr(), id.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *JSONRPC) MakeNotification(method String, params Variant, ) Dictionary {
  classNameV := StringNameFromStr("JSONRPC")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_notification")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2949127017) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{method.AsCTypePtr(), params.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *JSONRPC) MakeResponseError(code int64, message String, id Variant, ) Dictionary {
  classNameV := StringNameFromStr("JSONRPC")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_response_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 928596297) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&code) , message.AsCTypePtr(), id.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&code)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
