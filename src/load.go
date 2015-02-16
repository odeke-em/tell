package tell

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

const SamplesDirPath = "samples"

func sampleFile(suffix string) string {
	fmt.Println("not yet implemented", suffix)
	return filepath.Join(SamplesDirPath, suffix)
}
