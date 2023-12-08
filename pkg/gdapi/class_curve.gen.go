// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Curve struct {
  obj gdc.ObjectPtr
}

func (me *Curve) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Curve) BaseClass() string {
  return "Curve"
}

type CurveTangentMode int
const (
  CurveTangentFree CurveTangentMode = 0
  CurveTangentLinear CurveTangentMode = 1
  CurveTangentModeCount CurveTangentMode = 2
)
