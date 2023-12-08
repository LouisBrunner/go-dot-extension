// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Plane struct {
  obj gdc.ObjectPtr
}

func (me *Plane) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Plane) BaseClass() string {
  return "Plane"
}

// TODO: needed?
// const (
// // )

var (
  PlanePlaneYz = "Plane(1, 0, 0, 0)" // TODO: construct correctly
  PlanePlaneXz = "Plane(0, 1, 0, 0)" // TODO: construct correctly
  PlanePlaneXy = "Plane(0, 0, 1, 0)" // TODO: construct correctly
)

func  (me *Plane) Normalized() { // TODO: return value
  // TODO: implement
}

func  (me *Plane) GetCenter() { // TODO: return value
  // TODO: implement
}

func  (me *Plane) IsEqualApprox(to_plane Plane, ) { // TODO: return value
  // TODO: implement
}

func  (me *Plane) IsFinite() { // TODO: return value
  // TODO: implement
}

func  (me *Plane) IsPointOver(point Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Plane) DistanceTo(point Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Plane) HasPoint(point Vector3, tolerance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Plane) Project(point Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Plane) Intersect3(b Plane, c Plane, ) { // TODO: return value
  // TODO: implement
}

func  (me *Plane) IntersectsRay(from Vector3, dir Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Plane) IntersectsSegment(from Vector3, to Vector3, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
