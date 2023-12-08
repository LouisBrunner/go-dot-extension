// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ScriptCreateDialog struct {
  obj gdc.ObjectPtr
}

func createScriptCreateDialog(obj gdc.ObjectPtr) *ScriptCreateDialog {
  return &ScriptCreateDialog{
    obj: obj,
  }
}

func (me *ScriptCreateDialog) BaseClass() string {
  return "ScriptCreateDialog"
}
