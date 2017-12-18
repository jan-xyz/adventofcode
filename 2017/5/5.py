#!/usr/bin/env python


if __name__ == '__main__':
    with open('./puzzle_input') as file:
        puzzle_input = [int(line) for line in file.read().splitlines()]

    inside = True
    index = 0
    steps = 0
    jump = 0
    while inside:
        print("at {0} of {1} with next jump being {2}".format(index, len(puzzle_input), jump))
        if index >= len(puzzle_input) or index < 0:
            inside = False
        else:
            jump = puzzle_input[index]
            puzzle_input[index] += 1
            index = index + jump
            steps += 1

    print("steps : {0}".format(steps))

