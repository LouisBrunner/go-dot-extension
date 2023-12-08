// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Texture2DArray struct {
  obj gdc.ObjectPtr
}

func (me *Texture2DArray) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Texture2DArray) BaseClass() string {
  return "Texture2DArray"
}
