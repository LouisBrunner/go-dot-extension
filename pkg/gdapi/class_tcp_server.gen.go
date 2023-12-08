// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TCPServer struct {
  obj gdc.ObjectPtr
}

func (me *TCPServer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TCPServer) BaseClass() string {
  return "TCPServer"
}
