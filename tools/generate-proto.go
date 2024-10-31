package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
    rootDir, err := os.Getwd()
    if err != nil {
        fmt.Printf("Failed to get working directory: %v\n", err)
        os.Exit(1)
    }

    rootDir = filepath.ToSlash(rootDir)
    protoDir := filepath.ToSlash(filepath.Join(rootDir, "proto"))

    fmt.Printf("Root directory: %s\n", rootDir)
    fmt.Printf("Proto directory: %s\n", protoDir)

    var protoFiles []string
    err = filepath.Walk(protoDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && strings.HasSuffix(path, ".proto") {
            protoFiles = append(protoFiles, filepath.ToSlash(path))
        }
        return nil
    })
    if err != nil {
        fmt.Printf("Failed to walk proto directory: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("\nFound proto files:")
    for _, file := range protoFiles {
        fmt.Println(file)
    }

    for _, protoFile := range protoFiles {
        dir := filepath.ToSlash(filepath.Dir(protoFile))
        
        if err := os.MkdirAll(dir, 0755); err != nil {
            fmt.Printf("Failed to create directory %s: %v\n", dir, err)
            continue
        }

        fmt.Printf("\nCompiling: %s\n", protoFile)
        
        args := []string{
            "--proto_path=" + filepath.ToSlash(rootDir),
            "--go_out=" + filepath.ToSlash(rootDir),
            "--go_opt=paths=source_relative",
            "--go-grpc_out=" + filepath.ToSlash(rootDir),
            "--go-grpc_opt=paths=source_relative",
        }

        relPath, err := filepath.Rel(rootDir, protoFile)
        if err != nil {
            fmt.Printf("Failed to get relative path: %v\n", err)
            continue
        }
        args = append(args, filepath.ToSlash(relPath))

        cmd := exec.Command("protoc", args...)
        
        fmt.Println("Executing command:", "protoc", strings.Join(args, " "))

        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        cmd.Dir = rootDir

        if err := cmd.Run(); err != nil {
            fmt.Printf("Failed to compile %s: %v\n", protoFile, err)
            continue 
        }
    }

    fmt.Println("\nProto compilation process completed!")
}