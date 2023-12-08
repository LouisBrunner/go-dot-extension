// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type IP struct {
  obj gdc.ObjectPtr
}

func (me *IP) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *IP) BaseClass() string {
  return "IP"
}

// TODO: needed?
// const (
// // )

var (
  IPResolverMaxQueries = "256" // TODO: construct correctly
  IPResolverInvalidId = "-1" // TODO: construct correctly
)

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

func  (me *IP) ResolveHostname(host String, ip_type IPType, ) { // TODO: return value
  // TODO: implement
}

func  (me *IP) ResolveHostnameAddresses(host String, ip_type IPType, ) { // TODO: return value
  // TODO: implement
}

func  (me *IP) ResolveHostnameQueueItem(host String, ip_type IPType, ) { // TODO: return value
  // TODO: implement
}

func  (me *IP) GetResolveItemStatus(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *IP) GetResolveItemAddress(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *IP) GetResolveItemAddresses(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *IP) EraseResolveItem(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *IP) GetLocalAddresses() { // TODO: return value
  // TODO: implement
}

func  (me *IP) GetLocalInterfaces() { // TODO: return value
  // TODO: implement
}

func  (me *IP) ClearCache(hostname String, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
