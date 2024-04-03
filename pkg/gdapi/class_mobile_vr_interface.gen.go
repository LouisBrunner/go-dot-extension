// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MobileVRInterface struct {
  XRInterface
}

func (me *MobileVRInterface) BaseClass() string {
  return "MobileVRInterface"
}

func NewMobileVRInterface() *MobileVRInterface {
  str := StringNameFromStr("MobileVRInterface") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MobileVRInterface{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *MobileVRInterface) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MobileVRInterface) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MobileVRInterface) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MobileVRInterface) SetEyeHeight(eye_height float64, )  {
  classNameV := StringNameFromStr("MobileVRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_eye_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&eye_height), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MobileVRInterface) GetEyeHeight() float64 {
  classNameV := StringNameFromStr("MobileVRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_eye_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MobileVRInterface) SetIod(iod float64, )  {
  classNameV := StringNameFromStr("MobileVRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_iod")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&iod), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MobileVRInterface) GetIod() float64 {
  classNameV := StringNameFromStr("MobileVRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_iod")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MobileVRInterface) SetDisplayWidth(display_width float64, )  {
  classNameV := StringNameFromStr("MobileVRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_display_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&display_width), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MobileVRInterface) GetDisplayWidth() float64 {
  classNameV := StringNameFromStr("MobileVRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_display_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MobileVRInterface) SetDisplayToLens(display_to_lens float64, )  {
  classNameV := StringNameFromStr("MobileVRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_display_to_lens")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&display_to_lens), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MobileVRInterface) GetDisplayToLens() float64 {
  classNameV := StringNameFromStr("MobileVRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_display_to_lens")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MobileVRInterface) SetOversample(oversample float64, )  {
  classNameV := StringNameFromStr("MobileVRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_oversample")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&oversample), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MobileVRInterface) GetOversample() float64 {
  classNameV := StringNameFromStr("MobileVRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_oversample")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MobileVRInterface) SetK1(k float64, )  {
  classNameV := StringNameFromStr("MobileVRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_k1")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&k), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MobileVRInterface) GetK1() float64 {
  classNameV := StringNameFromStr("MobileVRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_k1")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MobileVRInterface) SetK2(k float64, )  {
  classNameV := StringNameFromStr("MobileVRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_k2")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&k), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MobileVRInterface) GetK2() float64 {
  classNameV := StringNameFromStr("MobileVRInterface")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_k2")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
