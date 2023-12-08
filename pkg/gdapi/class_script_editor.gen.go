// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ScriptEditor struct {
  obj gdc.ObjectPtr
}

func createScriptEditor(obj gdc.ObjectPtr) *ScriptEditor {
  return &ScriptEditor{
    obj: obj,
  }
}

func (me *ScriptEditor) BaseClass() string {
  return "ScriptEditor"
}
