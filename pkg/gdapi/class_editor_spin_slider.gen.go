// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorSpinSlider struct {
  obj gdc.ObjectPtr
}

func createEditorSpinSlider(obj gdc.ObjectPtr) *EditorSpinSlider {
  return &EditorSpinSlider{
    obj: obj,
  }
}

func (me *EditorSpinSlider) BaseClass() string {
  return "EditorSpinSlider"
}
