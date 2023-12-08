// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorInspectorPlugin struct {
  obj gdc.ObjectPtr
}

func (me *EditorInspectorPlugin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorInspectorPlugin) BaseClass() string {
  return "EditorInspectorPlugin"
}
