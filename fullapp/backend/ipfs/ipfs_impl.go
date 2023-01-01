package main

import (
	// "bytes"
	"fmt"
	"os"

	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	sh := shell.NewLocalShell()
	info, err := sh.ID()
	if err != nil {
		fmt.Println("failed to get info about local node")
	}
	for i := range info.Addresses {
		fmt.Printf("The id: %s, for address: %s has publickey: %s", info.ID, info.Addresses[i], info.PublicKey)
		
	}

	

	cid, err := sh.Add(strings.NewReader("hello world!"))
	if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s", err)
        os.Exit(1)
	}
    fmt.Printf("added %s", cid)
}

// func connectToIPFS() *shell.Shell {
// 	return shell.NewShell("localhost:5001")
// }

// func (sh *shell.Shell) addToIPFS() {
// 	cid, err := s.Add(strings.NewReader("hello world!"))
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error: %s", err)
// 		os.Exit(1)
// 	}
// 	fmt.Printf("added %s\n", cid)
// 	out := fmt.Sprintf("%s.txt", cid)
// 	err = sh.Get(cid, out)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error: %s", err)
// 		os.Exit(1)
// 	}
// }

// func (sh *shell.Shell) getDataFromIPFS() {
// 	data, err := sh.Cat(cid)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error: %s", err)
// 		os.Exit(1)
// 	}
// 	// ...so we convert it to a string by passing it through
// 	// a buffer first. A 'costly' but useful process.
// 	// https://golangcode.com/convert-io-readcloser-to-a-string/
// 	buf := new(bytes.Buffer)
// 	buf.ReadFrom(data)
// 	newStr := buf.String()
// 	fmt.Printf("data %s", newStr)
// }

