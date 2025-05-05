def greet(name):
    """Returns a personalized greeting."""
    if not name.strip():
        return "Hello, Stranger!"
    return f"Hello, {name.strip().title()}!"
