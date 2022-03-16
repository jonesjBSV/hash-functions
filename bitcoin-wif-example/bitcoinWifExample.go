package main

import (
	"encoding/hex"
	"fmt"
	"github.com/libsv/go-bk/base58"
	"github.com/libsv/go-bk/bec"
	"github.com/libsv/go-bk/chaincfg"
	"github.com/libsv/go-bk/crypto"
	"github.com/libsv/go-bk/wif"
)

func main() {

	privKey, _ := bec.NewPrivateKey(bec.S256())

	pKey := privKey.Serialise()

	version := make([]byte, 0)

	version = append(version, 0x80) //0xef for testnet

	pKey = append(version, pKey...)

	pKey = append(pKey, 0x01)

	fmt.Printf("Extended PrivKey: %s\n", hex.EncodeToString(pKey))

	chksum := crypto.Sha256d(pKey)[:4]

	pKey = append(pKey, chksum...)

	w := base58.Encode(pKey)
	wi, _ := wif.NewWIF(privKey, &chaincfg.MainNet, true)

	fmt.Printf("%s\n", hex.EncodeToString(privKey.Serialise()))

	fmt.Printf("WIF: %s\n", w)

	fmt.Printf("%s\n", wi.String())

}
