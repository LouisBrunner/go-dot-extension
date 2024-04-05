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

type ptrsForBitMapList struct {
  fnCreate gdc.MethodBindPtr
  fnCreateFromImageAlpha gdc.MethodBindPtr
  fnSetBitv gdc.MethodBindPtr
  fnSetBit gdc.MethodBindPtr
  fnGetBitv gdc.MethodBindPtr
  fnGetBit gdc.MethodBindPtr
  fnSetBitRect gdc.MethodBindPtr
  fnGetTrueBitCount gdc.MethodBindPtr
  fnGetSize gdc.MethodBindPtr
  fnResize gdc.MethodBindPtr
  fnGrowMask gdc.MethodBindPtr
  fnConvertToImage gdc.MethodBindPtr
  fnOpaqueToPolygons gdc.MethodBindPtr
}

var ptrsForBitMap ptrsForBitMapList

func initBitMapPtrs(iface gdc.Interface) {

  className := StringNameFromStr("BitMap")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("create")
    defer methodName.Destroy()
    ptrsForBitMap.fnCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
  }
  {
    methodName := StringNameFromStr("create_from_image_alpha")
    defer methodName.Destroy()
    ptrsForBitMap.fnCreateFromImageAlpha = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 106271684))
  }
  {
    methodName := StringNameFromStr("set_bitv")
    defer methodName.Destroy()
    ptrsForBitMap.fnSetBitv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4153096796))
  }
  {
    methodName := StringNameFromStr("set_bit")
    defer methodName.Destroy()
    ptrsForBitMap.fnSetBit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1383440665))
  }
  {
    methodName := StringNameFromStr("get_bitv")
    defer methodName.Destroy()
    ptrsForBitMap.fnGetBitv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3900751641))
  }
  {
    methodName := StringNameFromStr("get_bit")
    defer methodName.Destroy()
    ptrsForBitMap.fnGetBit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2522259332))
  }
  {
    methodName := StringNameFromStr("set_bit_rect")
    defer methodName.Destroy()
    ptrsForBitMap.fnSetBitRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 472162941))
  }
  {
    methodName := StringNameFromStr("get_true_bit_count")
    defer methodName.Destroy()
    ptrsForBitMap.fnGetTrueBitCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_size")
    defer methodName.Destroy()
    ptrsForBitMap.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
  }
  {
    methodName := StringNameFromStr("resize")
    defer methodName.Destroy()
    ptrsForBitMap.fnResize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
  }
  {
    methodName := StringNameFromStr("grow_mask")
    defer methodName.Destroy()
    ptrsForBitMap.fnGrowMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3317281434))
  }
  {
    methodName := StringNameFromStr("convert_to_image")
    defer methodName.Destroy()
    ptrsForBitMap.fnConvertToImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4190603485))
  }
  {
    methodName := StringNameFromStr("opaque_to_polygons")
    defer methodName.Destroy()
    ptrsForBitMap.fnOpaqueToPolygons = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 48478126))
  }
}

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
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBitMap.fnCreate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BitMap) CreateFromImageAlpha(image Image, threshold float64, )  {
  cargs := []gdc.ConstTypePtr{image.AsCTypePtr(), gdc.ConstTypePtr(&threshold) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBitMap.fnCreateFromImageAlpha), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BitMap) SetBitv(position Vector2i, bit bool, )  {
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&bit) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBitMap.fnSetBitv), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BitMap) SetBit(x int64, y int64, bit bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x) , gdc.ConstTypePtr(&y) , gdc.ConstTypePtr(&bit) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBitMap.fnSetBit), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BitMap) GetBitv(position Vector2i, ) bool {
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBitMap.fnGetBitv), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BitMap) GetBit(x int64, y int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x) , gdc.ConstTypePtr(&y) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&x)
  pinner.Pin(&y)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBitMap.fnGetBit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BitMap) SetBitRect(rect Rect2i, bit bool, )  {
  cargs := []gdc.ConstTypePtr{rect.AsCTypePtr(), gdc.ConstTypePtr(&bit) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBitMap.fnSetBitRect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BitMap) GetTrueBitCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBitMap.fnGetTrueBitCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BitMap) GetSize() Vector2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBitMap.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BitMap) Resize(new_size Vector2i, )  {
  cargs := []gdc.ConstTypePtr{new_size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBitMap.fnResize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BitMap) GrowMask(pixels int64, rect Rect2i, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixels) , rect.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBitMap.fnGrowMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BitMap) ConvertToImage() Image {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBitMap.fnConvertToImage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BitMap) OpaqueToPolygons(rect Rect2i, epsilon float64, ) []PackedVector2Array {
  cargs := []gdc.ConstTypePtr{rect.AsCTypePtr(), gdc.ConstTypePtr(&epsilon) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&epsilon)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBitMap.fnOpaqueToPolygons), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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
