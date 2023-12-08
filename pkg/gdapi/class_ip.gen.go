// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

const (
  IPRESOLVER_MAX_QUERIES = 256
  IPRESOLVER_INVALID_ID = -1
)

type IPResolverStatus int
const (
  IPResolverStatusNone IPResolverStatus = 0
  IPResolverStatusWaiting IPResolverStatus = 1
  IPResolverStatusDone IPResolverStatus = 2
  IPResolverStatusError IPResolverStatus = 3
)

type IPType int
const (
  IPTypeNone IPType = 0
  IPTypeIpv4 IPType = 1
  IPTypeIpv6 IPType = 2
  IPTypeAny IPType = 3
)
