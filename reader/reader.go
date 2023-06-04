package reader

import (
	"bufio"
	"os"
)

func BuildCalList(dir string) (error, []string) {
	file, err := os.Open(dir)
	if err != nil {
		return err, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cals []string
	for scanner.Scan() {
		cals = append(cals, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return err, nil
	}

	return nil, cals
}
