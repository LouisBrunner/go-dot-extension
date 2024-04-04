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

type SkinReference struct {
  RefCounted
}

func (me *SkinReference) BaseClass() string {
  return "SkinReference"
}

func NewSkinReference() *SkinReference {
  str := StringNameFromStr("SkinReference") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SkinReference{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *SkinReference) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SkinReference) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SkinReference) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SkinReference) GetSkeleton() RID {
  classNameV := StringNameFromStr("SkinReference")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkinReference) GetSkin() Skin {
  classNameV := StringNameFromStr("SkinReference")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2074563878) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSkin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
