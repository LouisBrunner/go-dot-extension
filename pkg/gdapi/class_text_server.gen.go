// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextServer struct {
  obj gdc.ObjectPtr
}

func createTextServer(obj gdc.ObjectPtr) *TextServer {
  return &TextServer{
    obj: obj,
  }
}

func (me *TextServer) BaseClass() string {
  return "TextServer"
}
