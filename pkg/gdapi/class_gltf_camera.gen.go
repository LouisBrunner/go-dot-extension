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

type GLTFCamera struct {
  Resource
}

func (me *GLTFCamera) BaseClass() string {
  return "GLTFCamera"
}

func NewGLTFCamera() *GLTFCamera {
  str := StringNameFromStr("GLTFCamera") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GLTFCamera{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{camera_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewGLTFCamera()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFCamera) ToNode() Camera3D {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2285090890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCamera3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  GLTFCameraFromDictionary(dictionary Dictionary, ) GLTFCamera {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("from_dictionary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2495512509) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{dictionary.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewGLTFCamera()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFCamera) ToDictionary() Dictionary {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_dictionary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFCamera) GetPerspective() bool {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_perspective")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFCamera) SetPerspective(perspective bool, )  {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_perspective")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&perspective) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFCamera) GetFov() float64 {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fov")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFCamera) SetFov(fov float64, )  {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fov")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fov) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFCamera) GetSizeMag() float64 {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size_mag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFCamera) SetSizeMag(size_mag float64, )  {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size_mag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size_mag) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFCamera) GetDepthFar() float64 {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_far")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFCamera) SetDepthFar(zdepth_far float64, )  {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth_far")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zdepth_far) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFCamera) GetDepthNear() float64 {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_near")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFCamera) SetDepthNear(zdepth_near float64, )  {
  classNameV := StringNameFromStr("GLTFCamera")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth_near")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zdepth_near) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
