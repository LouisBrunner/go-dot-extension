// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PolygonPathFinder struct {
  obj gdc.ObjectPtr
}

func createPolygonPathFinder(obj gdc.ObjectPtr) *PolygonPathFinder {
  return &PolygonPathFinder{
    obj: obj,
  }
}

func (me *PolygonPathFinder) BaseClass() string {
  return "PolygonPathFinder"
}
