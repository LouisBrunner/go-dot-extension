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

type BoxMesh struct {
  PrimitiveMesh
}

func (me *BoxMesh) BaseClass() string {
  return "BoxMesh"
}

func NewBoxMesh() *BoxMesh {
  str := StringNameFromStr("BoxMesh") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &BoxMesh{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *BoxMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *BoxMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *BoxMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *BoxMesh) SetSize(size Vector3, )  {
  classNameV := StringNameFromStr("BoxMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BoxMesh) GetSize() Vector3 {
  classNameV := StringNameFromStr("BoxMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BoxMesh) SetSubdivideWidth(subdivide int64, )  {
  classNameV := StringNameFromStr("BoxMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_subdivide_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subdivide) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BoxMesh) GetSubdivideWidth() int64 {
  classNameV := StringNameFromStr("BoxMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_subdivide_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BoxMesh) SetSubdivideHeight(divisions int64, )  {
  classNameV := StringNameFromStr("BoxMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_subdivide_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&divisions) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BoxMesh) GetSubdivideHeight() int64 {
  classNameV := StringNameFromStr("BoxMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_subdivide_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BoxMesh) SetSubdivideDepth(divisions int64, )  {
  classNameV := StringNameFromStr("BoxMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_subdivide_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&divisions) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BoxMesh) GetSubdivideDepth() int64 {
  classNameV := StringNameFromStr("BoxMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_subdivide_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
