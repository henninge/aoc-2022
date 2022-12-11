import 'dart:io';
import 'dart:math';

class Monkey {
  int pos = 0;
  List<int> items = List.empty(growable: true);
  String operation = "";
  int opParam = 0;
  int testDivider = 1;
  int targetTrue = 0;
  int targetFalse = 0;
  int inspected = 0;

  Monkey.fromLines(List<String> lines) {
    pos = int.parse(lines[0].substring(7, 8));
    items = List.of(
        lines[1].substring(18).split(', ').map((str) => int.parse(str)));

    var op = lines[2].substring(23, 24);
    var other = lines[2].substring(25);
    if (op == "+") {
      operation = op;
      opParam = int.parse(other);
    } else {
      // op == "*"
      if (other == "old") {
        operation = "sq";
      } else {
        operation = op;
        opParam = int.parse(other);
      }
    }
    testDivider = int.parse(lines[3].substring(21));
    targetTrue = int.parse(lines[4].substring(29));
    targetFalse = int.parse(lines[5].substring(30));
  }

  void catchItem(int item) {
    items.add(item);
  }

  void throwItems(List<Monkey> monkeys) {
    for (int item in items) {
      switch (operation) {
        case "+":
          item = item + opParam;
          break;
        case "*":
          item = item * opParam;
          break;
        case "sq":
          item = item * item;
          break;
      }
      item = (item / 3).floor();
      monkeys[item % testDivider == 0 ? targetTrue : targetFalse]
          .catchItem(item);
      inspected++;
    }
    items = List<int>.empty(growable: true);
  }

  String AsString() {
    return "Monkey ${pos}: ${items.join(', ')}";
  }
}

void main() {
  var file = File("input.txt");
  var lines = file.readAsLinesSync();
  var numMonkeys = (lines.length / 7).floor() + 1;
  var monkeys = List<Monkey>.generate(
      numMonkeys, (index) => Monkey.fromLines(lines.sublist(index * 7)));

  for (int round = 1; round <= 20; round++) {
    for (final monkey in monkeys) {
      monkey.throwItems(monkeys);
    }
    print("After round ${round}");
    for (final monkey in monkeys) {
      print(monkey.AsString());
    }
  }

  monkeys.sort((Monkey a, Monkey b) => b.inspected.compareTo(a.inspected));

  for (final monkey in monkeys) {
    print("Monkey ${monkey.pos}: ${monkey.inspected}");
  }
  print("Monkey Business: ${monkeys[0].inspected * monkeys[1].inspected}");
}
