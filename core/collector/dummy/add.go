package dummy

import (
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func StartDummyData(filePath string, until int) {
	f, err := os.Create(filePath + "/data")
	check(err)
	defer f.Close()

	sum := 0
	for until == -1 || sum < until {
		time.Sleep(10 * time.Millisecond)
		_, err := f.WriteString("writes wefwfew32rf23f2fwefwef23fwdsfwfwsfwe23fsdfwefwefwefwe23ffwefe\n")
		check(err)
		sum++
	}
}

func Clear(filePath string) {
	os.Remove(filePath)
}
