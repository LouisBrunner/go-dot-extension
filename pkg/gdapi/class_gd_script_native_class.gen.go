// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GDScriptNativeClass struct {
  obj gdc.ObjectPtr
}

func (me *GDScriptNativeClass) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GDScriptNativeClass) BaseClass() string {
  return "GDScriptNativeClass"
}
