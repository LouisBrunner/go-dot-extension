// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorNode3DGizmo struct {
  obj gdc.ObjectPtr
}

func createEditorNode3DGizmo(obj gdc.ObjectPtr) *EditorNode3DGizmo {
  return &EditorNode3DGizmo{
    obj: obj,
  }
}

func (me *EditorNode3DGizmo) BaseClass() string {
  return "EditorNode3DGizmo"
}
