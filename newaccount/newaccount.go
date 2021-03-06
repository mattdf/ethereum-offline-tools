package main

import (
	"fmt"
	"log"

	"os"

	"github.com/ethereum/go-ethereum/accounts"
	t0mkcommon "github.com/t0mk/ethereum-offline-tools/common"
)

func usage() {
	fmt.Println("This will create new account in given keystore dir.")
	fmt.Println("Usage:")
	fmt.Println(os.Args[0], "keystore_dir")
	fmt.Println("Example:")
	fmt.Println(os.Args[0], "/home/tomk/.ethereum/keystore")
	os.Exit(2)
}

func main() {
	// keystore_dir, from_address, to_address, amount, nonce

	if len(os.Args) != 2 {
		usage()
	}
	stat, err := os.Stat(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if !stat.IsDir() {
		log.Fatal(os.Args[1], "is not a directory")
	}

	am := accounts.NewManager(os.Args[1], 262144, 1)

	base_prompt := "New password to encrypt your private key,"
	p := t0mkcommon.GetCheckedPassword(base_prompt)
	fmt.Println("Password OK, now creating new account")

	new_account, err := am.NewAccount(p)
	if err != nil {
		log.Fatal("Failed to create new account", err)
	} else {
		fmt.Println("Created new account alright")
	}
	fmt.Println()
	fmt.Println("New account in", new_account.File)
	fmt.Println("New account address", new_account.Address.Hex())

}
