// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BitMap struct {
  obj gdc.ObjectPtr
}

func (me *BitMap) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *BitMap) BaseClass() string {
  return "BitMap"
}



// Enums

func (me *BitMap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *BitMap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *BitMap) Create(size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *BitMap) CreateFromImageAlpha(image Image, threshold float32, )  {
  panic("TODO: implement")
}

func  (me *BitMap) SetBitv(position Vector2i, bit bool, )  {
  panic("TODO: implement")
}

func  (me *BitMap) SetBit(x int, y int, bit bool, )  {
  panic("TODO: implement")
}

func  (me *BitMap) GetBitv(position Vector2i, )  {
  panic("TODO: implement")
}

func  (me *BitMap) GetBit(x int, y int, )  {
  panic("TODO: implement")
}

func  (me *BitMap) SetBitRect(rect Rect2i, bit bool, )  {
  panic("TODO: implement")
}

func  (me *BitMap) GetTrueBitCount()  {
  panic("TODO: implement")
}

func  (me *BitMap) GetSize()  {
  panic("TODO: implement")
}

func  (me *BitMap) Resize(new_size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *BitMap) GrowMask(pixels int, rect Rect2i, )  {
  panic("TODO: implement")
}

func  (me *BitMap) ConvertToImage()  {
  panic("TODO: implement")
}

func  (me *BitMap) OpaqueToPolygons(rect Rect2i, epsilon float32, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
