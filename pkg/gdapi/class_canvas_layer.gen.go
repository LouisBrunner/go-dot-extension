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

type ptrsForCanvasLayerList struct {
  fnSetLayer gdc.MethodBindPtr
  fnGetLayer gdc.MethodBindPtr
  fnSetVisible gdc.MethodBindPtr
  fnIsVisible gdc.MethodBindPtr
  fnShow gdc.MethodBindPtr
  fnHide gdc.MethodBindPtr
  fnSetTransform gdc.MethodBindPtr
  fnGetTransform gdc.MethodBindPtr
  fnGetFinalTransform gdc.MethodBindPtr
  fnSetOffset gdc.MethodBindPtr
  fnGetOffset gdc.MethodBindPtr
  fnSetRotation gdc.MethodBindPtr
  fnGetRotation gdc.MethodBindPtr
  fnSetScale gdc.MethodBindPtr
  fnGetScale gdc.MethodBindPtr
  fnSetFollowViewport gdc.MethodBindPtr
  fnIsFollowingViewport gdc.MethodBindPtr
  fnSetFollowViewportScale gdc.MethodBindPtr
  fnGetFollowViewportScale gdc.MethodBindPtr
  fnSetCustomViewport gdc.MethodBindPtr
  fnGetCustomViewport gdc.MethodBindPtr
  fnGetCanvas gdc.MethodBindPtr
}

var ptrsForCanvasLayer ptrsForCanvasLayerList

func initCanvasLayerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CanvasLayer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_layer")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnSetLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_layer")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnGetLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_visible")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnSetVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_visible")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnIsVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("show")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnShow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("hide")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnHide = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_transform")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761652528))
  }
  {
    methodName := StringNameFromStr("get_transform")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
  }
  {
    methodName := StringNameFromStr("get_final_transform")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnGetFinalTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
  }
  {
    methodName := StringNameFromStr("set_offset")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnSetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_offset")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnGetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_rotation")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnSetRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_rotation")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnGetRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_scale")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnSetScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_scale")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnGetScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_follow_viewport")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnSetFollowViewport = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_following_viewport")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnIsFollowingViewport = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_follow_viewport_scale")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnSetFollowViewportScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_follow_viewport_scale")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnGetFollowViewportScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_custom_viewport")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnSetCustomViewport = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
  }
  {
    methodName := StringNameFromStr("get_custom_viewport")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnGetCustomViewport = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3160264692))
  }
  {
    methodName := StringNameFromStr("get_canvas")
    defer methodName.Destroy()
    ptrsForCanvasLayer.fnGetCanvas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
  }
}

type CanvasLayer struct {
  Node
}

func (me *CanvasLayer) BaseClass() string {
  return "CanvasLayer"
}

func NewCanvasLayer() *CanvasLayer {
  str := StringNameFromStr("CanvasLayer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CanvasLayer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CanvasLayer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CanvasLayer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CanvasLayer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CanvasLayer) SetLayer(layer int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnSetLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasLayer) GetLayer() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnGetLayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasLayer) SetVisible(visible bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnSetVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasLayer) IsVisible() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnIsVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasLayer) Show()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnShow), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasLayer) Hide()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnHide), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasLayer) SetTransform(transform Transform2D, )  {
  cargs := []gdc.ConstTypePtr{transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasLayer) GetTransform() Transform2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasLayer) GetFinalTransform() Transform2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnGetFinalTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasLayer) SetOffset(offset Vector2, )  {
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnSetOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasLayer) GetOffset() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnGetOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasLayer) SetRotation(radians float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnSetRotation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasLayer) GetRotation() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnGetRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasLayer) SetScale(scale Vector2, )  {
  cargs := []gdc.ConstTypePtr{scale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnSetScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasLayer) GetScale() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnGetScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasLayer) SetFollowViewport(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnSetFollowViewport), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasLayer) IsFollowingViewport() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnIsFollowingViewport), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasLayer) SetFollowViewportScale(scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnSetFollowViewportScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasLayer) GetFollowViewportScale() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnGetFollowViewportScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasLayer) SetCustomViewport(viewport Node, )  {
  cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnSetCustomViewport), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasLayer) GetCustomViewport() Node {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnGetCustomViewport), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasLayer) GetCanvas() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasLayer.fnGetCanvas), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type CanvasLayerVisibilityChangedSignalFn func()

func (me *CanvasLayer) ConnectVisibilityChanged(subs SignalSubscribers, fn CanvasLayerVisibilityChangedSignalFn) {
  sig := StringNameFromStr("visibility_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CanvasLayer) DisconnectVisibilityChanged(subs SignalSubscribers, fn CanvasLayerVisibilityChangedSignalFn) {
  sig := StringNameFromStr("visibility_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
