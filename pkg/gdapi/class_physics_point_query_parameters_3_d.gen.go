// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsPointQueryParameters3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsPointQueryParameters3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsPointQueryParameters3D) BaseClass() string {
  return "PhysicsPointQueryParameters3D"
}



// Enums

func (me *PhysicsPointQueryParameters3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsPointQueryParameters3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsPointQueryParameters3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicsPointQueryParameters3D) SetPosition(position Vector3, )  {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsPointQueryParameters3D) GetPosition() Vector3 {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsPointQueryParameters3D) SetCollisionMask(collision_mask int, )  {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsPointQueryParameters3D) GetCollisionMask() int {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsPointQueryParameters3D) SetExclude(exclude RID, )  {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_exclude")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(exclude.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsPointQueryParameters3D) GetExclude() RID {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_exclude")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsPointQueryParameters3D) SetCollideWithBodies(enable bool, )  {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collide_with_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsPointQueryParameters3D) IsCollideWithBodiesEnabled() bool {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collide_with_bodies_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsPointQueryParameters3D) SetCollideWithAreas(enable bool, )  {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collide_with_areas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsPointQueryParameters3D) IsCollideWithAreasEnabled() bool {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collide_with_areas_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *PhysicsPointQueryParameters3D) GetPropPosition() Vector3 {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters3D) SetPropPosition(value Vector3) {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters3D) GetPropCollisionMask() int {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters3D) SetPropCollisionMask(value int) {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters3D) GetPropExclude() RID {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters3D) SetPropExclude(value RID) {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters3D) GetPropCollideWithBodies() bool {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters3D) SetPropCollideWithBodies(value bool) {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters3D) GetPropCollideWithAreas() bool {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters3D) SetPropCollideWithAreas(value bool) {
  panic("TODO: implement")
}