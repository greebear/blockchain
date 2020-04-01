package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"log"
	"runtime"
)

func init(){
	log.SetFlags(log.Ldate|log.Lshortfile)
}
// The public key and plaintext are passed in for encryption
func EccEncrypt(plainText,key []byte)( cryptText []byte,err error){
	block, _:= pem.Decode(key)

	// painc
	defer func(){
		if err:=recover();err!=nil{
			switch err.(type){
			case runtime.Error:
				log.Println("runtime err: ",err,"Check that the key is correct")
			default:
				log.Println("error:",err)
			}
		}
	}()

	tempPublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err!=nil{
		return nil,err
	}
	// Decode to get the private key in the ecdsa package
	publicKeyForEcies:=tempPublicKey.(*ecdsa.PublicKey)
	// Convert to the public key in the ecies package in the ethereum package
	publicKey:=ImportECDSAPublic(publicKeyForEcies)
	crypttext,err:=Encrypt(rand.Reader, publicKey, plainText, nil, nil)

	return  crypttext,err


}
// The private key and plaintext are passed in for decryption
func EccDecrypt(cryptText,key []byte)( msg []byte,err error){
	block, _:= pem.Decode(key)

	// painc
	defer func(){
		if err:=recover();err!=nil{
			switch err.(type){
			case runtime.Error:
				log.Println("runtime err:",err,"Check that the key is correct")
			default:
				log.Println("error:",err)
			}
		}
	}()

	tempPrivateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err!=nil{
		return nil,err
	}
	// Decode to get the private key in the ecdsa package
	// Convert to the private key in the ecies package in the ethereum package
	privateKey:=ImportECDSA(tempPrivateKey)

	plainText,err:=privateKey.Decrypt(cryptText,nil,nil)
	if err!=nil{
		return nil,err
	}
	return plainText,nil
}