// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type NavigationPathQueryResult3D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationPathQueryResult3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationPathQueryResult3D) BaseClass() string {
  return "NavigationPathQueryResult3D"
}



// Enums

type NavigationPathQueryResult3DPathSegmentType int
const (
  NavigationPathQueryResult3DPathSegmentTypePathSegmentTypeRegion NavigationPathQueryResult3DPathSegmentType = 0
  NavigationPathQueryResult3DPathSegmentTypePathSegmentTypeLink NavigationPathQueryResult3DPathSegmentType = 1
)

func (me *NavigationPathQueryResult3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationPathQueryResult3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationPathQueryResult3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NavigationPathQueryResult3D) SetPath(path PackedVector3Array, )  {
  classNameV := StringNameFromStr("NavigationPathQueryResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 334873810) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryResult3D) GetPath() PackedVector3Array {
  classNameV := StringNameFromStr("NavigationPathQueryResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 497664490) // FIXME: should cache?
  var ret PackedVector3Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPathQueryResult3D) SetPathTypes(path_types PackedInt32Array, )  {
  classNameV := StringNameFromStr("NavigationPathQueryResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_types")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path_types.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryResult3D) GetPathTypes() PackedInt32Array {
  classNameV := StringNameFromStr("NavigationPathQueryResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_types")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPathQueryResult3D) SetPathRids(path_rids RID, )  {
  classNameV := StringNameFromStr("NavigationPathQueryResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_rids")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path_rids.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryResult3D) GetPathRids() RID {
  classNameV := StringNameFromStr("NavigationPathQueryResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_rids")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPathQueryResult3D) SetPathOwnerIds(path_owner_ids PackedInt64Array, )  {
  classNameV := StringNameFromStr("NavigationPathQueryResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_owner_ids")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3709968205) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path_owner_ids.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPathQueryResult3D) GetPathOwnerIds() PackedInt64Array {
  classNameV := StringNameFromStr("NavigationPathQueryResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_owner_ids")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 235988956) // FIXME: should cache?
  var ret PackedInt64Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPathQueryResult3D) Reset()  {
  classNameV := StringNameFromStr("NavigationPathQueryResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
