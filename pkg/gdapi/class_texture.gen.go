// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Texture struct {
  obj gdc.ObjectPtr
}

func (me *Texture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Texture) BaseClass() string {
  return "Texture"
}
