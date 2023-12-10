// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsDirectSpaceState3DExtension struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectSpaceState3DExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectSpaceState3DExtension) BaseClass() string {
  return "PhysicsDirectSpaceState3DExtension"
}



// Enums

func (me *PhysicsDirectSpaceState3DExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsDirectSpaceState3DExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectSpaceState3DExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicsDirectSpaceState3DExtension) IsBodyExcludedFromQuery(body RID, ) bool {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState3DExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_body_excluded_from_query")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties