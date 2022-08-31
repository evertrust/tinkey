package main

import (
	"log"
	"os"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/core/registry"
	"github.com/google/tink/go/insecurecleartextkeyset"
	"github.com/google/tink/go/keyset"
)

func GenerateKeyset() *keyset.Handle {
	handle, err := keyset.NewHandle(aead.AES256GCMKeyTemplate())
	if err != nil {
		log.Fatal(err)
	}
	return handle
}

func WriteKeySet(path, masterKeyURI string, handle *keyset.Handle) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err.Error())
	}
	jsonWriter := keyset.NewJSONWriter(f)
	if masterKeyURI != "" {
		kmsClient, err := registry.GetKMSClient(masterKeyURI)
		if err != nil {
			log.Fatal(err)
		}
		registry.RegisterKMSClient(kmsClient)
		masterKey, err := kmsClient.GetAEAD(masterKeyURI)
		if err != nil {
			log.Fatal(err)
		}
		handle.Write(jsonWriter, masterKey)
	} else {
		insecurecleartextkeyset.Write(handle, jsonWriter)
	}
}

func ReadKeySet(path, masterKeyURI string) *keyset.Handle {
	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		log.Fatal(err.Error())
	}
	jsonReader := keyset.NewJSONReader(f)
	var handle *keyset.Handle
	if masterKeyURI != "" {
		kmsClient, err := registry.GetKMSClient(masterKeyURI)
		if err != nil {
			log.Fatal(err)
		}
		registry.RegisterKMSClient(kmsClient)
		masterKey, err := kmsClient.GetAEAD(masterKeyURI)
		if err != nil {
			log.Fatal(err)
		}
		handle, err = keyset.Read(jsonReader, masterKey)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		handle, err = insecurecleartextkeyset.Read(jsonReader)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	return handle
}
