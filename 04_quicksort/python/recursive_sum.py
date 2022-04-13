def recursive_sum(my_list):
    if not my_list:
        return 0
    return my_list[0] + recursive_sum(my_list[1:])


if __name__ == '__main__':
    print(recursive_sum([1, 2, 3, 4]))
