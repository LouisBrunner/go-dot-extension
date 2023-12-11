// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type IP struct {
  obj gdc.ObjectPtr
}

func (me *IP) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *IP) BaseClass() string {
  return "IP"
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
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 396864159) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(host.AsCTypePtr()), gdc.ConstTypePtr(&ip_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *IP) ResolveHostnameAddresses(host String, ip_type IPType, ) PackedStringArray {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("resolve_hostname_addresses")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3462780090) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(host.AsCTypePtr()), gdc.ConstTypePtr(&ip_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *IP) ResolveHostnameQueueItem(host String, ip_type IPType, ) int {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("resolve_hostname_queue_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3936392508) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(host.AsCTypePtr()), gdc.ConstTypePtr(&ip_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *IP) GetResolveItemStatus(id int, ) IPResolverStatus {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resolve_item_status")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3812250196) // FIXME: should cache?
  var ret IPResolverStatus
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *IP) GetResolveItemAddress(id int, ) String {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resolve_item_address")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *IP) GetResolveItemAddresses(id int, ) Array {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resolve_item_addresses")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 663333327) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *IP) EraseResolveItem(id int, )  {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("erase_resolve_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *IP) GetLocalAddresses() PackedStringArray {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_local_addresses")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *IP) GetLocalInterfaces() Dictionary {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_local_interfaces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *IP) ClearCache(hostname String, )  {
  classNameV := StringNameFromStr("IP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_cache")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3005725572) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(hostname.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
