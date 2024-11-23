# Concurrent Prime Number Checker

## Problem Description
Write a function that determines whether a given number is prime or not using concurrent operations. A prime number is a natural number greater than 1 that is only divisible by 1 and itself.

### Input
- An integer `n` where 1 ≤ n ≤ 10^9

### Output
- Return `true` if the number is prime
- Return `false` if the number is not prime

### Examples
Input: 17
Output: true
Explanation: 17 is only divisible by 1 and itself

Input: 4
Output: false
Explanation: 4 is divisible by 1, 2, and 4

### Notes
- Implement the solution using goroutines to check divisibility concurrently
- Use channels for communication between goroutines
- Consider implementing a worker pool pattern for better performance 