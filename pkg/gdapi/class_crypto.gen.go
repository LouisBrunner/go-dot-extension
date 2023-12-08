// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Crypto struct {
  obj gdc.ObjectPtr
}

func createCrypto(obj gdc.ObjectPtr) *Crypto {
  return &Crypto{
    obj: obj,
  }
}

func (me *Crypto) BaseClass() string {
  return "Crypto"
}
