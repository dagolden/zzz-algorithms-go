// Copyright 2015 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// given input consisting of integer pairs denoting connections,
// output pairs that are connected (???)
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func readNodeCount(r *bufio.Reader) (v int) {
	_, err := fmt.Fscanf(r, "%d\n", &v)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func readPair(r *bufio.Reader) (p, q int, err error) {
	_, err = fmt.Fscanf(r, "%d %d\n", &p, &q)
	if err == io.EOF {
		return 0, 0, err
	}
	if err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	in := bufio.NewReader(os.Stdin)
	nc := readNodeCount(in)
	fmt.Println("Node count:", nc)

	for {
		p, q, err := readPair(in)
		if err != nil {
			break
		}
		fmt.Println("Pair:", p, q)
	}

}
