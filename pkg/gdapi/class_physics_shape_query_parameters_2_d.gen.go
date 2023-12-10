// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsShapeQueryParameters2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsShapeQueryParameters2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsShapeQueryParameters2D) BaseClass() string {
  return "PhysicsShapeQueryParameters2D"
}



// Enums

func (me *PhysicsShapeQueryParameters2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsShapeQueryParameters2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsShapeQueryParameters2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicsShapeQueryParameters2D) SetShape(shape Resource, )  {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 968641751) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsShapeQueryParameters2D) GetShape() Resource {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 121922552) // FIXME: should cache?
  var ret Resource
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsShapeQueryParameters2D) SetShapeRid(shape RID, )  {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shape_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsShapeQueryParameters2D) GetShapeRid() RID {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shape_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsShapeQueryParameters2D) SetTransform(transform Transform2D, )  {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761652528) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsShapeQueryParameters2D) GetTransform() Transform2D {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsShapeQueryParameters2D) SetMotion(motion Vector2, )  {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_motion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(motion.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsShapeQueryParameters2D) GetMotion() Vector2 {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_motion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsShapeQueryParameters2D) SetMargin(margin float32, )  {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsShapeQueryParameters2D) GetMargin() float32 {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsShapeQueryParameters2D) SetCollisionMask(collision_mask int, )  {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsShapeQueryParameters2D) GetCollisionMask() int {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsShapeQueryParameters2D) SetExclude(exclude RID, )  {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_exclude")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(exclude.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsShapeQueryParameters2D) GetExclude() RID {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_exclude")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsShapeQueryParameters2D) SetCollideWithBodies(enable bool, )  {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collide_with_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsShapeQueryParameters2D) IsCollideWithBodiesEnabled() bool {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collide_with_bodies_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsShapeQueryParameters2D) SetCollideWithAreas(enable bool, )  {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collide_with_areas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsShapeQueryParameters2D) IsCollideWithAreasEnabled() bool {
  classNameV := StringNameFromStr("PhysicsShapeQueryParameters2D")
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

func (me *PhysicsShapeQueryParameters2D) GetPropCollisionMask() int {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) SetPropCollisionMask(value int) {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) GetPropExclude() RID {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) SetPropExclude(value RID) {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) GetPropMargin() float32 {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) SetPropMargin(value float32) {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) GetPropMotion() Vector2 {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) SetPropMotion(value Vector2) {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) GetPropShape() Shape2D {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) SetPropShape(value Shape2D) {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) GetPropShapeRid() RID {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) SetPropShapeRid(value RID) {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) GetPropTransform() Transform2D {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) SetPropTransform(value Transform2D) {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) GetPropCollideWithBodies() bool {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) SetPropCollideWithBodies(value bool) {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) GetPropCollideWithAreas() bool {
  panic("TODO: implement")
}

func (me *PhysicsShapeQueryParameters2D) SetPropCollideWithAreas(value bool) {
  panic("TODO: implement")
}