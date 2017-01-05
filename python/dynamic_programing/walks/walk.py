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
        for coordonate in left_up_coordinates(idx, dimension):
            yield coordonate


def left_up_coordinates(idx, dimension):
    """ Generate coordonates for traversing left and up. """
    coordinates = list()
    left = dimension - idx
    up = dimension
    while left != up:
        coordinates.append((left, up))
        coordinates.append((up, left))
        up = up - 1

    coordinates.append((left, up))

    return coordinates

if __name__ == "__main__":
    import doctest
    doctest.testmod()
