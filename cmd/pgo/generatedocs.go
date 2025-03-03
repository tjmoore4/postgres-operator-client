// Copyright 2021 - 2022 Crunchy Data Solutions, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra/doc"

	"github.com/crunchydata/postgres-operator-client/internal/cmd"
)

const fmTemplate = `---
title: "%s"
---
`

func main() {

	fmt.Println("generate CLI markdown")

	filePrepender := func(filename string) string {
		name := filepath.Base(filename)
		base := strings.TrimSuffix(name, path.Ext(name))
		fmt.Println(base)
		return fmt.Sprintf(fmTemplate, strings.ReplaceAll(base, "_", " "))
	}

	linkHandler := func(name string) string {
		base := strings.TrimSuffix(name, path.Ext(name))
		return "/reference/" + strings.ToLower(base) + "/"
	}

	pgo := cmd.NewPGOCommand(os.Stdin, os.Stdout, os.Stderr)
	err := doc.GenMarkdownTreeCustom(pgo, "./", filePrepender, linkHandler)
	if err != nil {
		log.Fatal(err)
	}
}
