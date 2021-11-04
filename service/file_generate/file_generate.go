package service

import (
	"os"
	"log"
    "io"
)

// GenerateCode generates the code for the Sphero
func GenerateCode() (*os.File, error) {
    file, err := copyFile()
    if err != nil {
        log.Println("error:file\n",err)
        return nil, err
    }

    return file, nil
}

func copyFile() (*os.File, error) {
    in, err := os.Open("./lib/sphero_template.ts")
    if err != nil {
        return nil, err
    }
    defer in.Close()

    out, err := os.Create("./lib/sphero.ts")
    if err != nil {
        return nil, err
    }
    defer out.Close()

    _, err = io.Copy(out, in)
    if err != nil {
        return nil, err
    }

    return out, nil
}