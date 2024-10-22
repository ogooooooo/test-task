package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Text struct {
	Content string
}

func (t *Text) textModifier() {
	
	t.Content = strings.Join(strings.Fields(t.Content), " ")

	t.Content = processMinus(t.Content)

	t.Content = strings.ReplaceAll(t.Content, "+", "!")

	sum := sumAndRemoveDigits(&t.Content)

	if sum > 0 {
		t.Content += fmt.Sprintf(" %d", sum)
	}

	fmt.Println(t.Content)
}

func processMinus(content string) string {
	re := regexp.MustCompile(`(\S)(-)(\S)`)
	for re.MatchString(content) {
		content = re.ReplaceAllString(content, "$3$1")
	}
	return strings.ReplaceAll(content, "-", "")
}

func sumAndRemoveDigits(content *string) int {
	sum := 0
	re := regexp.MustCompile(`[0-9]`)
	digits := re.FindAllString(*content, -1)

	for _, digit := range digits {
		num, _ := strconv.Atoi(digit)
		sum += num
	}

	*content = re.ReplaceAllString(*content, "")
	return sum
}

func main() {
	text := &Text{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите строку:")

	for scanner.Scan() {
		text.Content = scanner.Text()
		text.textModifier()
		break 
	}
}
