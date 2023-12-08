// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CharacterBody3D struct {
  obj gdc.ObjectPtr
}

func (me *CharacterBody3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CharacterBody3D) BaseClass() string {
  return "CharacterBody3D"
}
