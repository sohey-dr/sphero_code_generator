package service

import (
	"os"
	"log"
)

type Element string

// GenerateCode generates the code for the Sphero
func GenerateCode(fileName string) (*os.File, error) {
    file, err := os.Open("./lib/" + fileName)
    if err != nil {
        log.Println("error:file\n",err)
        return nil, err
    }

    return file, nil
}
