// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RandomNumberGenerator struct {
  obj gdc.ObjectPtr
}

func (me *RandomNumberGenerator) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RandomNumberGenerator) BaseClass() string {
  return "RandomNumberGenerator"
}



// Enums

func (me *RandomNumberGenerator) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RandomNumberGenerator) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *RandomNumberGenerator) SetSeed(seed int, )  {
  panic("TODO: implement")
}

func  (me *RandomNumberGenerator) GetSeed()  {
  panic("TODO: implement")
}

func  (me *RandomNumberGenerator) SetState(state int, )  {
  panic("TODO: implement")
}

func  (me *RandomNumberGenerator) GetState()  {
  panic("TODO: implement")
}

func  (me *RandomNumberGenerator) Randi()  {
  panic("TODO: implement")
}

func  (me *RandomNumberGenerator) Randf()  {
  panic("TODO: implement")
}

func  (me *RandomNumberGenerator) Randfn(mean float32, deviation float32, )  {
  panic("TODO: implement")
}

func  (me *RandomNumberGenerator) RandfRange(from float32, to float32, )  {
  panic("TODO: implement")
}

func  (me *RandomNumberGenerator) RandiRange(from int, to int, )  {
  panic("TODO: implement")
}

func  (me *RandomNumberGenerator) Randomize()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
