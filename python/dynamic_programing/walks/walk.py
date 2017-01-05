"""
Different ways to traverse a matrix.
"""


def walk_left_and_up(dimension):
    """
    Direction of walking is left, then up.
    Example:
        1 2 3
        4 5 6
        7 8 9
    left and up means: 9, 8, 6, 5, 7, 3, 4, 2, 1

    >>> list(walk_left_and_up(2))
    [(2, 2), (1, 2), (2, 1), (1, 1), (0, 2), (2, 0), (0, 1), (1, 0), (0, 0)]
    """
    for idx in range(dimension + 1):
        for coordonate in left_up(idx, dimension):
            yield coordonate


def left_up(idx, dimension):
    """ Generate coordonates for traversing left and up. """
    left = dimension - idx
    up = dimension
    while left != up:
        yield (left, up)
        yield (up, left)
        up = up - 1
    yield (left, up)

if __name__ == "__main__":
    import doctest
    doctest.testmod()
