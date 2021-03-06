package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("bismillah")
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}
	var err error
	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list()
	case "add":
		err = add(strings.Join(flag.Args()[1:], ""))
	default:
		err = fmt.Errorf("unknown subcommand %s", cmd)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}
	}
}

const dbpath = "mydb.pb"

func add(text string) error {
	task := &todo.Task{
		Text: text,
		Done: false,
	}
	b, err := proto.Marshal(task)
	if err != nil {
		return fmt.Errorf("could not encode task: %v", err)
	}
	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND|0666)
	if err != nil {
		return fmt.Errorf("could not open %s: %v", dbPath, err)
	}
	_, err := f.Write(b)
	if err != nil {
		return fmt.Errorf("could not write task to file: %v", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close file: %s: %v", dbpath, err)
	}
	fmt.Println(proto.MarshalTextString(task))
	return nil
}
func list() error {
	return nil
}
