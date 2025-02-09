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

    zipName := "mk1.zip"

    if !physfs.Mount(zipName, "/", 1) {
        fmt.Printf("Error: %v\n", physfs.GetError())
    } else {
        fmt.Printf("Mounted %v [OK]\n", zipName)
    }
    
    // if !physfs.Mount(path.Dir(exePath), "/", 1) {
    //     fmt.Printf("Error: %v\n", physfs.GetError())
    // } else {
    //     fmt.Printf("Mounted %s\n", path.Dir(exePath))
    // }

    // content, _ := physfs.ReadFile("example.txt")
    // if content != nil {
    //     fmt.Printf("File contents:\n%s\n", string(content))
    // } else {
    //     fmt.Printf("Error: %v\n", physfs.GetError())
    // }

    // Test OpenWrite and Close
    // physfs.SetWriteDir("src")
    // f := physfs.OpenWrite("example2.txt")
    // if f == nil {
    //     fmt.Printf("Error: %v\n", physfs.GetError())
    // } else {
    //     fmt.Printf("Opened example2.txt for writing\n")
    //     f.Write(([]byte)("Hello, world!\n"))
    //     physfs.Close(f)
    // }

    // Test OpenWrite and Close
    // physfs.SetWriteDir(".")
    // f = physfs.OpenAppend("src/example3.txt")
    // if f == nil {
    //     fmt.Printf("Error: %v\n", physfs.GetError())
    // } else {
    //     fmt.Printf("Opened example2.txt for writing\n")
    //     f.Write(([]byte)("Hello, world!\n"))
    //     physfs.Close(f)
    // }

    // Test Walk directory
    // test_dir := ""
    // if len(os.Args) == 2 {
    //     test_dir = os.Args[1]
    // } else {
    //     test_dir = "stages"
    // }
    // physfs.Walk(test_dir, func(path string, isDir bool) error {
	// 	if isDir {
	// 		fmt.Println("[DIR]  ", path)
	// 	} else {
	// 		fmt.Println("[FILE] ", path)
	// 	}
	// 	return nil // Continue traversal
	// })

    // Test FindFile
    // test_dir := ""
    // test_file := ""
    // if len(os.Args) == 2 {
    //     test_dir = "data"
    //     test_file = os.Args[1]
    // } else if len(os.Args) == 3 {
    //     test_dir = os.Args[1]
    //     test_file = os.Args[2]
    // } else {
    //     test_dir = "data"
    //     test_file = "system.def"
    // }
    // fmt.Printf("Find file [%v] in [%v] %v\n", test_file, test_dir, physfs.GetDirSeparator())
    // validFilePath := physfs.FindFile(test_dir, test_file)
    // if validFilePath == "" {
    //     fmt.Printf("Cannot find %v in %v\n", test_file, test_dir)
    // } else {
    //     fmt.Printf("Found %v in %v\n", test_file, validFilePath)
    // }

    // Test FindFile
    // test_file := ""
    // if len(os.Args) == 2 {
    //     test_file = os.Args[1]
    // } else {
    //     test_file = "data/system.def"
    // }
    // validFilePath := physfs.CheckFile(test_file)
    // if validFilePath == "" {
    //     fmt.Printf("FAIL\n")
    // } else {
    //     fmt.Printf("%v found in %v\n", test_file, validFilePath)
    // }

    // Test FindFileExt
    // test_file := ""
    // test_dirs := []string{}
    // if len(os.Args) >= 3 {
    //     for i := 1; i < len(os.Args)-1; i++ {
    //         test_dirs = append(test_dirs, os.Args[i])
    //     }
    //     test_file = os.Args[len(os.Args) - 1]
    // } else {
    //     test_file = "system.def"
    //     test_dirs = append(test_dirs, "data")
    // }
    // validFilePath := physfs.FindFileExt(test_dirs, test_file)
    // if validFilePath == "" {
    //     fmt.Printf("FAIL\n")
    // } else {
    //     fmt.Printf("%v found in %v\n", test_file, validFilePath)
    // }

    test_dir := ""
    pattern := ""
    if len(os.Args) == 3 {
        test_dir = os.Args[1]
        pattern = os.Args[2]
    } else {
        test_dir = "data"
        pattern = "*.def"
    }
    validFilesPath := physfs.FindFileMatch(test_dir, pattern)
	if len(validFilesPath) == 0 {
		fmt.Printf("FAIL\n")
    } else {
        fmt.Printf("Find Files in [%v] with pattern '%v'\n", test_dir, pattern)
		for k, v := range validFilesPath {
			fmt.Printf("%v: %v\n", k, v)
		}
	}

    if !physfs.Unmount(zipName) {
        fmt.Printf("Error: %v\n", physfs.GetError())
    } else {
        fmt.Printf("Unmounted %v [OK]\n", zipName)
    }

    // if !physfs.Unmount(path.Dir(exePath)) {
    //     fmt.Printf("Error: %v\n", physfs.GetError())
    // } else {
    //     fmt.Printf("Unmounted %s\n", path.Dir(exePath))
    // }
}