def loop_sum(arr):
    total = 0
    for x in arr:
        total += x
    return total


if __name__ == '__main__':
    print(loop_sum([1, 2, 3, 4]))
