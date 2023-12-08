// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorResourceTooltipPlugin struct {
  obj gdc.ObjectPtr
}

func createEditorResourceTooltipPlugin(obj gdc.ObjectPtr) *EditorResourceTooltipPlugin {
  return &EditorResourceTooltipPlugin{
    obj: obj,
  }
}

func (me *EditorResourceTooltipPlugin) BaseClass() string {
  return "EditorResourceTooltipPlugin"
}
