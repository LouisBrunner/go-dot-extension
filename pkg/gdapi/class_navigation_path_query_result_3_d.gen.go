// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationPathQueryResult3D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationPathQueryResult3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationPathQueryResult3D) BaseClass() string {
  return "NavigationPathQueryResult3D"
}

type NavigationPathQueryResult3DPathSegmentType int
const (
  NavigationPathQueryResult3DPathSegmentTypeRegion NavigationPathQueryResult3DPathSegmentType = 0
  NavigationPathQueryResult3DPathSegmentTypeLink NavigationPathQueryResult3DPathSegmentType = 1
)
