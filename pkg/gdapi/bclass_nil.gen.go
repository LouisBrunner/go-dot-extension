// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Nil struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewNil() Nil {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeNil))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeNil, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Nil{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewNilFromVariant(from Variant, ) Nil {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeNil))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeNil, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Nil{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Nil) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Nil) Type() gdc.VariantType {
  return gdc.VariantTypeNil
}

func (me *Nil) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Nil) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

// Operators

func (me *Nil) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Nil) OrVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Nil) Not() bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Nil) AndBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Nil) OrBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Nil) XorBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsInt(right int) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsInt(right int) bool {
  panic("TODO: implement")
}

func (me *Nil) AndInt(right int) bool {
  panic("TODO: implement")
}

func (me *Nil) OrInt(right int) bool {
  panic("TODO: implement")
}

func (me *Nil) XorInt(right int) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Nil) AndFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Nil) OrFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Nil) XorFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsString(right String) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsString(right String) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsVector2(right Vector2) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsVector2(right Vector2) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsVector2I(right Vector2i) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsVector2I(right Vector2i) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsRect2(right Rect2) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsRect2(right Rect2) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsRect2I(right Rect2i) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsRect2I(right Rect2i) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsVector3(right Vector3) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsVector3(right Vector3) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsVector3I(right Vector3i) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsVector3I(right Vector3i) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsTransform2D(right Transform2D) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsTransform2D(right Transform2D) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsVector4(right Vector4) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsVector4(right Vector4) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsVector4I(right Vector4i) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsVector4I(right Vector4i) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsPlane(right Plane) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsPlane(right Plane) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsQuaternion(right Quaternion) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsQuaternion(right Quaternion) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsAABB(right AABB) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsAABB(right AABB) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsBasis(right Basis) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsBasis(right Basis) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsTransform3D(right Transform3D) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsTransform3D(right Transform3D) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsProjection(right Projection) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsProjection(right Projection) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsColor(right Color) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsColor(right Color) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsStringName(right StringName) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsStringName(right StringName) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsNodePath(right NodePath) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsNodePath(right NodePath) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsRID(right RID) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsRID(right RID) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsObject(right Object) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsObject(right Object) bool {
  panic("TODO: implement")
}

func (me *Nil) AndObject(right Object) bool {
  panic("TODO: implement")
}

func (me *Nil) OrObject(right Object) bool {
  panic("TODO: implement")
}

func (me *Nil) XorObject(right Object) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsCallable(right Callable) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsCallable(right Callable) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsSignal(right Signal) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsSignal(right Signal) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Nil) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsArray(right Array) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsArray(right Array) bool {
  panic("TODO: implement")
}

func (me *Nil) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsPackedByteArray(right PackedByteArray) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsPackedByteArray(right PackedByteArray) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsPackedInt32Array(right PackedInt32Array) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsPackedInt32Array(right PackedInt32Array) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsPackedInt64Array(right PackedInt64Array) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsPackedInt64Array(right PackedInt64Array) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsPackedFloat32Array(right PackedFloat32Array) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsPackedFloat32Array(right PackedFloat32Array) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsPackedFloat64Array(right PackedFloat64Array) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsPackedFloat64Array(right PackedFloat64Array) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsPackedStringArray(right PackedStringArray) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsPackedStringArray(right PackedStringArray) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsPackedVector2Array(right PackedVector2Array) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsPackedVector2Array(right PackedVector2Array) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsPackedVector3Array(right PackedVector3Array) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsPackedVector3Array(right PackedVector3Array) bool {
  panic("TODO: implement")
}

func (me *Nil) EqualsPackedColorArray(right PackedColorArray) bool {
  panic("TODO: implement")
}

func (me *Nil) NotEqualsPackedColorArray(right PackedColorArray) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
