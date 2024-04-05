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

type ptrsForCSGShape3DList struct {
  fnIsRootShape gdc.MethodBindPtr
  fnSetOperation gdc.MethodBindPtr
  fnGetOperation gdc.MethodBindPtr
  fnSetSnap gdc.MethodBindPtr
  fnGetSnap gdc.MethodBindPtr
  fnSetUseCollision gdc.MethodBindPtr
  fnIsUsingCollision gdc.MethodBindPtr
  fnSetCollisionLayer gdc.MethodBindPtr
  fnGetCollisionLayer gdc.MethodBindPtr
  fnSetCollisionMask gdc.MethodBindPtr
  fnGetCollisionMask gdc.MethodBindPtr
  fnSetCollisionMaskValue gdc.MethodBindPtr
  fnGetCollisionMaskValue gdc.MethodBindPtr
  fnSetCollisionLayerValue gdc.MethodBindPtr
  fnGetCollisionLayerValue gdc.MethodBindPtr
  fnSetCollisionPriority gdc.MethodBindPtr
  fnGetCollisionPriority gdc.MethodBindPtr
  fnSetCalculateTangents gdc.MethodBindPtr
  fnIsCalculatingTangents gdc.MethodBindPtr
  fnGetMeshes gdc.MethodBindPtr
}

var ptrsForCSGShape3D ptrsForCSGShape3DList

func initCSGShape3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CSGShape3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("is_root_shape")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnIsRootShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_operation")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnSetOperation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 811425055))
  }
  {
    methodName := StringNameFromStr("get_operation")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnGetOperation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2662425879))
  }
  {
    methodName := StringNameFromStr("set_snap")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnSetSnap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_snap")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnGetSnap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_use_collision")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnSetUseCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_using_collision")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnIsUsingCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_collision_layer")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnSetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_collision_layer")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnGetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_collision_mask")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_collision_mask")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_collision_mask_value")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnSetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_collision_mask_value")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnGetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("set_collision_layer_value")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnSetCollisionLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_collision_layer_value")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnGetCollisionLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("set_collision_priority")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnSetCollisionPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_collision_priority")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnGetCollisionPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_calculate_tangents")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnSetCalculateTangents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_calculating_tangents")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnIsCalculatingTangents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_meshes")
    defer methodName.Destroy()
    ptrsForCSGShape3D.fnGetMeshes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
}

type CSGShape3D struct {
  GeometryInstance3D
}

func (me *CSGShape3D) BaseClass() string {
  return "CSGShape3D"
}

func NewCSGShape3D() *CSGShape3D {
  str := StringNameFromStr("CSGShape3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CSGShape3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type CSGShape3DOperation int
const (
  CSGShape3DOperationOperationUnion CSGShape3DOperation = 0
  CSGShape3DOperationOperationIntersection CSGShape3DOperation = 1
  CSGShape3DOperationOperationSubtraction CSGShape3DOperation = 2
)

func (me *CSGShape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CSGShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CSGShape3D) IsRootShape() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnIsRootShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGShape3D) SetOperation(operation CSGShape3DOperation, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&operation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnSetOperation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGShape3D) GetOperation() CSGShape3DOperation {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CSGShape3DOperation

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnGetOperation), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CSGShape3D) SetSnap(snap float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&snap) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnSetSnap), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGShape3D) GetSnap() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnGetSnap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGShape3D) SetUseCollision(operation bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&operation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnSetUseCollision), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGShape3D) IsUsingCollision() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnIsUsingCollision), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGShape3D) SetCollisionLayer(layer int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnSetCollisionLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGShape3D) GetCollisionLayer() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnGetCollisionLayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGShape3D) SetCollisionMask(mask int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGShape3D) GetCollisionMask() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGShape3D) SetCollisionMaskValue(layer_number int64, value bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnSetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGShape3D) GetCollisionMaskValue(layer_number int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnGetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGShape3D) SetCollisionLayerValue(layer_number int64, value bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnSetCollisionLayerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGShape3D) GetCollisionLayerValue(layer_number int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnGetCollisionLayerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGShape3D) SetCollisionPriority(priority float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnSetCollisionPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGShape3D) GetCollisionPriority() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnGetCollisionPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGShape3D) SetCalculateTangents(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnSetCalculateTangents), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGShape3D) IsCalculatingTangents() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnIsCalculatingTangents), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGShape3D) GetMeshes() Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGShape3D.fnGetMeshes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
