// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type JavaScriptBridge struct {
  obj gdc.ObjectPtr
}

func createJavaScriptBridge(obj gdc.ObjectPtr) *JavaScriptBridge {
  return &JavaScriptBridge{
    obj: obj,
  }
}

func (me *JavaScriptBridge) BaseClass() string {
  return "JavaScriptBridge"
}
