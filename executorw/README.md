# Go Code Executor

This program is a Go code executor that allows users to execute Go code provided by the user and display the result in a web interface. It is built using the Gin framework and the Executor package.

Note that this approach has security implications, as you are executing user-supplied code at runtime. Be sure to properly validate and sanitize any code input to avoid vulnerabilities.

If your goal is to execute Go code dynamically at runtime without generating temporary files, you may consider using an interpreter or an in-memory Go execution environment. For example, you can use the [Yaegi]


## Requirements

- Go 1.16 or higher

## Installation

1. Clone this repository to your working directory:

   ```shell
   git clone https://github.com/alexfu00/executorw.git
   ```
2. Navigate to the cloned repository:

   ```shell
   cd your_repository
   ```
3. Install the dependencies using the go get command:

   ```shell
   go get github.com/gin-gonic/gin   
   ```

   or

   ```shell
   cd executorw
   go mod tidy   
   ```

## Usage

  1. Run the program using the following command:

     ```shell
     go run main.go
     ```
  2. Open your browser and navigate to http://localhost:8080

  3. Enter your Go code in the text area and click the "Execute" button to execute the code.

        For example, you can enter the following code:

     ```go
     package main

     import "fmt"

     func main() {
        num1 := 10
        num2 := 20
        sum := num1 + num2
        fmt.Println("The result is:", sum)
     }
     ```

  4. The result of the execution will be displayed in the "Output" section.
  
