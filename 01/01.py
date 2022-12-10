import sys

cals = []
summe = 0
for line in sys.stdin.readlines():
    line = line.rstrip('\n')
    if line == "":
        cals.append(summe)
        summe = 0
    else:
        summe += int(line)
cals.append(summe)

print(max(cals))

print(sum(sorted(cals)[-3:]))