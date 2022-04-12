def binary_search(bucket, item):
    low = 0
    high = len(bucket) - 1

    while low <= high:
        mid = (low + high) // 2  # 取整除，返回商的整数部分，向下取整
        guess = bucket[mid]
        if guess == item:
            return mid
        if guess > item:
            high = mid - 1
        else:
            low = mid + 1
    return None


if __name__ == '__main__':
    my_bucket = [1, 3, 5, 7, 9]
    print(binary_search(my_bucket, 3))
    print(binary_search(my_bucket, -1))
