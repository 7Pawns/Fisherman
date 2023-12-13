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
	scriptPath string
}

func NewBait(scriptPath string) *Bait {
	bait := new(Bait)
	bait.scriptPath = scriptPath
	return bait
}

func (bait Bait) GetBaitScript(params ...any) string {
	script, err := os.ReadFile(bait.scriptPath)
	checkError(err)

	scriptFmt := string(script)
	return fmt.Sprintf(scriptFmt, params...)
}
