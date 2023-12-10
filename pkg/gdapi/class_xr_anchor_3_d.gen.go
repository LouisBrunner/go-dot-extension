// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type XRAnchor3D struct {
  obj gdc.ObjectPtr
}

func (me *XRAnchor3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XRAnchor3D) BaseClass() string {
  return "XRAnchor3D"
}



// Enums

func (me *XRAnchor3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *XRAnchor3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XRAnchor3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *XRAnchor3D) GetSize() Vector3 {
  classNameV := StringNameFromStr("XRAnchor3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRAnchor3D) GetPlane() Plane {
  classNameV := StringNameFromStr("XRAnchor3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_plane")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2753500971) // FIXME: should cache?
  var ret Plane
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties