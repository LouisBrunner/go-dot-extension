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

type JSONRPCErrorCode int
const (
  JSONRPCErrorCodeParseError JSONRPCErrorCode = -32700
  JSONRPCErrorCodeInvalidRequest JSONRPCErrorCode = -32600
  JSONRPCErrorCodeMethodNotFound JSONRPCErrorCode = -32601
  JSONRPCErrorCodeInvalidParams JSONRPCErrorCode = -32602
  JSONRPCErrorCodeInternalError JSONRPCErrorCode = -32603
)

func  (me *JSONRPC) SetScope(scope String, target Object, ) { // TODO: return value
  // TODO: implement
}

func  (me *JSONRPC) ProcessAction(action Variant, recurse bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *JSONRPC) ProcessString(action String, ) { // TODO: return value
  // TODO: implement
}

func  (me *JSONRPC) MakeRequest(method String, params Variant, id Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *JSONRPC) MakeResponse(result Variant, id Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *JSONRPC) MakeNotification(method String, params Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *JSONRPC) MakeResponseError(code int, message String, id Variant, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
