package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
        "strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())


	result := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		teacherName := scanner.Text()

		scanner.Scan()
		studentScores := strings.Fields(scanner.Text())

		sum := 0
		for _, score := range studentScores {
				scoreInt, _ := strconv.Atoi(score)
				sum += scoreInt
		}

		avg := float64(sum) / float64(len(studentScores))

		teacherEvaluation := ""
		if avg >= 80 {
				teacherEvaluation = "Excellent"
		} else if avg >= 60 {
				teacherEvaluation = "Very Good"
		} else if avg >= 40 {
				teacherEvaluation = "Good"
		} else {
				teacherEvaluation = "Fair"
		}

		result[i] = fmt.Sprintf("%s %s", teacherName, teacherEvaluation)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s\n", result[i])
	}
}