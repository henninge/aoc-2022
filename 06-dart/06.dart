import 'dart:async';
import 'dart:io';
import 'dart:convert';

void main() {
  readLine().listen(processLine);
}

Stream<String> readLine() =>
    stdin.transform(utf8.decoder).transform(const LineSplitter());

void processLine(String line) {
  int packet = getMarker(line, 4);
  int message = getMarker(line, 14);
  print('Packet: $packet Message: $message');
}

bool isUnique(String window) {
  return Set.unmodifiable(window.runes).length == window.length;
}

int getMarker(String line, int length) {
  for (int i = 0; i < line.length - length; i++) {
    if (isUnique(line.substring(i, i + length))) {
      return i + length;
    }
  }
  return 0;
}
