package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

const (
	defaultEccPrivateFileName="eccprivate.pem"
	defaultEccPublishFileName="eccpublic.pem"

	defaultEccPrivateKeyPrefix="ECC PRIVATE KEY"
	defaultEccPublicKeyPrefix="ECC PUBLIC KEY"
)

func init(){
	log.SetFlags(log.Ldate|log.Lshortfile)
}

func GetEccKey() error{
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err!=nil{
		return err
	}
	// x509PrivateKey
	x509PrivateKey, err := x509.MarshalECPrivateKey(privateKey)
	if err!=nil{
		return err
	}

	block := pem.Block{
		Type:  defaultEccPrivateKeyPrefix,
		Bytes: x509PrivateKey,
	}
	file, err := os.Create(defaultEccPrivateFileName)
	if err!=nil{
		return err
	}
	defer file.Close()
	if err=pem.Encode(file, &block);err!=nil{
		return err
	}

	// x509PublicKey
	x509PublicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err!=nil {
		return err
	}
	publicBlock := pem.Block{
		Type:  defaultEccPublicKeyPrefix,
		Bytes: x509PublicKey,
	}
	publicFile, err := os.Create(defaultEccPublishFileName)
	if err!=nil {
		return err
	}
	defer publicFile.Close()
	if err=pem.Encode(publicFile,&publicBlock);err!=nil{
		return err
	}
	return nil
}