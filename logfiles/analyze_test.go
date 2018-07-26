package logfiles

import (
	"fmt"
	"testing"
)

func TestConnectionStatusDetailLine(t *testing.T) {
	status, _ := ConnectionStatusLine("Jun 20 23:28:39.298283 RL00122 rlc[694]: RLM RL00122 Connection Status")
	if status {
		fmt.Printf("Detected Connection Status Line")
		_, idx, val := ConnectionStatusDetailLine("                                         Mem Used(rlc):  48664     ")
		fmt.Printf("Mem used rlc idx %d val %s\n", idx, val)
	}

}
