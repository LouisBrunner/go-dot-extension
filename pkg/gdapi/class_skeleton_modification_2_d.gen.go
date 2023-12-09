// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SkeletonModification2D struct {
  obj gdc.ObjectPtr
}

func (me *SkeletonModification2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SkeletonModification2D) BaseClass() string {
  return "SkeletonModification2D"
}



// Enums

func (me *SkeletonModification2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SkeletonModification2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SkeletonModification2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SkeletonModification2D) XExecute(delta float32, )  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2D) XSetupModification(modification_stack SkeletonModificationStack2D, )  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2D) XDrawEditorGizmo()  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2D) SetEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2D) GetEnabled()  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2D) GetModificationStack()  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2D) SetIsSetup(is_setup bool, )  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2D) GetIsSetup()  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2D) SetExecutionMode(execution_mode int, )  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2D) GetExecutionMode()  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2D) ClampAngle(angle float32, min float32, max float32, invert bool, )  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2D) SetEditorDrawGizmo(draw_gizmo bool, )  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2D) GetEditorDrawGizmo()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
