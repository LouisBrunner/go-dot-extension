// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ImageTexture3D struct {
  obj gdc.ObjectPtr
}

func (me *ImageTexture3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ImageTexture3D) BaseClass() string {
  return "ImageTexture3D"
}



// Enums

func (me *ImageTexture3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ImageTexture3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ImageTexture3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ImageTexture3D) Create(format ImageFormat, width int, height int, depth int, use_mipmaps bool, data Image, ) Error {
  classNameV := StringNameFromStr("ImageTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130379827) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&depth), gdc.ConstTypePtr(&use_mipmaps), gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ImageTexture3D) Update(data Image, )  {
  classNameV := StringNameFromStr("ImageTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties