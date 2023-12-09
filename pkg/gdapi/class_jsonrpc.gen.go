// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *JSONRPC) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *JSONRPC) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *JSONRPC) SetScope(scope String, target Object, )  {
  panic("TODO: implement")
}

func  (me *JSONRPC) ProcessAction(action Variant, recurse bool, )  {
  panic("TODO: implement")
}

func  (me *JSONRPC) ProcessString(action String, )  {
  panic("TODO: implement")
}

func  (me *JSONRPC) MakeRequest(method String, params Variant, id Variant, )  {
  panic("TODO: implement")
}

func  (me *JSONRPC) MakeResponse(result Variant, id Variant, )  {
  panic("TODO: implement")
}

func  (me *JSONRPC) MakeNotification(method String, params Variant, )  {
  panic("TODO: implement")
}

func  (me *JSONRPC) MakeResponseError(code int, message String, id Variant, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
