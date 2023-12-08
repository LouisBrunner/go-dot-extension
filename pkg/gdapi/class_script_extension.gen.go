// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ScriptExtension struct {
  obj gdc.ObjectPtr
}

func createScriptExtension(obj gdc.ObjectPtr) *ScriptExtension {
  return &ScriptExtension{
    obj: obj,
  }
}

func (me *ScriptExtension) BaseClass() string {
  return "ScriptExtension"
}
