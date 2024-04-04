// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type PhysicsServer2DExtension struct {
  PhysicsServer2D
}

func (me *PhysicsServer2DExtension) BaseClass() string {
  return "PhysicsServer2DExtension"
}

func NewPhysicsServer2DExtension() *PhysicsServer2DExtension {
  str := StringNameFromStr("PhysicsServer2DExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PhysicsServer2DExtension{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PhysicsServer2DExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsServer2DExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsServer2DExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicsServer2DExtension) BodyTestMotionIsExcludingBody(body RID, ) bool {
  classNameV := StringNameFromStr("PhysicsServer2DExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_test_motion_is_excluding_body")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsServer2DExtension) BodyTestMotionIsExcludingObject(object int64, ) bool {
  classNameV := StringNameFromStr("PhysicsServer2DExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_test_motion_is_excluding_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&object) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&object)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
