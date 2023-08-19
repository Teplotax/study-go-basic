package main

/*
A palindrome is a word, phrase, number, or other sequence of characters which reads the same backward or forward.

Given a string , print Yes if it is a palindrome, print No otherwise.

Constraints
The input will consist at most lower case english letters.

Sample Input: madam
Sample Output: Yes
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func isPalindrome(input string) bool {
	word := strings.ToLower(input)
	reverse := ""

	for i := len(word) - 1; i >= 0; i-- {
		reverse += string(word[i])
	}
	return strings.EqualFold(word, reverse)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	word := strings.TrimSpace(readLine(reader))

	if isPalindrome(word) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
