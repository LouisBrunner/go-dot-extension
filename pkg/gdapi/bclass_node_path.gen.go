// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NodePath struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewNodePath() NodePath {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeNodePath))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeNodePath, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return NodePath{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewNodePathFromNodePath(from NodePath, ) NodePath {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeNodePath))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeNodePath, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return NodePath{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewNodePathFromString(from String, ) NodePath {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeNodePath))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeNodePath, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return NodePath{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *NodePath) Destroy() {
  if me.ptr == nil {
    return
  }
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypeNodePath)
	me.iface.CallPtrDestructor(dtr, gdc.TypePtr(me.ptr))
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *NodePath) Type() gdc.VariantType {
  return gdc.VariantTypeNodePath
}

func (me *NodePath) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *NodePath) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *NodePath) IsAbsolute() bool {
  panic("TODO: implement")
}

func  (me *NodePath) GetNameCount() int {
  panic("TODO: implement")
}

func  (me *NodePath) GetName(idx int, ) StringName {
  panic("TODO: implement")
}

func  (me *NodePath) GetSubnameCount() int {
  panic("TODO: implement")
}

func  (me *NodePath) Hash() int {
  panic("TODO: implement")
}

func  (me *NodePath) GetSubname(idx int, ) StringName {
  panic("TODO: implement")
}

func  (me *NodePath) GetConcatenatedNames() StringName {
  panic("TODO: implement")
}

func  (me *NodePath) GetConcatenatedSubnames() StringName {
  panic("TODO: implement")
}

func  (me *NodePath) GetAsPropertyPath() NodePath {
  panic("TODO: implement")
}

func  (me *NodePath) IsEmpty() bool {
  panic("TODO: implement")
}

// Operators

func (me *NodePath) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *NodePath) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *NodePath) Not() bool {
  panic("TODO: implement")
}

func (me *NodePath) EqualsNodePath(right NodePath) bool {
  panic("TODO: implement")
}

func (me *NodePath) NotEqualsNodePath(right NodePath) bool {
  panic("TODO: implement")
}

func (me *NodePath) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *NodePath) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
