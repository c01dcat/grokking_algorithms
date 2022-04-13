def recursive_max(my_list):
    if not my_list:
        return None
    if len(my_list) == 1:
        return my_list[0]
    else:
        list_max = recursive_max(my_list[1:])
        return my_list[0] if my_list[0] > list_max else list_max


if __name__ == '__main__':
    print(recursive_max([1, 2, 3, 4]))
