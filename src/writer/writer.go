package writer

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Process struct {
	Pid      int
	Priority int
	Cycles   int
}

func writeToFile(file string, processes []Process) {
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Write([]string{"PID", "Priority", "Cycles"})
	for _, p := range processes {
		strPid := strconv.Itoa(p.Pid)
		strPriority := strconv.Itoa(p.Priority)
		strCycles := strconv.Itoa(p.Cycles)

		var arr = []string{strPid, strPriority, strCycles}
		err = w.Write(arr)
		if err != nil {
			panic(err)
		}
	}
	w.Flush()
}
