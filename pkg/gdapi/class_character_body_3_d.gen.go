// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CharacterBody3D struct {
  obj gdc.ObjectPtr
}

func createCharacterBody3D(obj gdc.ObjectPtr) *CharacterBody3D {
  return &CharacterBody3D{
    obj: obj,
  }
}

func (me *CharacterBody3D) BaseClass() string {
  return "CharacterBody3D"
}
