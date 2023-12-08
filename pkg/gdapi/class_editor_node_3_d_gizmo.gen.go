// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorNode3DGizmo struct {
  obj gdc.ObjectPtr
}

func (me *EditorNode3DGizmo) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorNode3DGizmo) BaseClass() string {
  return "EditorNode3DGizmo"
}
