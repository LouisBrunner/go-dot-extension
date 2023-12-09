// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Crypto struct {
  obj gdc.ObjectPtr
}

func (me *Crypto) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Crypto) BaseClass() string {
  return "Crypto"
}



// Enums

func (me *Crypto) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Crypto) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Crypto) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Crypto) GenerateRandomBytes(size int, )  {
  panic("TODO: implement")
}

func  (me *Crypto) GenerateRsa(size int, )  {
  panic("TODO: implement")
}

func  (me *Crypto) GenerateSelfSignedCertificate(key CryptoKey, issuer_name String, not_before String, not_after String, )  {
  panic("TODO: implement")
}

func  (me *Crypto) Sign(hash_type HashingContextHashType, hash PackedByteArray, key CryptoKey, )  {
  panic("TODO: implement")
}

func  (me *Crypto) Verify(hash_type HashingContextHashType, hash PackedByteArray, signature PackedByteArray, key CryptoKey, )  {
  panic("TODO: implement")
}

func  (me *Crypto) Encrypt(key CryptoKey, plaintext PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *Crypto) Decrypt(key CryptoKey, ciphertext PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *Crypto) HmacDigest(hash_type HashingContextHashType, key PackedByteArray, msg PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *Crypto) ConstantTimeCompare(trusted PackedByteArray, received PackedByteArray, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
