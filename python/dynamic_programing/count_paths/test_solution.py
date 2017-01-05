import unittest
from solution import solution


class TestSolution(unittest.TestCase):
    def test_matrix_clean_2x2(self):
        matrix = [[0, 0], [0, 0]]
        start = (0, 0)
        finish = (1, 1)
        self.assertEqual(2, solution(matrix, start, finish))

    def test_matrix_not_clean_3x3(self):
        start = (0, 0)

        matrix = [[0, 0, 0], [0, 1, 0], [0, 0, 0]]
        finish = (2, 2)
        self.assertEqual(2, solution(matrix, start, finish))

        matrix = [[0, 0, 0], [0, 0, 0], [0, 1, 0]]
        finish = (2, 2)
        self.assertEqual(3, solution(matrix, start, finish))


if __name__ == '__main__':
    unittest.main()
