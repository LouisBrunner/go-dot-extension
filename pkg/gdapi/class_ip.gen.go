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

type IP struct {
  Object
}

func (me *IP) BaseClass() string {
  return "IP"
}

func NewIP() *IP {
  str := StringNameFromStr("IP") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &IP{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  IPResolverMaxQueries = "256" // TODO: construct correctly
  IPResolverInvalidId = "-1" // TODO: construct correctly
)

// Enums

type IPResolverStatus int
const (
  IPResolverStatusResolverStatusNone IPResolverStatus = 0
  IPResolverStatusResolverStatusWaiting IPResolverStatus = 1
  IPResolverStatusResolverStatusDone IPResolverStatus = 2
  IPResolverStatusResolverStatusError IPResolverStatus = 3
)

type IPType int
const (
  IPTypeTypeNone IPType = 0
  IPTypeTypeIpv4 IPType = 1
  IPTypeTypeIpv6 IPType = 2
  IPTypeTypeAny IPType = 3
)

func (me *IP) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *IP) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *IP) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *IP) ResolveHostname(host String, ip_type IPType, ) String {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("resolve_hostname")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4283295457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&ip_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&ip_type)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *IP) ResolveHostnameAddresses(host String, ip_type IPType, ) PackedStringArray {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("resolve_hostname_addresses")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 773767525) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&ip_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()
  pinner.Pin(&ip_type)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *IP) ResolveHostnameQueueItem(host String, ip_type IPType, ) int64 {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("resolve_hostname_queue_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1749894742) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&ip_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&ip_type)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *IP) GetResolveItemStatus(id int64, ) IPResolverStatus {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resolve_item_status")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3812250196) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret IPResolverStatus
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *IP) GetResolveItemAddress(id int64, ) String {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resolve_item_address")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *IP) GetResolveItemAddresses(id int64, ) Array {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resolve_item_addresses")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 663333327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *IP) EraseResolveItem(id int64, )  {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("erase_resolve_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *IP) GetLocalAddresses() PackedStringArray {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_local_addresses")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *IP) GetLocalInterfaces() []Dictionary {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_local_interfaces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *IP) ClearCache(hostname String, )  {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_cache")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3005725572) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{hostname.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
