// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ScriptEditorBase struct {
  obj gdc.ObjectPtr
}

func createScriptEditorBase(obj gdc.ObjectPtr) *ScriptEditorBase {
  return &ScriptEditorBase{
    obj: obj,
  }
}

func (me *ScriptEditorBase) BaseClass() string {
  return "ScriptEditorBase"
}
