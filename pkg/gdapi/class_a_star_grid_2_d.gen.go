// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AStarGrid2D struct {
  obj gdc.ObjectPtr
}

func (me *AStarGrid2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AStarGrid2D) BaseClass() string {
  return "AStarGrid2D"
}
