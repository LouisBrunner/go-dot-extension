// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Skeleton2D struct {
  obj gdc.ObjectPtr
}

func (me *Skeleton2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Skeleton2D) BaseClass() string {
  return "Skeleton2D"
}
