score = {
    'A': 1,
    'B': 2,
    'C': 3,
    'X': 1,
    'Y': 2,
    'Z': 3,
}

def play(op, me):
    iop = score[op]
    ime = score[me]
    if iop == ime:
        return 3
    if iop == 1 and ime == 3:
        return 0
    if iop == 3 and ime == 1:
        return 6
    return 6 if ime > iop else 0

total = 0
with open('input.txt') as inp:
    for line in inp.readlines():
        op, me = line.rstrip().split(' ')
        total += score[me] + play(op, me)

print(total)