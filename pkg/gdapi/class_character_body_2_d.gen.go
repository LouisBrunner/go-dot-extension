// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CharacterBody2D struct {
  obj gdc.ObjectPtr
}

func createCharacterBody2D(obj gdc.ObjectPtr) *CharacterBody2D {
  return &CharacterBody2D{
    obj: obj,
  }
}

func (me *CharacterBody2D) BaseClass() string {
  return "CharacterBody2D"
}
