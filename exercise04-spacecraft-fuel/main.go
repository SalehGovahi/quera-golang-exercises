package main

import (
        "bufio"
        // "fmt"
        "os"
        "strconv"
        "strings"
)

func main() {
        // Read the number of names
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        numNames, _ := strconv.Atoi(scanner.Text())

        // Read the names and their corresponding numbers
        names := make([]string, numNames)
        nums := make([][]int, numNames)

        for i := 0; i < numNames; i++ {
			scanner.Scan()
			line := scanner.Text()
			parts := strings.Split(line, " ")

			names[i] = parts[0]
			nums[i] = make([]int, len(parts)-1)

			for j := 1; j < len(parts); j++ {
				num, _ := strconv.Atoi(parts[j])
				nums[i][j-1] = num
			}

			for m := 1; m < len(parts); m++ {
				num, _ := strconv.Atoi(parts[m])
				nums[i][m-1] = num
			}
        }
}