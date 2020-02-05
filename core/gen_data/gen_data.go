package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	. "github.com/candid82/joker/core"
	_ "github.com/candid82/joker/std/html"
	_ "github.com/candid82/joker/std/string"
)

const template string = `// Generated by gen_data. Don't modify manually!

{slowInit}package core

var {name}Data = []byte("{content}")
`

const hextable = "0123456789abcdef"

func main() {
	namespaces := map[string]struct{}{}

	GLOBAL_ENV.FindNamespace(MakeSymbol("user")).ReferAll(GLOBAL_ENV.CoreNamespace)
	for _, f := range CoreSourceFiles {
		GLOBAL_ENV.SetCurrentNamespace(GLOBAL_ENV.CoreNamespace)
		content, err := ioutil.ReadFile("data/" + f.Filename)
		if err != nil {
			panic(err)
		}
		content, err = PackReader(NewReader(bytes.NewReader(content), f.Name), "")
		PanicOnErr(err)

		dst := make([]byte, len(content)*4)
		for i, v := range content {
			dst[i*4] = '\\'
			dst[i*4+1] = 'x'
			dst[i*4+2] = hextable[v>>4]
			dst[i*4+3] = hextable[v&0x0f]
		}

		nsName := GLOBAL_ENV.CurrentNamespace().Name.Name()

		slowInit := ""
		if _, found := namespaces[nsName]; !found && nsName != "user" {
			slowInit = `// +build !fast_init

`
		}

		namespaces[nsName] = struct{}{}

		name := f.Filename[0 : len(f.Filename)-5] // assumes .joke extension
		fileContent := strings.ReplaceAll(template, "{name}", name)
		fileContent = strings.ReplaceAll(fileContent, "{slowInit}", slowInit)
		fileContent = strings.Replace(fileContent, "{content}", string(dst), 1)
		ioutil.WriteFile("a_"+name+"_data.go", []byte(fileContent), 0666)
	}

	const dataTemplate = `// Generated by gen_data. Don't modify manually!

package core

func init() {
	coreNamespaces = []string{
{coreNamespaces}
	}
}
`

	coreNamespaces := []string{}
	for ns, _ := range namespaces {
		coreNamespaces = append(coreNamespaces, fmt.Sprintf(`
		"%s",`[1:],
			ns))
	}
	dataContent := strings.Replace(dataTemplate, "{coreNamespaces}", strings.Join(coreNamespaces, "\n"), 1)
	ioutil.WriteFile("a_data.go", []byte(dataContent), 0666)
}
