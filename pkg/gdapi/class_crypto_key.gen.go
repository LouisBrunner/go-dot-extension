// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CryptoKey struct {
  obj gdc.ObjectPtr
}

func createCryptoKey(obj gdc.ObjectPtr) *CryptoKey {
  return &CryptoKey{
    obj: obj,
  }
}

func (me *CryptoKey) BaseClass() string {
  return "CryptoKey"
}
