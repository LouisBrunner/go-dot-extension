// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector2i struct {
  obj gdc.ObjectPtr
}

func (me *Vector2i) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Vector2i) BaseClass() string {
  return "Vector2i"
}

// TODO: needed?
// const (
//   Vector2iAxisX = 0
//   Vector2iAxisY = 1
// // )

var (
  Vector2iZero = "Vector2i(0, 0)" // TODO: construct correctly
  Vector2iOne = "Vector2i(1, 1)" // TODO: construct correctly
  Vector2iLeft = "Vector2i(-1, 0)" // TODO: construct correctly
  Vector2iRight = "Vector2i(1, 0)" // TODO: construct correctly
  Vector2iUp = "Vector2i(0, -1)" // TODO: construct correctly
  Vector2iDown = "Vector2i(0, 1)" // TODO: construct correctly
)

type Vector2iAxis int
const (
  Vector2iAxisAxisX Vector2iAxis = 0
  Vector2iAxisAxisY Vector2iAxis = 1
)

func  (me *Vector2i) Aspect() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2i) MaxAxisIndex() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2i) MinAxisIndex() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2i) Length() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2i) LengthSquared() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2i) Sign() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2i) Abs() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2i) Clamp(min Vector2i, max Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2i) Snapped(step Vector2i, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
