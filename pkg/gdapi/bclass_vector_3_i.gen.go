// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector3i struct {
  obj gdc.ObjectPtr
}

func (me *Vector3i) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Vector3i) BaseClass() string {
  return "Vector3i"
}

// TODO: needed?
// const (
//   Vector3iAxisX = 0
//   Vector3iAxisY = 1
//   Vector3iAxisZ = 2
// // )

var (
  Vector3iZero = "Vector3i(0, 0, 0)" // TODO: construct correctly
  Vector3iOne = "Vector3i(1, 1, 1)" // TODO: construct correctly
  Vector3iLeft = "Vector3i(-1, 0, 0)" // TODO: construct correctly
  Vector3iRight = "Vector3i(1, 0, 0)" // TODO: construct correctly
  Vector3iUp = "Vector3i(0, 1, 0)" // TODO: construct correctly
  Vector3iDown = "Vector3i(0, -1, 0)" // TODO: construct correctly
  Vector3iForward = "Vector3i(0, 0, -1)" // TODO: construct correctly
  Vector3iBack = "Vector3i(0, 0, 1)" // TODO: construct correctly
)

type Vector3iAxis int
const (
  Vector3iAxisAxisX Vector3iAxis = 0
  Vector3iAxisAxisY Vector3iAxis = 1
  Vector3iAxisAxisZ Vector3iAxis = 2
)

func  (me *Vector3i) MinAxisIndex() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3i) MaxAxisIndex() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3i) Length() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3i) LengthSquared() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3i) Sign() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3i) Abs() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3i) Clamp(min Vector3i, max Vector3i, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3i) Snapped(step Vector3i, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
