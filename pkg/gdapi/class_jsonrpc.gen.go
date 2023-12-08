// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  JSONRPCParseError JSONRPCErrorCode = -32700
  JSONRPCInvalidRequest JSONRPCErrorCode = -32600
  JSONRPCMethodNotFound JSONRPCErrorCode = -32601
  JSONRPCInvalidParams JSONRPCErrorCode = -32602
  JSONRPCInternalError JSONRPCErrorCode = -32603
)
