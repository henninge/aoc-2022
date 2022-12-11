import 'dart:io';
import 'dart:math';

class Monkey {
  int pos = 0;
  List<BigInt> items = List.empty(growable: true);
  String operation = "";
  BigInt opParam = BigInt.zero;
  BigInt worryDivider = BigInt.from(3);
  BigInt testDivider = BigInt.zero;
  int targetTrue = 0;
  int targetFalse = 0;
  int inspected = 0;

  Monkey.fromLines(List<String> lines) {
    pos = int.parse(lines[0].substring(7, 8));
    items = List.of(lines[1]
        .substring(18)
        .split(', ')
        .map((str) => BigInt.from(int.parse(str))));

    var op = lines[2].substring(23, 24);
    var other = lines[2].substring(25);
    if (op == "+") {
      operation = op;
      opParam = BigInt.from(int.parse(other));
    } else {
      // op == "*"
      if (other == "old") {
        operation = "sq";
      } else {
        operation = op;
        opParam = BigInt.from(int.parse(other));
      }
    }
    testDivider = BigInt.from(int.parse(lines[3].substring(21)));
    targetTrue = int.parse(lines[4].substring(29));
    targetFalse = int.parse(lines[5].substring(30));
  }

  void catchItem(BigInt item) {
    items.add(item);
  }

  void throwItems(List<Monkey> monkeys) {
    for (BigInt item in items) {
      switch (operation) {
        case "+":
          item = item + opParam;
          break;
        case "*":
          item = item * opParam;
          break;
        case "sq":
          item = item.pow(2);
          break;
      }
      if (worryDivider == BigInt.from(3)) {
        item = item ~/ worryDivider;
      } else {
        item = item % worryDivider;
      }
      int target = item % testDivider == BigInt.zero ? targetTrue : targetFalse;
      monkeys[target].catchItem(item);
      inspected++;
    }
    items = List<BigInt>.empty(growable: true);
  }

  String AsString() {
    return "Monkey ${pos} (${inspected}): ${items.join(', ')}";
  }
}

const bool DIVIDE_WORRY = false;
const int ROUNDS = 10000;

void main() {
  var file = File("input.txt");
  var lines = file.readAsLinesSync();
  var numMonkeys = (lines.length / 7).floor() + 1;
  var monkeys = List<Monkey>.generate(
      numMonkeys, (index) => Monkey.fromLines(lines.sublist(index * 7)));

  if (!DIVIDE_WORRY) {
    var worryDivider = monkeys
        .map((m) => m.testDivider)
        .reduce((value, element) => value * element);
    for (final monkey in monkeys) {
      monkey.worryDivider = worryDivider;
    }
  }

  for (final monkey in monkeys) {
    print(monkey.AsString());
  }

  for (int round = 1; round <= ROUNDS; round++) {
    for (final monkey in monkeys) {
      monkey.throwItems(monkeys);
    }
    print("After round ${round}");
    // for (final monkey in monkeys) {
    //   print(monkey.AsString());
    // }
  }

  monkeys.sort((Monkey a, Monkey b) => b.inspected.compareTo(a.inspected));

  for (final monkey in monkeys) {
    print("Monkey ${monkey.pos}: ${monkey.inspected}");
  }
  print("Monkey Business: ${monkeys[0].inspected * monkeys[1].inspected}");
}
