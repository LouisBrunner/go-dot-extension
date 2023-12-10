// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MultiMesh struct {
  obj gdc.ObjectPtr
}

func (me *MultiMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiMesh) BaseClass() string {
  return "MultiMesh"
}



// Enums

type MultiMeshTransformFormat int
const (
  MultiMeshTransformFormatTransform2D MultiMeshTransformFormat = 0
  MultiMeshTransformFormatTransform3D MultiMeshTransformFormat = 1
)

func (me *MultiMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MultiMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MultiMesh) SetMesh(mesh Mesh, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 194775623) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mesh.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiMesh) GetMesh() Mesh {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1808005922) // FIXME: should cache?
  var ret Mesh
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiMesh) SetUseColors(enable bool, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiMesh) IsUsingColors() bool {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiMesh) SetUseCustomData(enable bool, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_custom_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiMesh) IsUsingCustomData() bool {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_custom_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiMesh) SetTransformFormat(format MultiMeshTransformFormat, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transform_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2404750322) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiMesh) GetTransformFormat() MultiMeshTransformFormat {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transform_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2444156481) // FIXME: should cache?
  var ret MultiMeshTransformFormat
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiMesh) SetInstanceCount(count int, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_instance_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiMesh) GetInstanceCount() int {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiMesh) SetVisibleInstanceCount(count int, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visible_instance_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiMesh) GetVisibleInstanceCount() int {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visible_instance_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiMesh) SetInstanceTransform(instance int, transform Transform3D, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_instance_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3616898986) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiMesh) SetInstanceTransform2D(instance int, transform Transform2D, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_instance_transform_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 30160968) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiMesh) GetInstanceTransform(instance int, ) Transform3D {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965739696) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiMesh) GetInstanceTransform2D(instance int, ) Transform2D {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_transform_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3836996910) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiMesh) SetInstanceColor(instance int, color Color, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_instance_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiMesh) GetInstanceColor(instance int, ) Color {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiMesh) SetInstanceCustomData(instance int, custom_data Color, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_instance_custom_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), gdc.ConstTypePtr(custom_data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiMesh) GetInstanceCustomData(instance int, ) Color {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_custom_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiMesh) GetAabb() AABB {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1068685055) // FIXME: should cache?
  var ret AABB
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiMesh) GetBuffer() PackedFloat32Array {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 675695659) // FIXME: should cache?
  var ret PackedFloat32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MultiMesh) SetBuffer(buffer PackedFloat32Array, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2899603908) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(buffer.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *MultiMesh) GetPropTransformFormat() int {
  panic("TODO: implement")
}

func (me *MultiMesh) SetPropTransformFormat(value int) {
  panic("TODO: implement")
}

func (me *MultiMesh) GetPropUseColors() bool {
  panic("TODO: implement")
}

func (me *MultiMesh) SetPropUseColors(value bool) {
  panic("TODO: implement")
}

func (me *MultiMesh) GetPropUseCustomData() bool {
  panic("TODO: implement")
}

func (me *MultiMesh) SetPropUseCustomData(value bool) {
  panic("TODO: implement")
}

func (me *MultiMesh) GetPropInstanceCount() int {
  panic("TODO: implement")
}

func (me *MultiMesh) SetPropInstanceCount(value int) {
  panic("TODO: implement")
}

func (me *MultiMesh) GetPropVisibleInstanceCount() int {
  panic("TODO: implement")
}

func (me *MultiMesh) SetPropVisibleInstanceCount(value int) {
  panic("TODO: implement")
}

func (me *MultiMesh) GetPropMesh() Mesh {
  panic("TODO: implement")
}

func (me *MultiMesh) SetPropMesh(value Mesh) {
  panic("TODO: implement")
}

func (me *MultiMesh) GetPropBuffer() PackedFloat32Array {
  panic("TODO: implement")
}

func (me *MultiMesh) SetPropBuffer(value PackedFloat32Array) {
  panic("TODO: implement")
}

func (me *MultiMesh) GetPropTransformArray() PackedVector3Array {
  panic("TODO: implement")
}

func (me *MultiMesh) SetPropTransformArray(value PackedVector3Array) {
  panic("TODO: implement")
}

func (me *MultiMesh) GetPropTransform2DArray() PackedVector2Array {
  panic("TODO: implement")
}

func (me *MultiMesh) SetPropTransform2DArray(value PackedVector2Array) {
  panic("TODO: implement")
}

func (me *MultiMesh) GetPropColorArray() PackedColorArray {
  panic("TODO: implement")
}

func (me *MultiMesh) SetPropColorArray(value PackedColorArray) {
  panic("TODO: implement")
}

func (me *MultiMesh) GetPropCustomDataArray() PackedColorArray {
  panic("TODO: implement")
}

func (me *MultiMesh) SetPropCustomDataArray(value PackedColorArray) {
  panic("TODO: implement")
}