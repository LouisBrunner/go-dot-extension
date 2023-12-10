// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsPointQueryParameters2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsPointQueryParameters2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsPointQueryParameters2D) BaseClass() string {
  return "PhysicsPointQueryParameters2D"
}



// Enums

func (me *PhysicsPointQueryParameters2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsPointQueryParameters2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsPointQueryParameters2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicsPointQueryParameters2D) SetPosition(position Vector2, )  {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsPointQueryParameters2D) GetPosition() Vector2 {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsPointQueryParameters2D) SetCanvasInstanceId(canvas_instance_id int, )  {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_canvas_instance_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&canvas_instance_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsPointQueryParameters2D) GetCanvasInstanceId() int {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_canvas_instance_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsPointQueryParameters2D) SetCollisionMask(collision_mask int, )  {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsPointQueryParameters2D) GetCollisionMask() int {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsPointQueryParameters2D) SetExclude(exclude RID, )  {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_exclude")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(exclude.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsPointQueryParameters2D) GetExclude() RID {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_exclude")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsPointQueryParameters2D) SetCollideWithBodies(enable bool, )  {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collide_with_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsPointQueryParameters2D) IsCollideWithBodiesEnabled() bool {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collide_with_bodies_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsPointQueryParameters2D) SetCollideWithAreas(enable bool, )  {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collide_with_areas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsPointQueryParameters2D) IsCollideWithAreasEnabled() bool {
  classNameV := StringNameFromStr("PhysicsPointQueryParameters2D")
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

func (me *PhysicsPointQueryParameters2D) GetPropPosition() Vector2 {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters2D) SetPropPosition(value Vector2) {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters2D) GetPropCanvasInstanceId() int {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters2D) SetPropCanvasInstanceId(value int) {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters2D) GetPropCollisionMask() int {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters2D) SetPropCollisionMask(value int) {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters2D) GetPropExclude() RID {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters2D) SetPropExclude(value RID) {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters2D) GetPropCollideWithBodies() bool {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters2D) SetPropCollideWithBodies(value bool) {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters2D) GetPropCollideWithAreas() bool {
  panic("TODO: implement")
}

func (me *PhysicsPointQueryParameters2D) SetPropCollideWithAreas(value bool) {
  panic("TODO: implement")
}