import unittest
from solution import my_func
from random import randint
import time
from functools import reduce


class TestCarlos(unittest.TestCase):
    def test_my_func(self):
        input_list = list()
        for i in range(70):
            input_list.append(float(randint(1, 9)))
        product = reduce(lambda x, y: x*y, input_list)
        expected_list = [product/x for x in input_list]
        startTime = time.time()
        self.assertEqual(expected_list, my_func(input_list))

        print("needs {0} seconds".format(time.time() - startTime))


if __name__ == '__main__':
    unittest.main()
