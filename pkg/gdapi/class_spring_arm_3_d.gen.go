// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SpringArm3D struct {
  obj gdc.ObjectPtr
}

func (me *SpringArm3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SpringArm3D) BaseClass() string {
  return "SpringArm3D"
}



// Enums

func (me *SpringArm3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SpringArm3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SpringArm3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SpringArm3D) GetHitLength() float32 {
  classNameV := StringNameFromStr("SpringArm3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hit_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SpringArm3D) SetLength(length float32, )  {
  classNameV := StringNameFromStr("SpringArm3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SpringArm3D) GetLength() float32 {
  classNameV := StringNameFromStr("SpringArm3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SpringArm3D) SetShape(shape Shape3D, )  {
  classNameV := StringNameFromStr("SpringArm3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1549710052) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SpringArm3D) GetShape() Shape3D {
  classNameV := StringNameFromStr("SpringArm3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3214262478) // FIXME: should cache?
  var ret Shape3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SpringArm3D) AddExcludedObject(RID RID, )  {
  classNameV := StringNameFromStr("SpringArm3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_excluded_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(RID.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SpringArm3D) RemoveExcludedObject(RID RID, ) bool {
  classNameV := StringNameFromStr("SpringArm3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_excluded_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3521089500) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(RID.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SpringArm3D) ClearExcludedObjects()  {
  classNameV := StringNameFromStr("SpringArm3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_excluded_objects")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SpringArm3D) SetCollisionMask(mask int, )  {
  classNameV := StringNameFromStr("SpringArm3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SpringArm3D) GetCollisionMask() int {
  classNameV := StringNameFromStr("SpringArm3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SpringArm3D) SetMargin(margin float32, )  {
  classNameV := StringNameFromStr("SpringArm3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SpringArm3D) GetMargin() float32 {
  classNameV := StringNameFromStr("SpringArm3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *SpringArm3D) GetPropCollisionMask() int {
  panic("TODO: implement")
}

func (me *SpringArm3D) SetPropCollisionMask(value int) {
  panic("TODO: implement")
}

func (me *SpringArm3D) GetPropShape() Shape3D {
  panic("TODO: implement")
}

func (me *SpringArm3D) SetPropShape(value Shape3D) {
  panic("TODO: implement")
}

func (me *SpringArm3D) GetPropSpringLength() float32 {
  panic("TODO: implement")
}

func (me *SpringArm3D) SetPropSpringLength(value float32) {
  panic("TODO: implement")
}

func (me *SpringArm3D) GetPropMargin() float32 {
  panic("TODO: implement")
}

func (me *SpringArm3D) SetPropMargin(value float32) {
  panic("TODO: implement")
}