import unittest
from greetings import greet

class TestGreetFunction(unittest.TestCase):

    def test_regular_name(self):
        self.assertEqual(greet("alice"), "Hello, Alice!")

    def test_name_with_spaces(self):
        self.assertEqual(greet("  bob  "), "Hello, Bob!")

    def test_empty_string(self):
        self.assertEqual(greet(""), "Hello, Stranger!")

    def test_all_uppercase(self):
        self.assertEqual(greet("JANE"), "Hello, Jane!")

if __name__ == '__main__':
    unittest.main()
