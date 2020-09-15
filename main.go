package main

import (
	"fmt"
	jsonpatch "github.com/evanphx/json-patch"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	keyEnvPatchSource = "PATCH_SOURCE"
	keyEnvPatchTarget = "PATCH_TARGET"
)

var (
	optPatchSource = os.Getenv(keyEnvPatchSource)
	optPatchTarget = os.Getenv(keyEnvPatchTarget)
)

func exit(err *error) {
	if *err != nil {
		log.Println("exited with error:", (*err).Error())
		os.Exit(1)
	} else {
		log.Println("exited")
	}
}

func main() {
	var err error
	defer exit(&err)

	optPatchSource = strings.TrimSpace(optPatchSource)
	if optPatchSource == "" {
		err = fmt.Errorf("missing environment variable: %s", keyEnvPatchSource)
		return
	}

	optPatchTarget = strings.TrimSpace(optPatchTarget)
	if optPatchTarget == "" {
		err = fmt.Errorf("missing environment variable: %s", keyEnvPatchTarget)
		return
	}

	var rawSource []byte
	if rawSource, err = ioutil.ReadFile(optPatchSource); err != nil {
		return
	}

	log.Printf("source: %s", string(rawSource))

	var rawTarget []byte
	if rawTarget, err = ioutil.ReadFile(optPatchTarget); err != nil {
		return
	}

	log.Printf("target: %s", string(rawTarget))

	var p jsonpatch.Patch
	if p, err = jsonpatch.DecodePatch(rawSource); err != nil {
		return
	}

	var result []byte
	if result, err = p.ApplyIndent(rawTarget, "  "); err != nil {
		return
	}

	log.Printf("result: %s", string(result))

	if err = ioutil.WriteFile(optPatchTarget, result, 0644); err != nil {
		return
	}
}
