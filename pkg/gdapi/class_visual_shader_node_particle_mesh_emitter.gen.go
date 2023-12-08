// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParticleMeshEmitter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeParticleMeshEmitter(obj gdc.ObjectPtr) *VisualShaderNodeParticleMeshEmitter {
  return &VisualShaderNodeParticleMeshEmitter{
    obj: obj,
  }
}

func (me *VisualShaderNodeParticleMeshEmitter) BaseClass() string {
  return "VisualShaderNodeParticleMeshEmitter"
}
