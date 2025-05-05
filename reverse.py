def reverse_string(text):
    return text[::-1]

# Main program
user_input = input("Enter a string: ")
reversed_text = reverse_string(user_input)

print(f"Reversed string: {reversed_text}")
