// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParticleSphereEmitter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeParticleSphereEmitter(obj gdc.ObjectPtr) *VisualShaderNodeParticleSphereEmitter {
  return &VisualShaderNodeParticleSphereEmitter{
    obj: obj,
  }
}

func (me *VisualShaderNodeParticleSphereEmitter) BaseClass() string {
  return "VisualShaderNodeParticleSphereEmitter"
}
