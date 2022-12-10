score = {
    'A': 1,
    'B': 2,
    'C': 3,
    'X': 0,
    'Y': 3,
    'Z': 6,
}

def play(op, res):
    iop = score[op]
    ires = score[res]
    if ires == 3:
        return ires + iop
    if ires == 6:
        return ires + iop % 3 + 1
    return ires + (iop + 1) % 3 + 1

total = 0
with open('input.txt') as inp:
    for line in inp.readlines():
        op, res = line.rstrip().split(' ')
        total += play(op, res)

print(total)