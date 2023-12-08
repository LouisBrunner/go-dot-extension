// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AStarGrid2D struct {
  obj gdc.ObjectPtr
}

func createAStarGrid2D(obj gdc.ObjectPtr) *AStarGrid2D {
  return &AStarGrid2D{
    obj: obj,
  }
}

func (me *AStarGrid2D) BaseClass() string {
  return "AStarGrid2D"
}
