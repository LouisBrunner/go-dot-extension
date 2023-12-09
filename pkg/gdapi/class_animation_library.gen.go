// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationLibrary struct {
  obj gdc.ObjectPtr
}

func (me *AnimationLibrary) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationLibrary) BaseClass() string {
  return "AnimationLibrary"
}



// Enums

func (me *AnimationLibrary) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationLibrary) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AnimationLibrary) AddAnimation(name StringName, animation Animation, )  {
  panic("TODO: implement")
}

func  (me *AnimationLibrary) RemoveAnimation(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationLibrary) RenameAnimation(name StringName, newname StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationLibrary) HasAnimation(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationLibrary) GetAnimation(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationLibrary) GetAnimationList()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
