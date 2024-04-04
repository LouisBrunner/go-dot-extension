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

type PhysicsDirectSpaceState2DExtension struct {
  PhysicsDirectSpaceState2D
}

func (me *PhysicsDirectSpaceState2DExtension) BaseClass() string {
  return "PhysicsDirectSpaceState2DExtension"
}

func NewPhysicsDirectSpaceState2DExtension() *PhysicsDirectSpaceState2DExtension {
  str := StringNameFromStr("PhysicsDirectSpaceState2DExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PhysicsDirectSpaceState2DExtension{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PhysicsDirectSpaceState2DExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsDirectSpaceState2DExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectSpaceState2DExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicsDirectSpaceState2DExtension) IsBodyExcludedFromQuery(body RID, ) bool {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState2DExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_body_excluded_from_query")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
