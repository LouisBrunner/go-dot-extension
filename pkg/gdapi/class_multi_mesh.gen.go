// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MultiMesh struct {
  Resource
}

func (me *MultiMesh) BaseClass() string {
  return "MultiMesh"
}

func NewMultiMesh() *MultiMesh {
  str := StringNameFromStr("MultiMesh") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MultiMesh{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewMesh()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
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
  cargs := []gdc.ConstTypePtr{}
  var ret MultiMeshTransformFormat

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *MultiMesh) SetInstanceCount(count int64, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_instance_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiMesh) GetInstanceCount() int64 {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiMesh) SetVisibleInstanceCount(count int64, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visible_instance_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiMesh) GetVisibleInstanceCount() int64 {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visible_instance_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MultiMesh) SetInstanceTransform(instance int64, transform Transform3D, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_instance_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3616898986) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), gdc.ConstTypePtr(transform.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiMesh) SetInstanceTransform2D(instance int64, transform Transform2D, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_instance_transform_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 30160968) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), gdc.ConstTypePtr(transform.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiMesh) GetInstanceTransform(instance int64, ) Transform3D {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965739696) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), }
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MultiMesh) GetInstanceTransform2D(instance int64, ) Transform2D {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_transform_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3836996910) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), }
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MultiMesh) SetInstanceColor(instance int64, color Color, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_instance_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), gdc.ConstTypePtr(color.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiMesh) GetInstanceColor(instance int64, ) Color {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), }
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MultiMesh) SetInstanceCustomData(instance int64, custom_data Color, )  {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_instance_custom_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), gdc.ConstTypePtr(custom_data.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MultiMesh) GetInstanceCustomData(instance int64, ) Color {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_custom_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), }
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MultiMesh) GetAabb() AABB {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1068685055) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewAABB()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MultiMesh) GetBuffer() PackedFloat32Array {
  classNameV := StringNameFromStr("MultiMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 675695659) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPackedFloat32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
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
// FIXME: can't seem to be able to use those from this side of the API

// Signals
