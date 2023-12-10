// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type JSONRPC struct {
  obj gdc.ObjectPtr
}

func (me *JSONRPC) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *JSONRPC) BaseClass() string {
  return "JSONRPC"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scope.AsCTypePtr()), gdc.ConstTypePtr(target.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *JSONRPC) ProcessAction(action Variant, recurse bool, ) Variant {
  classNameV := StringNameFromStr("JSONRPC")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("process_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2963479484) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), gdc.ConstTypePtr(&recurse), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *JSONRPC) ProcessString(action String, ) String {
  classNameV := StringNameFromStr("JSONRPC")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("process_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1703090593) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *JSONRPC) MakeRequest(method String, params Variant, id Variant, ) Dictionary {
  classNameV := StringNameFromStr("JSONRPC")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_request")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3423508980) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(method.AsCTypePtr()), gdc.ConstTypePtr(params.AsCTypePtr()), gdc.ConstTypePtr(id.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *JSONRPC) MakeResponse(result Variant, id Variant, ) Dictionary {
  classNameV := StringNameFromStr("JSONRPC")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_response")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 5053918) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(result.AsCTypePtr()), gdc.ConstTypePtr(id.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *JSONRPC) MakeNotification(method String, params Variant, ) Dictionary {
  classNameV := StringNameFromStr("JSONRPC")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_notification")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2949127017) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(method.AsCTypePtr()), gdc.ConstTypePtr(params.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *JSONRPC) MakeResponseError(code int, message String, id Variant, ) Dictionary {
  classNameV := StringNameFromStr("JSONRPC")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_response_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 928596297) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&code), gdc.ConstTypePtr(message.AsCTypePtr()), gdc.ConstTypePtr(id.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties