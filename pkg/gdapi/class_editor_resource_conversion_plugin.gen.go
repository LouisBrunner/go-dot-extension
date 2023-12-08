// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorResourceConversionPlugin struct {
  obj gdc.ObjectPtr
}

func createEditorResourceConversionPlugin(obj gdc.ObjectPtr) *EditorResourceConversionPlugin {
  return &EditorResourceConversionPlugin{
    obj: obj,
  }
}

func (me *EditorResourceConversionPlugin) BaseClass() string {
  return "EditorResourceConversionPlugin"
}
