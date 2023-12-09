// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *NavigationPathQueryResult3D) GetPath()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryResult3D) SetPathTypes(path_types PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryResult3D) GetPathTypes()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryResult3D) SetPathRids(path_rids RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryResult3D) GetPathRids()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryResult3D) SetPathOwnerIds(path_owner_ids PackedInt64Array, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryResult3D) GetPathOwnerIds()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryResult3D) Reset()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
