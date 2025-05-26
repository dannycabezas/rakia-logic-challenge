# Rakia Logic Challenge: Message Decoding

## Problem Description

You are given a message encoded using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26

Write a function or algorithm that takes a string of digits and returns the number of ways it
can be decoded back into its original message.

For example:
- Given the input "12", the possible decodings are "AB" and "L", so the output should be 2.
- For the input "226", the possible decodings are "BZ", "VF", and "BBF", making the output 3.
- With the input "0", there are no valid decodings, resulting in an output of 0.

Your solution should efficiently handle larger inputs as well.

## Solution

The provided solution is implemented in Go (`main.go`) using a dynamic programming approach to efficiently calculate the number of possible decodings.

The core logic resides in the `numDecodings` function, which takes a string of digits as input and returns the total number of ways it can be decoded. The function iterates through the input string from right to left, calculating the number of decodings possible at each position based on the decodings of subsequent positions. It considers both single-digit and two-digit interpretations, ensuring that two-digit numbers fall within the valid range (10-26).

The use of `*big.Int` ensures that the solution can handle very large numbers of decodings, which might occur with longer input strings.

## Tests

Unit tests for the `numDecodings` function are provided in the file `main_test.go`.  
You can run all tests usando el siguiente comando:

```bash
go test
```

## Examples

```
Input: "12"
Output: 2

Input: "226"
Output: 3

Input: "0"
Output: 0
```