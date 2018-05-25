package main

import (
	"fmt"
	"strconv"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"
	"bytes"
)

func main(){
	var difficulty bytes.Buffer
	difficulty.WriteString("0")

	for k:= 0; k<15; k++ {

		start := time.Now()

		for i := 0; ; i++ {

			hashFunc := sha256.New()
			hashFunc.Write([]byte("Proof of Work sample : " + strconv.Itoa(i)))
			hashResult := hex.EncodeToString(hashFunc.Sum(nil))

			if strings.HasPrefix(hashResult, difficulty.String()) {
				elapsed := time.Since(start)
				fmt.Printf("In %s and %d iteraction found the following difficulty : %s \n", elapsed, i+1, difficulty.String())
				break
			}

			hashFunc.Reset()
		}

		difficulty.WriteString("0")

	}
}