// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorNode3DGizmoPlugin struct {
  obj gdc.ObjectPtr
}

func createEditorNode3DGizmoPlugin(obj gdc.ObjectPtr) *EditorNode3DGizmoPlugin {
  return &EditorNode3DGizmoPlugin{
    obj: obj,
  }
}

func (me *EditorNode3DGizmoPlugin) BaseClass() string {
  return "EditorNode3DGizmoPlugin"
}
