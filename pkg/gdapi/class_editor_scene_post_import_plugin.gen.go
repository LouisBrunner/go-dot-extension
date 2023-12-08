// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorScenePostImportPlugin struct {
  obj gdc.ObjectPtr
}

func (me *EditorScenePostImportPlugin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorScenePostImportPlugin) BaseClass() string {
  return "EditorScenePostImportPlugin"
}
