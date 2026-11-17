Computational Thinking & Programming Algorithms
A collection of Go (Golang) implementations of the six exercises from
the Asynchronous Assignment: Iterative (FOR Loop) & Recursive Control
Flow with IF-ELSE.
Note: The assignment's official submission rules require
standardized pseudocode and handwritten answers. These Go programs are
for practice and learning; they are not a replacement for the required
pseudocode submission.

Student Parameters
- NIM: 103032640006
- K: 6 + 4 = 10
- M: 0 + 3 = 3
Programs
  File                    Problem                 Description
  problem1_1.go         1.1 -- SeriesMystery    Uses a for loop and
an if-else condition
to update a running
total. Uses K = 10.

  problem1_2.go         1.2 -- ComputePattern   Demonstrates recursion
with base case,
odd-number, and
even-number branches.
Uses M = 3.

  problem2_1.go         2.1 -- RoasterMonitor   Reads temperature
measurements, counts
ideal and overheated
readings, and
calculates the average.

  problem2_2.go         2.2 -- CustomSequence   Prints a modified
FizzBuzz sequence using
ordered conditional
checks.

  problem3_1.go         3.1 -- CountEvenDigits  Recursively counts even
digits using arithmetic
operations.
problem3_2.go         3.2 -- FastPower        Calculates a
non-negative integer
power using recursion
and exponentiation by
squaring.
Requirements
- Go installed on your computer.
- A terminal or command prompt.
Check that Go is installed:
go version
How to Run
Each example is a separate Go program and contains its own main()
function. Save each program in its own .go file, then run it from the
terminal.

For example:
go run problem1_1.go
Run another exercise by changing the filename:
go run problem1_2.go
go run problem2_1.go
go run problem2_2.go
go run problem3_1.go
go run problem3_2.go
Do not combine all six files into the same Go package directory and run
go run . without first organizing them into separate folders or
removing the duplicate main() functions.
Expected Results
  Program                 Example                     Expected result
  SeriesMystery           K = 10                    Final Total = 35

  ComputePattern          ComputePattern(3)         5

  RoasterMonitor          190, 200, 215, 210, 185   Ideal: 2; overheated:
1; average: 200

  CustomSequence          N = 15                    COMPUTATION is
printed for 15

  CountEvenDigits         CountEvenDigits(48271)    3

  CountEvenDigits         CountEvenDigits(1359)     0
FastPower               FastPower(2, 6)           64
Concepts Practiced
- Variables and data types (int, float64)
- for loops
- if, else if, and else
- Arithmetic operators and remainder (%)
- Integer division
- User input and output with fmt.Scan and fmt.Println
- Recursive functions and base cases
- Exponentiation by squaring
Important Logic Notes
- In CustomSequence, check divisibility by both 3 and 5 before
  checking divisibility by 3 alone. Otherwise, multiples of 15 (such
  as 15 and 30) will print the wrong label.
- In CountEvenDigits, % 10 extracts the final digit and integer
  division by 10 removes it. The implementation should also handle 0
  as an even digit.
- In FastPower, an even exponent is halved, and the resulting
  half-power is squared. This reduces the number of recursive steps to
  logarithmic growth for even-exponent decomposition.
