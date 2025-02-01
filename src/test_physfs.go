// go build -o main main.go physfs.go

package main

import (
    "fmt"
    "log"
    "os"
    // "path"
    "github.com/leonkasovan/ikemen/packages/physfs"
)

func main() {
    // exePath, _ := os.Executable()    
    if !physfs.Init(os.Args[0]) {
        log.Fatal("Failed to initialize PhysicsFS")
    }
    defer physfs.Deinit()

    if !physfs.Mount("example.zip", "/", 1) {
        fmt.Printf("Error: %v\n", physfs.GetError())
    } else {
        fmt.Printf("Mounted example.zip\n")
    }
    
    // if !physfs.Mount(path.Dir(exePath), "/", 1) {
    //     fmt.Printf("Error: %v\n", physfs.GetError())
    // } else {
    //     fmt.Printf("Mounted %s\n", path.Dir(exePath))
    // }

    content, _ := physfs.ReadFile("example.txt")
    if content != nil {
        fmt.Printf("File contents:\n%s\n", string(content))
    } else {
        fmt.Printf("Error: %v\n", physfs.GetError())
    }

    // Test OpenWrite and Close
    physfs.SetWriteDir("src")
    f := physfs.OpenWrite("example2.txt")
    if f == nil {
        fmt.Printf("Error: %v\n", physfs.GetError())
    } else {
        fmt.Printf("Opened example2.txt for writing\n")
        f.Write(([]byte)("Hello, world!\n"))
        physfs.Close(f)
    }

    // Test OpenWrite and Close
    physfs.SetWriteDir(".")
    f = physfs.OpenAppend("src/example3.txt")
    if f == nil {
        fmt.Printf("Error: %v\n", physfs.GetError())
    } else {
        fmt.Printf("Opened example2.txt for writing\n")
        f.Write(([]byte)("Hello, world!\n"))
        physfs.Close(f)
    }

    if !physfs.Unmount("example.zip") {
        fmt.Printf("Error: %v\n", physfs.GetError())
    } else {
        fmt.Printf("Unmounted example.zip\n")
    }

    // if !physfs.Unmount(path.Dir(exePath)) {
    //     fmt.Printf("Error: %v\n", physfs.GetError())
    // } else {
    //     fmt.Printf("Unmounted %s\n", path.Dir(exePath))
    // }
}