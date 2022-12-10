#!/usr/bin/python3

import sys

def is_unique(chars):
    return len(set(chars)) == len(chars)


def get_marker(inp, length):
    for pos in range(len(inp)-length):
        if is_unique(inp[pos:pos+length]):
            return pos+length


if __name__ == '__main__':
    for inp in sys.stdin.readlines():
        packet = get_marker(inp, 4)
        message = get_marker(inp, 14)
        print(f'Packet: {packet} Message: {message}')