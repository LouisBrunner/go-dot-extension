// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *RandomNumberGenerator) SetSeed(seed int, )  {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_seed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RandomNumberGenerator) GetSeed() int {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_seed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RandomNumberGenerator) SetState(state int, )  {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&state), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RandomNumberGenerator) GetState() int {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RandomNumberGenerator) Randi() int {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("randi")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RandomNumberGenerator) Randf() float32 {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("randf")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RandomNumberGenerator) Randfn(mean float32, deviation float32, ) float32 {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("randfn")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 837325100) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mean), gdc.ConstTypePtr(&deviation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RandomNumberGenerator) RandfRange(from float32, to float32, ) float32 {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("randf_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4269894367) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from), gdc.ConstTypePtr(&to), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RandomNumberGenerator) RandiRange(from int, to int, ) int {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("randi_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 50157827) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from), gdc.ConstTypePtr(&to), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RandomNumberGenerator) Randomize()  {
  classNameV := StringNameFromStr("RandomNumberGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("randomize")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *RandomNumberGenerator) GetPropSeed() int {
  panic("TODO: implement")
}

func (me *RandomNumberGenerator) SetPropSeed(value int) {
  panic("TODO: implement")
}

func (me *RandomNumberGenerator) GetPropState() int {
  panic("TODO: implement")
}

func (me *RandomNumberGenerator) SetPropState(value int) {
  panic("TODO: implement")
}