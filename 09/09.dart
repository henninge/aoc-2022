import 'dart:io';
import 'dart:math';

class Knot {
  int x = 0, y = 0;

  void followHead(Knot head) {
    if ((head.x - x).abs() <= 1 && (head.y - y).abs() <= 1) {
      // Still touching
      return;
    }
    // Move ±1 or 0 towards the head.
    x += (head.x - x).sign;
    y += (head.y - y).sign;
  }
}

class Rope {
  List<Knot> knots = [];
  Set<Point<int>> visited = {};

  Rope(int length) {
    knots = List<Knot>.generate(length, (index) => Knot());
    visited.add(Point<int>(0, 0));
  }

  int getVisited() {
    return visited.length;
  }

  void parseInstruction(String line) {
    final parts = line.split(' ');
    String dir = parts[0];
    int dist = int.parse(parts[1]);

    for (int i = 0; i < dist; i++) {
      switch (dir) {
        case "U":
          knots.first.y++;
          break;
        case "D":
          knots.first.y--;
          break;
        case "R":
          knots.first.x++;
          break;
        case "L":
          knots.first.x--;
          break;
      }
      for (int k = 1; k < knots.length; k++) {
        knots[k].followHead(knots[k - 1]);
      }
      visited.add(Point<int>(knots.last.x, knots.last.y));
    }
  }
}

void main() {
  // Rope rope = Rope(2);
  Rope rope = Rope(10);
  var file = File("input.txt");
  for (final line in file.readAsLinesSync()) {
    rope.parseInstruction(line);
  }

  print("Visited: ${rope.getVisited()}");
}
