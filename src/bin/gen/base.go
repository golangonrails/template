/*
  Codegen, use `go run tasks/g.go`
*/
package gen

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

const postFix = ".go"
const seperateLine = "_"

func generateID() string {
	return time.Now().Format("20060102150405")
}

func generateFile(name, dir, id, content string) (string, error) {
	migrationFilePath := fmt.Sprintf("%v/%v%s%v%v", dir, id, seperateLine, strings.ToLower(name), postFix)
	return migrationFilePath, ioutil.WriteFile(migrationFilePath, []byte(content), 0644)
}
