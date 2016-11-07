"""
Example of using iter and partial.
TODO: extend the example
"""
from functools import partial

SIZE = 1024


def read_from_file_version_1(path):
    with open(path) as f:
        block = []
        while True:
            block = f.read(SIZE)
            if block == "":
                break
            block.append(block)


def read_from_file_version_2(path):
    with open(path) as f:
        blocks = []
        for block in iter(partial(f.read, SIZE), ''):
            blocks.append(block)
