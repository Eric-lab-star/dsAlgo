package main

import (
	"bytes"
	"crypto/sha256"
	"encoding"
	"fmt"
	"log"
)

func main() {
	const (
		example1 = "this is a examle"
		examle2  = "second examle"
	)
	firstHash := sha256.New()
	firstHash.Write([]byte(example1))

	marshaler, ok := firstHash.(encoding.BinaryMarshaler) //convert firstHash to BinaryMarchaler interface type
	if !ok {
		log.Fatal("first hash does not implement encoding.BinaryMarshaler")
	}
	data, err := marshaler.MarshalBinary()
	if err != nil {
		log.Fatal("unable to marshal first Hash:", err)
	}
	secondHash := sha256.New()
	unmarshaler, ok := secondHash.(encoding.BinaryUnmarshaler)
	if !ok {
		log.Fatal("second Hash does not implement encoding.BinaryUnmarshaler")
	}
	if err := unmarshaler.UnmarshalBinary(data); err != nil {
		log.Fatal("unable to unmarshal hash:", err)
	}
	fmt.Printf("%x\n", secondHash.Sum(nil))
	fmt.Printf("%x\n", firstHash.Sum(nil))
	fmt.Println(bytes.Equal(firstHash.Sum(nil), secondHash.Sum(nil)))
}
