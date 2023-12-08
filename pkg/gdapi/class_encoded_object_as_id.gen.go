// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EncodedObjectAsID struct {
  obj gdc.ObjectPtr
}

func createEncodedObjectAsID(obj gdc.ObjectPtr) *EncodedObjectAsID {
  return &EncodedObjectAsID{
    obj: obj,
  }
}

func (me *EncodedObjectAsID) BaseClass() string {
  return "EncodedObjectAsID"
}
