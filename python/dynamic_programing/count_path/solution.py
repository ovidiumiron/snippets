#NxN grid
#You can move down or the right
#There is some cel blocked in the grid.
#Provide a function to calculate how many path there is from start to finish.
def solution(matrix, start, finish):
    memo = {
        (finish[0] - 1, finish[1]): 1,
        (finish[0], finish[1] - 1): 1
    }

    def s(matrix, memo, st, finish):
        if st in memo:
            return memo[st]

        tmp = 0

        next_row = st[0] + 1
        if next_row <= finish[0] and matrix[next_row][st[1]] == 0:
            tmp += s(matrix, memo, (next_row, st[1]), finish)

        next_column = st[1] + 1
        if next_column <= finish[1] and matrix[st[0]][next_column] == 0:
            tmp += s(matrix, memo, (st[0], next_column), finish)

        memo[st] = tmp
        return memo[st]

    return s(matrix, memo, start, finish)
