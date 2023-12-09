// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGPolygon3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGPolygon3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGPolygon3D) BaseClass() string {
  return "CSGPolygon3D"
}



// Enums

type CSGPolygon3DMode int
const (
  CSGPolygon3DModeModeDepth CSGPolygon3DMode = 0
  CSGPolygon3DModeModeSpin CSGPolygon3DMode = 1
  CSGPolygon3DModeModePath CSGPolygon3DMode = 2
)

type CSGPolygon3DPathRotation int
const (
  CSGPolygon3DPathRotationPathRotationPolygon CSGPolygon3DPathRotation = 0
  CSGPolygon3DPathRotationPathRotationPath CSGPolygon3DPathRotation = 1
  CSGPolygon3DPathRotationPathRotationPathFollow CSGPolygon3DPathRotation = 2
)

type CSGPolygon3DPathIntervalType int
const (
  CSGPolygon3DPathIntervalTypePathIntervalDistance CSGPolygon3DPathIntervalType = 0
  CSGPolygon3DPathIntervalTypePathIntervalSubdivide CSGPolygon3DPathIntervalType = 1
)

func (me *CSGPolygon3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGPolygon3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CSGPolygon3D) SetPolygon(polygon PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) GetPolygon()  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) SetMode(mode CSGPolygon3DMode, )  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) GetMode()  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) SetDepth(depth float32, )  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) GetDepth()  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) SetSpinDegrees(degrees float32, )  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) GetSpinDegrees()  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) SetSpinSides(spin_sides int, )  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) GetSpinSides()  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) SetPathNode(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) GetPathNode()  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) SetPathIntervalType(interval_type CSGPolygon3DPathIntervalType, )  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) GetPathIntervalType()  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) SetPathInterval(interval float32, )  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) GetPathInterval()  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) SetPathSimplifyAngle(degrees float32, )  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) GetPathSimplifyAngle()  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) SetPathRotation(path_rotation CSGPolygon3DPathRotation, )  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) GetPathRotation()  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) SetPathLocal(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) IsPathLocal()  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) SetPathContinuousU(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) IsPathContinuousU()  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) SetPathUDistance(distance float32, )  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) GetPathUDistance()  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) SetPathJoined(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) IsPathJoined()  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) SetMaterial(material Material, )  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) GetMaterial()  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) SetSmoothFaces(smooth_faces bool, )  {
  panic("TODO: implement")
}

func  (me *CSGPolygon3D) GetSmoothFaces()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
