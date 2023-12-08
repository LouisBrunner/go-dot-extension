// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type JSONRPC struct {
  obj gdc.ObjectPtr
}

func createJSONRPC(obj gdc.ObjectPtr) *JSONRPC {
  return &JSONRPC{
    obj: obj,
  }
}

func (me *JSONRPC) BaseClass() string {
  return "JSONRPC"
}
