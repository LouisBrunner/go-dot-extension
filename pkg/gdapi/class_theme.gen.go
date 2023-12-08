// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Theme struct {
  obj gdc.ObjectPtr
}

func (me *Theme) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Theme) BaseClass() string {
  return "Theme"
}

type ThemeDataType int
const (
  ThemeDataTypeColor ThemeDataType = 0
  ThemeDataTypeConstant ThemeDataType = 1
  ThemeDataTypeFont ThemeDataType = 2
  ThemeDataTypeFontSize ThemeDataType = 3
  ThemeDataTypeIcon ThemeDataType = 4
  ThemeDataTypeStylebox ThemeDataType = 5
  ThemeDataTypeMax ThemeDataType = 6
)
