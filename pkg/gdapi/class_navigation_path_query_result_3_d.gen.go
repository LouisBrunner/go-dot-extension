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

type ptrsForNavigationPathQueryResult3DList struct {
  fnSetPath gdc.MethodBindPtr
  fnGetPath gdc.MethodBindPtr
  fnSetPathTypes gdc.MethodBindPtr
  fnGetPathTypes gdc.MethodBindPtr
  fnSetPathRids gdc.MethodBindPtr
  fnGetPathRids gdc.MethodBindPtr
  fnSetPathOwnerIds gdc.MethodBindPtr
  fnGetPathOwnerIds gdc.MethodBindPtr
  fnReset gdc.MethodBindPtr
}

var ptrsForNavigationPathQueryResult3D ptrsForNavigationPathQueryResult3DList

func initNavigationPathQueryResult3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("NavigationPathQueryResult3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_path")
    defer methodName.Destroy()
    ptrsForNavigationPathQueryResult3D.fnSetPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 334873810))
  }
  {
    methodName := StringNameFromStr("get_path")
    defer methodName.Destroy()
    ptrsForNavigationPathQueryResult3D.fnGetPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 497664490))
  }
  {
    methodName := StringNameFromStr("set_path_types")
    defer methodName.Destroy()
    ptrsForNavigationPathQueryResult3D.fnSetPathTypes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
  }
  {
    methodName := StringNameFromStr("get_path_types")
    defer methodName.Destroy()
    ptrsForNavigationPathQueryResult3D.fnGetPathTypes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
  }
  {
    methodName := StringNameFromStr("set_path_rids")
    defer methodName.Destroy()
    ptrsForNavigationPathQueryResult3D.fnSetPathRids = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_path_rids")
    defer methodName.Destroy()
    ptrsForNavigationPathQueryResult3D.fnGetPathRids = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("set_path_owner_ids")
    defer methodName.Destroy()
    ptrsForNavigationPathQueryResult3D.fnSetPathOwnerIds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3709968205))
  }
  {
    methodName := StringNameFromStr("get_path_owner_ids")
    defer methodName.Destroy()
    ptrsForNavigationPathQueryResult3D.fnGetPathOwnerIds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 235988956))
  }
  {
    methodName := StringNameFromStr("reset")
    defer methodName.Destroy()
    ptrsForNavigationPathQueryResult3D.fnReset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
}

type NavigationPathQueryResult3D struct {
  RefCounted
}

func (me *NavigationPathQueryResult3D) BaseClass() string {
  return "NavigationPathQueryResult3D"
}

func NewNavigationPathQueryResult3D() *NavigationPathQueryResult3D {
  str := StringNameFromStr("NavigationPathQueryResult3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &NavigationPathQueryResult3D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult3D.fnSetPath), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPathQueryResult3D) GetPath() PackedVector3Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult3D.fnGetPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPathQueryResult3D) SetPathTypes(path_types PackedInt32Array, )  {
  cargs := []gdc.ConstTypePtr{path_types.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult3D.fnSetPathTypes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPathQueryResult3D) GetPathTypes() PackedInt32Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult3D.fnGetPathTypes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPathQueryResult3D) SetPathRids(path_rids []RID, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path_rids) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult3D.fnSetPathRids), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPathQueryResult3D) GetPathRids() []RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult3D.fnGetPathRids), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[RID](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *NavigationPathQueryResult3D) SetPathOwnerIds(path_owner_ids PackedInt64Array, )  {
  cargs := []gdc.ConstTypePtr{path_owner_ids.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult3D.fnSetPathOwnerIds), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPathQueryResult3D) GetPathOwnerIds() PackedInt64Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt64Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult3D.fnGetPathOwnerIds), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPathQueryResult3D) Reset()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult3D.fnReset), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
