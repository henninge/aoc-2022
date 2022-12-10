import 'dart:io';
import 'dart:math';

List<List<int>> forest = [];

int calcVisibilty() {
  Set<Point<int>> visible = {};
  int row, col, largest;

  int checkTree(int largest, row, col) {
    if (forest[row][col] > largest) {
      visible.add(Point<int>(row, col));
      largest = forest[row][col];
    }
    return largest;
  }

  // Each row forward
  for (row = 0; row < forest.length; row++) {
    largest = -1;
    for (col = 0; col < forest[row].length; col++) {
      largest = checkTree(largest, row, col);
      if (largest == 9) break;
    }
  }

  // Each row backward
  for (row = 0; row < forest.length; row++) {
    largest = -1;
    for (col = forest[row].length - 1; col >= 0; col--) {
      largest = checkTree(largest, row, col);
      if (largest == 9) break;
    }
  }

  // Each col forward
  for (col = 0; col < forest[0].length; col++) {
    largest = -1;
    for (row = 0; row < forest.length; row++) {
      largest = checkTree(largest, row, col);
      if (largest == 9) break;
    }
  }

  // Each col row
  for (col = 0; col < forest[0].length; col++) {
    largest = -1;
    for (row = forest.length - 1; row >= 0; row--) {
      largest = checkTree(largest, row, col);
      if (largest == 9) break;
    }
  }
  return visible.length;
}

enum Direction { north, east, south, west }

int viewDistance(int row, col, Direction dir) {
  if (dir == Direction.north && row == 0 ||
      dir == Direction.west && col == forest[0].length - 1 ||
      dir == Direction.south && row == forest.length - 1 ||
      dir == Direction.east && col == 0) {
    return 0;
  }

  int delta;
  switch (dir) {
    case Direction.north:
      delta = -1;
      break;
    case Direction.east:
      delta = 1;
      break;
    case Direction.south:
      delta = 1;
      break;
    case Direction.west:
      delta = -1;
      break;
  }
  var centerHeight = forest[row][col];
  if (dir == Direction.north || dir == Direction.south) {
    for (int r = row + delta; r >= 0 && r < forest.length; r += delta) {
      if (forest[r][col] >= centerHeight) {
        return (row - r).abs();
      }
    }
    return dir == Direction.north ? row : forest.length - row - 1;
  } else {
    for (int c = col + delta; c >= 0 && c < forest[0].length; c += delta) {
      if (forest[row][c] >= centerHeight) {
        return (col - c).abs();
      }
    }
    return dir == Direction.west ? col : forest[0].length - col - 1;
  }
}

int calcScenicScore(int row, col) {
  int score = 1;
  for (final dir in Direction.values) {
    int dist = viewDistance(row, col, dir);
    score *= dist;
  }
  return score;
}

int calcMaxScenicScore() {
  int maxScore = 0;
  for (int row = 0; row < forest.length; row++) {
    for (int col = 0; col < forest[0].length; col++) {
      int score = calcScenicScore(row, col);
      maxScore = max(score, maxScore);
    }
  }
  return maxScore;
}

void main() {
  var file = File("input.txt");
  for (final line in file.readAsLinesSync()) {
    forest.add(
        List<int>.generate(line.length, (index) => int.parse(line[index])));
  }

  print("Visible: ${calcVisibilty()}");
  print("Scenic Score: ${calcMaxScenicScore()}");
}
