// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextMesh struct {
  obj gdc.ObjectPtr
}

func (me *TextMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextMesh) BaseClass() string {
  return "TextMesh"
}
