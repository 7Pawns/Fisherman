package baits

import (
	"fmt"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

type Bait struct {
	path string
}

func NewBait(scriptPath string) *Bait {
	bait := new(Bait)
	bait.path = scriptPath
	return bait
}

func (bait Bait) BaitToString(params ...any) string {
	script, err := os.ReadFile(bait.path)
	checkError(err)

	scriptFmt := string(script)
	return fmt.Sprintf(scriptFmt, params...)
}
