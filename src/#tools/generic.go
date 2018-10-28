package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	namedArgs, _ := ParseArgs()

	if _, ok := namedArgs["file"]; !ok {
		log.Fatal("ERROR: Must specify file name")
	}
	if _, ok := namedArgs["types"]; !ok {
		log.Fatal("ERROR: Must specify types")
	}
	if _, ok := namedArgs["dir"]; !ok {
		namedArgs["dir"] = "."
	}

	replacements := make(map[string]string)
	for _, v := range strings.Split(namedArgs["types"], ",") {
		arr := strings.Split(v, "=")
		if len(arr) != 2 {
			log.Fatalf("ERROR: Incorrect types format: %s", v)
		}
		replacements[arr[0]] = arr[1]
	}

	filePath, found := FindFile(namedArgs["dir"], namedArgs["file"])
	if !found {
		log.Fatal("ERROR: File not found")
	}

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("ERROR: Reading file '%s': %v", filePath, err)
	}

	code := string(buf)
	for k, v := range replacements {
		code = strings.Replace(code, k, v, -1)
	}

	if outputFile := namedArgs["output"]; outputFile != "" {
		if err := os.MkdirAll(path.Dir(outputFile), 0); err != nil {
			log.Fatalf("ERROR: Creating directory for file '%s': %v", outputFile, err)
		}
		if err := ioutil.WriteFile(outputFile, []byte(code), 0); err != nil {
			log.Fatalf("ERROR: Writing file '%s': %v", outputFile, err)
		}
	} else {
		fmt.Println(code)
	}
}

func ParseArgs() (map[string]string, map[string]bool) {
	namedArgs := make(map[string]string)
	otherArgs := make(map[string]bool)
	for _, v := range os.Args[1:] {
		idx := strings.IndexRune(v, '=')
		if idx < 0 {
			otherArgs[v] = true
		} else {
			flag := v[:idx]
			if len(flag) > 3 && flag[:2] == "--" && flag[2] != '-' {
				namedArgs[flag[2:]] = v[idx+1:]
			}
		}
	}
	return namedArgs, otherArgs
}

func FindFile(dir, file string) (string, bool) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", false
	}
	for _, fi := range infos {
		baseName := fi.Name()
		if fi.IsDir() {
			if filePath, found := FindFile(path.Join(dir, baseName), file); found {
				return filePath, true
			}
		} else {
			if baseName == file {
				return path.Join(dir, baseName), true
			}
		}
	}
	return "", false
}
