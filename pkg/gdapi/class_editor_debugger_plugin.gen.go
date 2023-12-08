// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorDebuggerPlugin struct {
  obj gdc.ObjectPtr
}

func createEditorDebuggerPlugin(obj gdc.ObjectPtr) *EditorDebuggerPlugin {
  return &EditorDebuggerPlugin{
    obj: obj,
  }
}

func (me *EditorDebuggerPlugin) BaseClass() string {
  return "EditorDebuggerPlugin"
}
