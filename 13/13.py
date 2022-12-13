import json
import functools


def checkOrder(first_list, second_list):
    i = -1
    while True:
        i += 1

        if i == len(first_list) and i == len(second_list):
            return None
        if i == len(first_list):
            return True
        if i == len(second_list):
            return False

        first = first_list[i]
        second = second_list[i]

        if isinstance(first, int) and isinstance(second, int):
            if first == second:
                continue
            return first < second

        if isinstance(first, int):
            first = [first]
        if isinstance(second, int):
            second = [second]

        result = checkOrder(first, second)
        if result is None:
            continue
        return result


def cmp_packets(first, second):
    result = checkOrder(first, second)
    if result is None:
        return 0
    return result and -1 or 1


if __name__ == '__main__':
    packets = [
        json.loads(line.strip())
        for line in open('input.txt').readlines()
        if line.strip() != '']

    packet_sum = 0
    for i in range(0, int(len(packets)/2)):
        result = checkOrder(packets[i*2], packets[i*2+1])
        print(f'Packet {i+1}: {result}')
        if result:
            packet_sum += i + 1
    print(f'Packet sum: {packet_sum}')

    packets.extend([[[2]],[[6]]])
    packets.sort(key=functools.cmp_to_key(cmp_packets))

    start = packets.index([[2]]) + 1
    end = packets.index([[6]]) + 1

    print(f'Deocoder key: {start*end}')
