// Copyright 2015 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// given list of files on the command line, each with sorted words,
// merge them, retaining the sort order
package main

import (
	"bufio"
	"fmt"
	"github.com/dagolden/zzz-algorithms-go/pq"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	// skip command name
	args := os.Args[1:]

	q := pq.New(func(i, j interface{}) bool { return i.(string) < j.(string) })

	for _, f := range args {
		in, err := os.Open(f)
		if err != nil {
			log.Fatal("Couldn't open file", f, ":", err)
		}
		scanner := bufio.NewScanner(in)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			q.Push(scanner.Text())
		}
	}

	if q.Size() == 0 {
		return
	}

	fmt.Print(q.Pop().(string))
	for q.Size() != 0 {
		fmt.Print(" " + q.Pop().(string))
	}
	fmt.Print("\n")

	return
}
