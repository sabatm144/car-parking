package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
)

func main() {

    mode := flag.Int("mode", 1, "Select a mode to run")
    filePath := flag.String("filePath", "", "Say your input file path, rest i will do for you")

    flag.Parse()

    // get the actual option values.
    fmt.Println("mode:", *mode)
    fmt.Println("filePath:", *filePath)

    if *mode == 2 {
        fmt.Println("Processing file, please wait")
        if *filePath == "" {
            fmt.Println("Hey you didn’t mention your input file path")
            flag.PrintDefaults()
            os.Exit(1)
        }

        //TODO: read a file path and execute the commands

        os.Exit(1)
    }

    fmt.Println("hey you look good, say what I want to do for you")
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan() {
        text := scanner.Text()

        //TODO: execute the command

        if text == "exit" {
            fmt.Println("Good bye")
            os.Exit(1)
        }
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }

}