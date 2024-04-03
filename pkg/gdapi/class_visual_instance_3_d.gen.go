// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualInstance3D struct {
  Node3D
}

func (me *VisualInstance3D) BaseClass() string {
  return "VisualInstance3D"
}

func NewVisualInstance3D() *VisualInstance3D {
  str := StringNameFromStr("VisualInstance3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualInstance3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualInstance3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualInstance3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualInstance3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualInstance3D) SetBase(base RID, )  {
  classNameV := StringNameFromStr("VisualInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_base")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(base.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualInstance3D) GetBase() RID {
  classNameV := StringNameFromStr("VisualInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_base")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VisualInstance3D) GetInstance() RID {
  classNameV := StringNameFromStr("VisualInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VisualInstance3D) SetLayerMask(mask int64, )  {
  classNameV := StringNameFromStr("VisualInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualInstance3D) GetLayerMask() int64 {
  classNameV := StringNameFromStr("VisualInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualInstance3D) SetLayerMaskValue(layer_number int64, value bool, )  {
  classNameV := StringNameFromStr("VisualInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualInstance3D) GetLayerMaskValue(layer_number int64, ) bool {
  classNameV := StringNameFromStr("VisualInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualInstance3D) SetSortingOffset(offset float64, )  {
  classNameV := StringNameFromStr("VisualInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sorting_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualInstance3D) GetSortingOffset() float64 {
  classNameV := StringNameFromStr("VisualInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sorting_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualInstance3D) SetSortingUseAabbCenter(enabled bool, )  {
  classNameV := StringNameFromStr("VisualInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sorting_use_aabb_center")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualInstance3D) IsSortingUseAabbCenter() bool {
  classNameV := StringNameFromStr("VisualInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_sorting_use_aabb_center")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualInstance3D) GetAabb() AABB {
  classNameV := StringNameFromStr("VisualInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1068685055) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewAABB()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
