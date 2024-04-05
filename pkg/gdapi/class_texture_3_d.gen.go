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

type ptrsForTexture3DList struct {
  fnXGetFormat gdc.MethodBindPtr
  fnXGetWidth gdc.MethodBindPtr
  fnXGetHeight gdc.MethodBindPtr
  fnXGetDepth gdc.MethodBindPtr
  fnXHasMipmaps gdc.MethodBindPtr
  fnXGetData gdc.MethodBindPtr
  fnGetFormat gdc.MethodBindPtr
  fnGetWidth gdc.MethodBindPtr
  fnGetHeight gdc.MethodBindPtr
  fnGetDepth gdc.MethodBindPtr
  fnHasMipmaps gdc.MethodBindPtr
  fnGetData gdc.MethodBindPtr
  fnCreatePlaceholder gdc.MethodBindPtr
}

var ptrsForTexture3D ptrsForTexture3DList

func initTexture3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Texture3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_format")
    defer methodName.Destroy()
    ptrsForTexture3D.fnGetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3847873762))
  }
  {
    methodName := StringNameFromStr("get_width")
    defer methodName.Destroy()
    ptrsForTexture3D.fnGetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_height")
    defer methodName.Destroy()
    ptrsForTexture3D.fnGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_depth")
    defer methodName.Destroy()
    ptrsForTexture3D.fnGetDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("has_mipmaps")
    defer methodName.Destroy()
    ptrsForTexture3D.fnHasMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_data")
    defer methodName.Destroy()
    ptrsForTexture3D.fnGetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("create_placeholder")
    defer methodName.Destroy()
    ptrsForTexture3D.fnCreatePlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 121922552))
  }
}

type Texture3D struct {
  Texture
}

func (me *Texture3D) BaseClass() string {
  return "Texture3D"
}

func NewTexture3D() *Texture3D {
  str := StringNameFromStr("Texture3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Texture3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Texture3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Texture3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Texture3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Texture3D) GetFormat() ImageFormat {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ImageFormat

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture3D.fnGetFormat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Texture3D) GetWidth() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture3D.fnGetWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Texture3D) GetHeight() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture3D.fnGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Texture3D) GetDepth() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture3D.fnGetDepth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Texture3D) HasMipmaps() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture3D.fnHasMipmaps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Texture3D) GetData() []Image {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture3D.fnGetData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Image](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *Texture3D) CreatePlaceholder() Resource {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewResource()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture3D.fnCreatePlaceholder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
