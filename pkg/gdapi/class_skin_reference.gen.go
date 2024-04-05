// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForSkinReferenceList struct {
  fnGetSkeleton gdc.MethodBindPtr
  fnGetSkin gdc.MethodBindPtr
}

var ptrsForSkinReference ptrsForSkinReferenceList

func initSkinReferencePtrs(iface gdc.Interface) {

  className := StringNameFromStr("SkinReference")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_skeleton")
    defer methodName.Destroy()
    ptrsForSkinReference.fnGetSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
  }
  {
    methodName := StringNameFromStr("get_skin")
    defer methodName.Destroy()
    ptrsForSkinReference.fnGetSkin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2074563878))
  }
}

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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkinReference.fnGetSkeleton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkinReference) GetSkin() Skin {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSkin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkinReference.fnGetSkin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
