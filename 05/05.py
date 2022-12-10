#!/usr/bin/python3

def parse_instruction(line):
    parts = line.split(' ')
    return int(parts[1]), int(parts[3])-1, int(parts[5])-1


def run_simulation(file_name, model_name):
    stacks_raw = []
    instructions_raw = []

    with open(file_name) as inp:
        lines = list(map(str.rstrip, inp.readlines()))

    sepPos = lines.index('')

    stacks_raw = lines[:sepPos-1]
    instructions_raw = lines[sepPos+1:]
    stack_count = int(lines[sepPos-1][-1])

    stacks = [[] for i in range(stack_count)]

    for stack_line in reversed(stacks_raw):
        for s, stack in enumerate(stacks):
            pos = s * 4 + 1
            if pos >= len(stack_line):
                break
            if stack_line[pos] == ' ':
                continue
            stack.append(stack_line[pos])

    for instruction in instructions_raw:
        amount, from_, to_ = parse_instruction(instruction)

        to_move = stacks[from_][-amount:]
        if model_name == "9000":
            to_move.reverse()
        stacks[to_].extend(to_move)
        del(stacks[from_][-amount:])

    for stack in stacks:
        print(stack[-1],end="")
    print()


if __name__ == '__main__':
    run_simulation('input.txt', '9001')
