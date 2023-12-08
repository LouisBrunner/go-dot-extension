// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorSettings struct {
  obj gdc.ObjectPtr
}

func (me *EditorSettings) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorSettings) BaseClass() string {
  return "EditorSettings"
}

const (
  EditorSettingsNOTIFICATION_EDITOR_SETTINGS_CHANGED = 10000
)
