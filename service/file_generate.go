package service

import (
	"log"
	"os"
    "strings"
)

// GenerateCode generates the code for the Sphero
func GenerateCode() error {
    fileContent, err := getTemplate()
    if err != nil {
        log.Println("error:file\n",err)
        return err
    }

    writeCode(fileContent)

    return nil
}

func writeCode(content []byte) error {
    var appendContent string
    for i := 0; i < 3; i++ {
        appendContent += "    " + string(GO_FORWARD) + "\n"
    }

    c := strings.Replace(string(content), "// appendContent", appendContent, -1)
    f, err := os.Create("sphero.ts")
    if err != nil {
        return err
    }

    _, err = f.Write([]byte(c))
    if err != nil {
        return err
    }

    return nil
}

func getTemplate() ([]byte, error) {
    fileContent, err := os.ReadFile("./lib/sphero_template.ts")
    if err != nil {
        return nil, err
    }

    return fileContent, nil
}
