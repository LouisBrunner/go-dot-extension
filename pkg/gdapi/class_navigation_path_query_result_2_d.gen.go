// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationPathQueryResult2D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationPathQueryResult2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationPathQueryResult2D) BaseClass() string {
  return "NavigationPathQueryResult2D"
}

type NavigationPathQueryResult2DPathSegmentType int
const (
  NavigationPathQueryResult2DPathSegmentTypeRegion NavigationPathQueryResult2DPathSegmentType = 0
  NavigationPathQueryResult2DPathSegmentTypeLink NavigationPathQueryResult2DPathSegmentType = 1
)
