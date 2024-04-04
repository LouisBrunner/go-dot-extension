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

type NavigationPathQueryResult2D struct {
  RefCounted
}

func (me *NavigationPathQueryResult2D) BaseClass() string {
  return "NavigationPathQueryResult2D"
}

func NewNavigationPathQueryResult2D() *NavigationPathQueryResult2D {
  str := StringNameFromStr("NavigationPathQueryResult2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &NavigationPathQueryResult2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type NavigationPathQueryResult2DPathSegmentType int
const (
  NavigationPathQueryResult2DPathSegmentTypePathSegmentTypeRegion NavigationPathQueryResult2DPathSegmentType = 0
  NavigationPathQueryResult2DPathSegmentTypePathSegmentTypeLink NavigationPathQueryResult2DPathSegmentType = 1
)

func (me *NavigationPathQueryResult2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationPathQueryResult2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationPathQueryResult2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NavigationPathQueryResult2D) SetPath(path PackedVector2Array, )  {
  classNameV := StringNameFromStr("NavigationPathQueryResult2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPathQueryResult2D) GetPath() PackedVector2Array {
  classNameV := StringNameFromStr("NavigationPathQueryResult2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPathQueryResult2D) SetPathTypes(path_types PackedInt32Array, )  {
  classNameV := StringNameFromStr("NavigationPathQueryResult2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_types")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path_types.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPathQueryResult2D) GetPathTypes() PackedInt32Array {
  classNameV := StringNameFromStr("NavigationPathQueryResult2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_types")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPathQueryResult2D) SetPathRids(path_rids []RID, )  {
  classNameV := StringNameFromStr("NavigationPathQueryResult2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_rids")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path_rids) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPathQueryResult2D) GetPathRids() []RID {
  classNameV := StringNameFromStr("NavigationPathQueryResult2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_rids")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[RID](ret)
}

func  (me *NavigationPathQueryResult2D) SetPathOwnerIds(path_owner_ids PackedInt64Array, )  {
  classNameV := StringNameFromStr("NavigationPathQueryResult2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_owner_ids")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3709968205) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path_owner_ids.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPathQueryResult2D) GetPathOwnerIds() PackedInt64Array {
  classNameV := StringNameFromStr("NavigationPathQueryResult2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_owner_ids")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 235988956) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt64Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPathQueryResult2D) Reset()  {
  classNameV := StringNameFromStr("NavigationPathQueryResult2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
