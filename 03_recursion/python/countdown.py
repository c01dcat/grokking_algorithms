def countdown(i):
    if i <= 0:
        return
    else:
        print(i)
        countdown(i - 1)


if __name__ == '__main__':
    countdown(5)
