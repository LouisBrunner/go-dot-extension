// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type RandomNumberGenerator struct {
  RefCounted
}

func (me *RandomNumberGenerator) BaseClass() string {
  return "RandomNumberGenerator"
}

func NewRandomNumberGenerator() *RandomNumberGenerator {
  str := StringNameFromStr("RandomNumberGenerator") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RandomNumberGenerator{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *RandomNumberGenerator) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RandomNumberGenerator) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RandomNumberGenerator) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RandomNumberGenerator) SetSeed(seed int64, )  {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_seed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RandomNumberGenerator) GetSeed() int64 {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_seed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RandomNumberGenerator) SetState(state int64, )  {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&state) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RandomNumberGenerator) GetState() int64 {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RandomNumberGenerator) Randi() int64 {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("randi")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RandomNumberGenerator) Randf() float64 {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("randf")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RandomNumberGenerator) Randfn(mean float64, deviation float64, ) float64 {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("randfn")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 837325100) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mean) , gdc.ConstTypePtr(&deviation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&mean)
  pinner.Pin(&deviation)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RandomNumberGenerator) RandfRange(from float64, to float64, ) float64 {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("randf_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4269894367) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from) , gdc.ConstTypePtr(&to) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&from)
  pinner.Pin(&to)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RandomNumberGenerator) RandiRange(from int64, to int64, ) int64 {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("randi_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 50157827) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from) , gdc.ConstTypePtr(&to) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&from)
  pinner.Pin(&to)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RandomNumberGenerator) Randomize()  {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("randomize")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
