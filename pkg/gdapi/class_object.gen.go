// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Object struct {
  obj gdc.ObjectPtr
}

func (me *Object) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Object) BaseClass() string {
  return "Object"
}

const (
  ObjectNOTIFICATION_POSTINITIALIZE = 0
  ObjectNOTIFICATION_PREDELETE = 1
)

type ObjectConnectFlags int
const (
  ObjectConnectDeferred ObjectConnectFlags = 1
  ObjectConnectPersist ObjectConnectFlags = 2
  ObjectConnectOneShot ObjectConnectFlags = 4
  ObjectConnectReferenceCounted ObjectConnectFlags = 8
)
