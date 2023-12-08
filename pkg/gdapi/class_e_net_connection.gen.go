// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ENetConnection struct {
  obj gdc.ObjectPtr
}

func createENetConnection(obj gdc.ObjectPtr) *ENetConnection {
  return &ENetConnection{
    obj: obj,
  }
}

func (me *ENetConnection) BaseClass() string {
  return "ENetConnection"
}
