// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SegmentShape2D struct {
  obj gdc.ObjectPtr
}

func (me *SegmentShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SegmentShape2D) BaseClass() string {
  return "SegmentShape2D"
}



// Enums

func (me *SegmentShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SegmentShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SegmentShape2D) SetA(a Vector2, )  {
  panic("TODO: implement")
}

func  (me *SegmentShape2D) GetA()  {
  panic("TODO: implement")
}

func  (me *SegmentShape2D) SetB(b Vector2, )  {
  panic("TODO: implement")
}

func  (me *SegmentShape2D) GetB()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
