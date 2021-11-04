package service

import (
	"log"
	"os"
    "strings"
)

// GenerateCode generates the code for the Sphero
func GenerateCode(programs []string) error {
    fileContent, err := getTemplate()
    if err != nil {
        log.Println("error:file\n",err)
        return err
    }

    writeCode(programs, fileContent)

    return nil
}

func writeCode(programs []string, content []byte) error {
    appendContent := createAppendContent(programs)

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

func createAppendContent(programs []string) string {
    var appendContent string
    for _, program := range programs {
        switch program {
            case "前進":
                appendContent += "    " + string(GO_FORWARD) + "\n"
            case "後進":
                appendContent += "    " + string(GO_BACKWARD) + "\n"
        }
    }

    return appendContent
}

func getTemplate() ([]byte, error) {
    fileContent, err := os.ReadFile("./lib/sphero_template.ts")
    if err != nil {
        return nil, err
    }

    return fileContent, nil
}
