// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BaseButton struct {
  obj gdc.ObjectPtr
}

func (me *BaseButton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *BaseButton) BaseClass() string {
  return "BaseButton"
}

type BaseButtonDrawMode int
const (
  BaseButtonDrawNormal BaseButtonDrawMode = 0
  BaseButtonDrawPressed BaseButtonDrawMode = 1
  BaseButtonDrawHover BaseButtonDrawMode = 2
  BaseButtonDrawDisabled BaseButtonDrawMode = 3
  BaseButtonDrawHoverPressed BaseButtonDrawMode = 4
)

type BaseButtonActionMode int
const (
  BaseButtonActionModeButtonPress BaseButtonActionMode = 0
  BaseButtonActionModeButtonRelease BaseButtonActionMode = 1
)
