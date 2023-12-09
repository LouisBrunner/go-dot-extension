// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func  GLTFCameraFromNode(camera_node Camera3D, )  {
  panic("TODO: implement")
}

func  (me *GLTFCamera) ToNode()  {
  panic("TODO: implement")
}

func  GLTFCameraFromDictionary(dictionary Dictionary, )  {
  panic("TODO: implement")
}

func  (me *GLTFCamera) ToDictionary()  {
  panic("TODO: implement")
}

func  (me *GLTFCamera) GetPerspective()  {
  panic("TODO: implement")
}

func  (me *GLTFCamera) SetPerspective(perspective bool, )  {
  panic("TODO: implement")
}

func  (me *GLTFCamera) GetFov()  {
  panic("TODO: implement")
}

func  (me *GLTFCamera) SetFov(fov float32, )  {
  panic("TODO: implement")
}

func  (me *GLTFCamera) GetSizeMag()  {
  panic("TODO: implement")
}

func  (me *GLTFCamera) SetSizeMag(size_mag float32, )  {
  panic("TODO: implement")
}

func  (me *GLTFCamera) GetDepthFar()  {
  panic("TODO: implement")
}

func  (me *GLTFCamera) SetDepthFar(zdepth_far float32, )  {
  panic("TODO: implement")
}

func  (me *GLTFCamera) GetDepthNear()  {
  panic("TODO: implement")
}

func  (me *GLTFCamera) SetDepthNear(zdepth_near float32, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
