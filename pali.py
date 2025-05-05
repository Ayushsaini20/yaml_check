def is_palindrome(n):
    return str(n) == str(n)[::-1]

# Example usage:
num = 121
print(f"{num} is palindrome:", is_palindrome(num))
