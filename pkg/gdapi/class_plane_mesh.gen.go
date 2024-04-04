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

type PlaneMesh struct {
  PrimitiveMesh
}

func (me *PlaneMesh) BaseClass() string {
  return "PlaneMesh"
}

func NewPlaneMesh() *PlaneMesh {
  str := StringNameFromStr("PlaneMesh") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PlaneMesh{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type PlaneMeshOrientation int
const (
  PlaneMeshOrientationFaceX PlaneMeshOrientation = 0
  PlaneMeshOrientationFaceY PlaneMeshOrientation = 1
  PlaneMeshOrientationFaceZ PlaneMeshOrientation = 2
)

func (me *PlaneMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PlaneMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PlaneMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PlaneMesh) SetSize(size Vector2, )  {
  classNameV := StringNameFromStr("PlaneMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PlaneMesh) GetSize() Vector2 {
  classNameV := StringNameFromStr("PlaneMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PlaneMesh) SetSubdivideWidth(subdivide int64, )  {
  classNameV := StringNameFromStr("PlaneMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_subdivide_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subdivide) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PlaneMesh) GetSubdivideWidth() int64 {
  classNameV := StringNameFromStr("PlaneMesh")
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

func  (me *PlaneMesh) SetSubdivideDepth(subdivide int64, )  {
  classNameV := StringNameFromStr("PlaneMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_subdivide_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subdivide) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PlaneMesh) GetSubdivideDepth() int64 {
  classNameV := StringNameFromStr("PlaneMesh")
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

func  (me *PlaneMesh) SetCenterOffset(offset Vector3, )  {
  classNameV := StringNameFromStr("PlaneMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_center_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PlaneMesh) GetCenterOffset() Vector3 {
  classNameV := StringNameFromStr("PlaneMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_center_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PlaneMesh) SetOrientation(orientation PlaneMeshOrientation, )  {
  classNameV := StringNameFromStr("PlaneMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_orientation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2751399687) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&orientation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PlaneMesh) GetOrientation() PlaneMeshOrientation {
  classNameV := StringNameFromStr("PlaneMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_orientation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227599250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret PlaneMeshOrientation

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
