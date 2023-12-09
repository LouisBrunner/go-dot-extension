// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationPathQueryResult2D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationPathQueryResult2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationPathQueryResult2D) BaseClass() string {
  return "NavigationPathQueryResult2D"
}



// Enums

type NavigationPathQueryResult2DPathSegmentType int
const (
  NavigationPathQueryResult2DPathSegmentTypePathSegmentTypeRegion NavigationPathQueryResult2DPathSegmentType = 0
  NavigationPathQueryResult2DPathSegmentTypePathSegmentTypeLink NavigationPathQueryResult2DPathSegmentType = 1
)

func (me *NavigationPathQueryResult2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationPathQueryResult2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *NavigationPathQueryResult2D) SetPath(path PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryResult2D) GetPath()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryResult2D) SetPathTypes(path_types PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryResult2D) GetPathTypes()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryResult2D) SetPathRids(path_rids RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryResult2D) GetPathRids()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryResult2D) SetPathOwnerIds(path_owner_ids PackedInt64Array, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryResult2D) GetPathOwnerIds()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryResult2D) Reset()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
