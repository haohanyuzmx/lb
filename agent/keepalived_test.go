package agent

import (
	"testing"
)

func TestKeepalivedConf(t *testing.T) {
	KeepalivedConf().ToConf("")
}
