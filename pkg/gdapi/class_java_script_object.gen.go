// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type JavaScriptObject struct {
  obj gdc.ObjectPtr
}

func createJavaScriptObject(obj gdc.ObjectPtr) *JavaScriptObject {
  return &JavaScriptObject{
    obj: obj,
  }
}

func (me *JavaScriptObject) BaseClass() string {
  return "JavaScriptObject"
}
