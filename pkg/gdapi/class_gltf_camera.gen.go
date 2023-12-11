// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GLTFCamera struct {
  obj gdc.ObjectPtr
}

func (me *GLTFCamera) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFCamera) BaseClass() string {
  return "GLTFCamera"
}



// Enums

func (me *GLTFCamera) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFCamera) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFCamera) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  GLTFCameraFromNode(camera_node Camera3D, ) GLTFCamera {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("from_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 237784) // FIXME: should cache?
  var ret GLTFCamera
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(camera_node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFCamera) ToNode() Camera3D {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2285090890) // FIXME: should cache?
  var ret Camera3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  GLTFCameraFromDictionary(dictionary Dictionary, ) GLTFCamera {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("from_dictionary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2495512509) // FIXME: should cache?
  var ret GLTFCamera
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(dictionary.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFCamera) ToDictionary() Dictionary {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_dictionary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFCamera) GetPerspective() bool {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_perspective")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFCamera) SetPerspective(perspective bool, )  {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_perspective")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&perspective), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFCamera) GetFov() float32 {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fov")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFCamera) SetFov(fov float32, )  {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fov")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fov), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFCamera) GetSizeMag() float32 {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size_mag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFCamera) SetSizeMag(size_mag float32, )  {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size_mag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size_mag), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFCamera) GetDepthFar() float32 {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_far")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFCamera) SetDepthFar(zdepth_far float32, )  {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth_far")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zdepth_far), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFCamera) GetDepthNear() float32 {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_near")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFCamera) SetDepthNear(zdepth_near float32, )  {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth_near")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zdepth_near), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
