import sys


def main():

    path = sys.argv[1]
    f = open(path, 'r')

    str = f.readline()
    print(str)

    f.close()


if __name__ == "__main__":
    main()
