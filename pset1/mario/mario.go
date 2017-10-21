// Copyright 2017 Arjun Rao
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// getHeight prompts the user to enter an integer, and reprompts on invalid input.
// It returns a positive integer between 1-23.
func getHeight() int {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Height: ")
		input, _ := inputReader.ReadString('\n')
		height, err := strconv.Atoi(strings.TrimSpace(input))
		if err == nil && height > 0 && height < 24 {
			return height
		}
	}
}

func main() {
	// Get height from the user.
	height := getHeight()
	for level := height; level > 0; level-- {
		// Add correct number of spaces before blocks in a level to a string to be printed.
		levelString := fmt.Sprintf("%*s", level+1, "##")

		// Add remaining # blocks to the string.
		for j := level; j < height; j++ {
			levelString += "#"
		}

		// Print the prepared string for 1 level.
		fmt.Println(levelString)
	}
}
