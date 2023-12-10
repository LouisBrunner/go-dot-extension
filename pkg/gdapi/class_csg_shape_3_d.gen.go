// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CSGShape3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGShape3D) BaseClass() string {
  return "CSGShape3D"
}



// Enums

type CSGShape3DOperation int
const (
  CSGShape3DOperationOperationUnion CSGShape3DOperation = 0
  CSGShape3DOperationOperationIntersection CSGShape3DOperation = 1
  CSGShape3DOperationOperationSubtraction CSGShape3DOperation = 2
)

func (me *CSGShape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CSGShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CSGShape3D) IsRootShape() bool {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_root_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGShape3D) SetOperation(operation CSGShape3DOperation, )  {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_operation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 811425055) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&operation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGShape3D) GetOperation() CSGShape3DOperation {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_operation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2662425879) // FIXME: should cache?
  var ret CSGShape3DOperation
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGShape3D) SetSnap(snap float32, )  {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&snap), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGShape3D) GetSnap() float32 {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGShape3D) SetUseCollision(operation bool, )  {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&operation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGShape3D) IsUsingCollision() bool {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGShape3D) SetCollisionLayer(layer int, )  {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGShape3D) GetCollisionLayer() int {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGShape3D) SetCollisionMask(mask int, )  {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGShape3D) GetCollisionMask() int {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGShape3D) SetCollisionMaskValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGShape3D) GetCollisionMaskValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGShape3D) SetCollisionLayerValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGShape3D) GetCollisionLayerValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGShape3D) SetCollisionPriority(priority float32, )  {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGShape3D) GetCollisionPriority() float32 {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGShape3D) SetCalculateTangents(enabled bool, )  {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_calculate_tangents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGShape3D) IsCalculatingTangents() bool {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_calculating_tangents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGShape3D) GetMeshes() Array {
  classNameV := StringNameFromStr("CSGShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_meshes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *CSGShape3D) GetPropOperation() int {
  panic("TODO: implement")
}

func (me *CSGShape3D) SetPropOperation(value int) {
  panic("TODO: implement")
}

func (me *CSGShape3D) GetPropSnap() float32 {
  panic("TODO: implement")
}

func (me *CSGShape3D) SetPropSnap(value float32) {
  panic("TODO: implement")
}

func (me *CSGShape3D) GetPropCalculateTangents() bool {
  panic("TODO: implement")
}

func (me *CSGShape3D) SetPropCalculateTangents(value bool) {
  panic("TODO: implement")
}

func (me *CSGShape3D) GetPropUseCollision() bool {
  panic("TODO: implement")
}

func (me *CSGShape3D) SetPropUseCollision(value bool) {
  panic("TODO: implement")
}

func (me *CSGShape3D) GetPropCollisionLayer() int {
  panic("TODO: implement")
}

func (me *CSGShape3D) SetPropCollisionLayer(value int) {
  panic("TODO: implement")
}

func (me *CSGShape3D) GetPropCollisionMask() int {
  panic("TODO: implement")
}

func (me *CSGShape3D) SetPropCollisionMask(value int) {
  panic("TODO: implement")
}

func (me *CSGShape3D) GetPropCollisionPriority() float32 {
  panic("TODO: implement")
}

func (me *CSGShape3D) SetPropCollisionPriority(value float32) {
  panic("TODO: implement")
}