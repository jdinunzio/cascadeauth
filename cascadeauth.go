/* ===========================================================================
  -*- coding: utf-8 -*-

  Author: José Dinuncio <jdinunci@uc.edu.ve>
  Copyright (c) 2010. RedUC
  All Rights Reserved.

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License.

=========================================================================== */
package main

import ("os"
        "fmt"
        "exec"
        "flag"
        "strings"
        )

const MAX_LINE_LEN = 256
const MAX_PROC = 10

func exitOnError(err *os.Error) {
    // Exit the program if an error
    if *err != nil {
        fmt.Printf("ERROR: %v\n", *err)
        os.Exit(1)
    }
}

func exitOnBool(b bool, msg string) {
    // Exit the program if b is true
    if b {
        fmt.Printf(msg)
        os.Exit(1)
    }
}

func getLine(f *os.File) string {
    // Read a line from a file
    var line [MAX_LINE_LEN]byte
    i := 0
    for i=0; i < MAX_LINE_LEN; i++ {
        _, e := f.Read(line[i:i+1])
        if e == os.EOF      { break }
        exitOnError(&e)
        if line[i] == '\n'  { i++; break }
    }
    return string(line[0:i])
}

func initCmd(file *os.File) []* exec.Cmd {
    // Create a command for each line in file
    var cmd     [MAX_PROC]* exec.Cmd;
    var e       os.Error

    environ := os.Environ();
    var i int
    for i=0; i < MAX_PROC; i++ {
        line := getLine(file)
        if line == "" { break }
        parts := strings.Fields(line)
        cmd[i], e = exec.Run(parts[0], parts, environ, exec.Pipe, exec.Pipe, exec.Pipe)
        exitOnError(&e)
    }
    return cmd[0:i]
}


func main() {
    // Find the name of the config gile
    flag.Parse()
    exitOnBool(flag.NArg() != 1, "Error. Use:\n\t cascadeauth filename\n")

    // Open the config file
    fname := flag.Arg(0)
    file, e := os.Open(fname, os.O_RDONLY, 0)
    exitOnError(&e)

    // Get the commands
    cmd := initCmd(file)

    // Main loop
    var pre, post string
    for ;; {
        pre = getLine(os.Stdin)
        for i:=0; i < len(cmd); i++ {
            cmd[i].Stdin.WriteString(pre)
            post = getLine(cmd[i].Stdout)
            if strings.HasPrefix(post, "OK") { break }
        }
        fmt.Printf(post)
    }
}
