// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ScriptLanguageExtension struct {
  obj gdc.ObjectPtr
}

func createScriptLanguageExtension(obj gdc.ObjectPtr) *ScriptLanguageExtension {
  return &ScriptLanguageExtension{
    obj: obj,
  }
}

func (me *ScriptLanguageExtension) BaseClass() string {
  return "ScriptLanguageExtension"
}
