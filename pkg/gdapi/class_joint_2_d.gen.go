// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Joint2D struct {
  obj gdc.ObjectPtr
}

func (me *Joint2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Joint2D) BaseClass() string {
  return "Joint2D"
}



// Enums

func (me *Joint2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Joint2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Joint2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Joint2D) SetNodeA(node NodePath, )  {
  classNameV := StringNameFromStr("Joint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_node_a")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Joint2D) GetNodeA() NodePath {
  classNameV := StringNameFromStr("Joint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_a")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Joint2D) SetNodeB(node NodePath, )  {
  classNameV := StringNameFromStr("Joint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_node_b")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Joint2D) GetNodeB() NodePath {
  classNameV := StringNameFromStr("Joint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_b")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Joint2D) SetBias(bias float32, )  {
  classNameV := StringNameFromStr("Joint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Joint2D) GetBias() float32 {
  classNameV := StringNameFromStr("Joint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Joint2D) SetExcludeNodesFromCollision(enable bool, )  {
  classNameV := StringNameFromStr("Joint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_exclude_nodes_from_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Joint2D) GetExcludeNodesFromCollision() bool {
  classNameV := StringNameFromStr("Joint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_exclude_nodes_from_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Joint2D) GetPropNodeA() NodePath {
  panic("TODO: implement")
}

func (me *Joint2D) SetPropNodeA(value NodePath) {
  panic("TODO: implement")
}

func (me *Joint2D) GetPropNodeB() NodePath {
  panic("TODO: implement")
}

func (me *Joint2D) SetPropNodeB(value NodePath) {
  panic("TODO: implement")
}

func (me *Joint2D) GetPropBias() float32 {
  panic("TODO: implement")
}

func (me *Joint2D) SetPropBias(value float32) {
  panic("TODO: implement")
}

func (me *Joint2D) GetPropDisableCollision() bool {
  panic("TODO: implement")
}

func (me *Joint2D) SetPropDisableCollision(value bool) {
  panic("TODO: implement")
}