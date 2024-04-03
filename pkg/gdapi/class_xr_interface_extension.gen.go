// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type XRInterfaceExtension struct {
  XRInterface
}

func (me *XRInterfaceExtension) BaseClass() string {
  return "XRInterfaceExtension"
}

func NewXRInterfaceExtension() *XRInterfaceExtension {
  str := StringNameFromStr("XRInterfaceExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &XRInterfaceExtension{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *XRInterfaceExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *XRInterfaceExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XRInterfaceExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *XRInterfaceExtension) GetColorTexture() RID {
  classNameV := StringNameFromStr("XRInterfaceExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRInterfaceExtension) GetDepthTexture() RID {
  classNameV := StringNameFromStr("XRInterfaceExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRInterfaceExtension) GetVelocityTexture() RID {
  classNameV := StringNameFromStr("XRInterfaceExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *XRInterfaceExtension) AddBlit(render_target RID, src_rect Rect2, dst_rect Rect2i, use_layer bool, layer int64, apply_lens_distortion bool, eye_center Vector2, k1 float64, k2 float64, upscale float64, aspect_ratio float64, )  {
  classNameV := StringNameFromStr("XRInterfaceExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_blit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 258596971) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(render_target.AsCTypePtr()), gdc.ConstTypePtr(src_rect.AsCTypePtr()), gdc.ConstTypePtr(dst_rect.AsCTypePtr()), gdc.ConstTypePtr(&use_layer), gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&apply_lens_distortion), gdc.ConstTypePtr(eye_center.AsCTypePtr()), gdc.ConstTypePtr(&k1), gdc.ConstTypePtr(&k2), gdc.ConstTypePtr(&upscale), gdc.ConstTypePtr(&aspect_ratio), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *XRInterfaceExtension) GetRenderTargetTexture(render_target RID, ) RID {
  classNameV := StringNameFromStr("XRInterfaceExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_render_target_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 41030802) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(render_target.AsCTypePtr()), }
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
