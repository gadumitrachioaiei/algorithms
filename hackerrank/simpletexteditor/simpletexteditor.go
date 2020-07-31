// Package main : https://www.hackerrank.com/challenges/simple-text-editor/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	countOps := 0
	_, err := fmt.Scanln(&countOps)
	if err != nil {
		panic(err)
	}
	lines := make([]string, countOps)
	r := bufio.NewReader(os.Stdin)
	for i := 0; i < len(lines); i++ {
		lines[i], err = r.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if len(lines) == 1 {
			panic("unexpected empty operation")
		}
		if lines[i][len(lines[i])-1] == '\n' {
			lines[i] = lines[i][:len(lines[i])-1]
		}
	}
	ops := make([]Operation, len(lines))
	for i, line := range lines {
		t, err := strconv.Atoi(line[:1])
		if err != nil {
			panic(err)
		}
		if Type(t) == Undo {
			ops[i] = Operation{Type: Undo}
			continue
		}
		if Type(t) == Append {
			ops[i] = Operation{Type: Append, Arg: line[2:]}
			continue
		}
		arg, err := strconv.Atoi(line[2:])
		if err != nil {
			panic(err)
		}
		ops[i] = Operation{Type: Type(t), Arg: arg}
	}
	s := ""
	// we need to apply the operations in order
	var changes []int
	for i, op := range ops {
		switch op.Type {
		case Append:
			ops[i].State = s
			s = s + op.Arg.(string)
			changes = append(changes, i)
		case Delete:
			ops[i].State = s
			s = s[:len(s)-op.Arg.(int)]
			changes = append(changes, i)
		case Print:
			fmt.Printf("%c\n", s[op.Arg.(int)-1])
		case Undo:
			op = ops[changes[len(changes)-1]]
			s = op.State
			changes = changes[:len(changes)-1]
		}
	}
}

type Type int

const (
	Append Type = iota + 1
	Delete
	Print
	Undo
)

type Operation struct {
	Type  Type
	Arg   interface{}
	State string
}
