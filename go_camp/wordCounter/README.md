# Program Execution

$ echo "word1 word2 word3 word4" | ./wordCounter
4

$ echo -e "word1 word2 \n word3 word4" | ./wordCounter -l
2

$ go test 
PASS
ok      wordCounter     0.001s

$ go test  -v
=== RUN   TestCountWords
--- PASS: TestCountWords (0.00s)
PASS
ok      wordCounter     0.002s


# Exercices
- update the program to split the stream by bytes
    - add a new input flag `-b` , when it is true, program should count based on number of bytes.
- Finish the test cases lines and bytes. 
