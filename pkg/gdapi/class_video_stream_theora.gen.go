// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type VideoStreamTheora struct {
  VideoStream
}

func (me *VideoStreamTheora) BaseClass() string {
  return "VideoStreamTheora"
}

func NewVideoStreamTheora() *VideoStreamTheora {
  str := StringNameFromStr("VideoStreamTheora") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VideoStreamTheora{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VideoStreamTheora) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VideoStreamTheora) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VideoStreamTheora) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
