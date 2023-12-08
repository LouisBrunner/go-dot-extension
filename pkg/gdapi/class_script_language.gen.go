// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ScriptLanguage struct {
  obj gdc.ObjectPtr
}

func createScriptLanguage(obj gdc.ObjectPtr) *ScriptLanguage {
  return &ScriptLanguage{
    obj: obj,
  }
}

func (me *ScriptLanguage) BaseClass() string {
  return "ScriptLanguage"
}
