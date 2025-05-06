def factorial_recursive(n):
    if n == 0 or n == 1:
        return 1
    return n * factorial_recursive(n - 1)

def factorial_iterative(n):
    result = 1
    for i in range(2, n + 1):
        result *= i
    return result

# Main program
num = int(input("Enter a non-negative integer: "))

if num < 0:
    print(" Factorial is not defined for negative numbers.")
else:
    print(f"Recursive: Factorial of {num} is {factorial_recursive(num)}")
    print(f"Iterative: Factorial of {num} is {factorial_iterative(num)}")
