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

func  (me *Crypto) GenerateRandomBytes(size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Crypto) GenerateRsa(size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Crypto) GenerateSelfSignedCertificate(key CryptoKey, issuer_name String, not_before String, not_after String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Crypto) Sign(hash_type HashingContextHashType, hash PackedByteArray, key CryptoKey, ) { // TODO: return value
  // TODO: implement
}

func  (me *Crypto) Verify(hash_type HashingContextHashType, hash PackedByteArray, signature PackedByteArray, key CryptoKey, ) { // TODO: return value
  // TODO: implement
}

func  (me *Crypto) Encrypt(key CryptoKey, plaintext PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *Crypto) Decrypt(key CryptoKey, ciphertext PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *Crypto) HmacDigest(hash_type HashingContextHashType, key PackedByteArray, msg PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *Crypto) ConstantTimeCompare(trusted PackedByteArray, received PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
