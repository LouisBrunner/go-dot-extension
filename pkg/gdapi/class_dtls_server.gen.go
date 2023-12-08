// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type DTLSServer struct {
  obj gdc.ObjectPtr
}

func createDTLSServer(obj gdc.ObjectPtr) *DTLSServer {
  return &DTLSServer{
    obj: obj,
  }
}

func (me *DTLSServer) BaseClass() string {
  return "DTLSServer"
}
