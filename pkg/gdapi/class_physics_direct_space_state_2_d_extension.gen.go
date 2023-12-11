// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsDirectSpaceState2DExtension struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectSpaceState2DExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectSpaceState2DExtension) BaseClass() string {
  return "PhysicsDirectSpaceState2DExtension"
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
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
