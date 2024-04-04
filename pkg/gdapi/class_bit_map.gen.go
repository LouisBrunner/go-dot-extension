// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type BitMap struct {
  Resource
}

func (me *BitMap) BaseClass() string {
  return "BitMap"
}

func NewBitMap() *BitMap {
  str := StringNameFromStr("BitMap") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &BitMap{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *BitMap) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *BitMap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *BitMap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *BitMap) Create(size Vector2i, )  {
  classNameV := StringNameFromStr("BitMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BitMap) CreateFromImageAlpha(image Image, threshold float64, )  {
  classNameV := StringNameFromStr("BitMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_from_image_alpha")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 106271684) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{image.AsCTypePtr(), gdc.ConstTypePtr(&threshold) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BitMap) SetBitv(position Vector2i, bit bool, )  {
  classNameV := StringNameFromStr("BitMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bitv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4153096796) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&bit) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BitMap) SetBit(x int64, y int64, bit bool, )  {
  classNameV := StringNameFromStr("BitMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1383440665) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x) , gdc.ConstTypePtr(&y) , gdc.ConstTypePtr(&bit) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BitMap) GetBitv(position Vector2i, ) bool {
  classNameV := StringNameFromStr("BitMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bitv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3900751641) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BitMap) GetBit(x int64, y int64, ) bool {
  classNameV := StringNameFromStr("BitMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2522259332) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x) , gdc.ConstTypePtr(&y) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&x)
  pinner.Pin(&y)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BitMap) SetBitRect(rect Rect2i, bit bool, )  {
  classNameV := StringNameFromStr("BitMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bit_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 472162941) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{rect.AsCTypePtr(), gdc.ConstTypePtr(&bit) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BitMap) GetTrueBitCount() int64 {
  classNameV := StringNameFromStr("BitMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_true_bit_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BitMap) GetSize() Vector2i {
  classNameV := StringNameFromStr("BitMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BitMap) Resize(new_size Vector2i, )  {
  classNameV := StringNameFromStr("BitMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("resize")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{new_size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BitMap) GrowMask(pixels int64, rect Rect2i, )  {
  classNameV := StringNameFromStr("BitMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("grow_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3317281434) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixels) , rect.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BitMap) ConvertToImage() Image {
  classNameV := StringNameFromStr("BitMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("convert_to_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4190603485) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BitMap) OpaqueToPolygons(rect Rect2i, epsilon float64, ) []PackedVector2Array {
  classNameV := StringNameFromStr("BitMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("opaque_to_polygons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 48478126) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{rect.AsCTypePtr(), gdc.ConstTypePtr(&epsilon) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&epsilon)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[PackedVector2Array](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
