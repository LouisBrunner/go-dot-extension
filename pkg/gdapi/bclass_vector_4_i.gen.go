// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector4i struct {
  obj gdc.ObjectPtr
}

func (me *Vector4i) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Vector4i) BaseClass() string {
  return "Vector4i"
}

// TODO: needed?
// const (
//   Vector4iAxisX = 0
//   Vector4iAxisY = 1
//   Vector4iAxisZ = 2
//   Vector4iAxisW = 3
// // )

var (
  Vector4iZero = "Vector4i(0, 0, 0, 0)" // TODO: construct correctly
  Vector4iOne = "Vector4i(1, 1, 1, 1)" // TODO: construct correctly
)

type Vector4iAxis int
const (
  Vector4iAxisAxisX Vector4iAxis = 0
  Vector4iAxisAxisY Vector4iAxis = 1
  Vector4iAxisAxisZ Vector4iAxis = 2
  Vector4iAxisAxisW Vector4iAxis = 3
)

func  (me *Vector4i) MinAxisIndex() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4i) MaxAxisIndex() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4i) Length() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4i) LengthSquared() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4i) Sign() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4i) Abs() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4i) Clamp(min Vector4i, max Vector4i, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector4i) Snapped(step Vector4i, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
