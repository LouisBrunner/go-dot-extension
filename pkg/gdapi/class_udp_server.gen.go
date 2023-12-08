// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type UDPServer struct {
  obj gdc.ObjectPtr
}

func createUDPServer(obj gdc.ObjectPtr) *UDPServer {
  return &UDPServer{
    obj: obj,
  }
}

func (me *UDPServer) BaseClass() string {
  return "UDPServer"
}
