import java.io.IOException;

import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;
import java.util.stream.Stream;
import java.util.PriorityQueue;

import java.lang.Character;
import java.lang.Integer;

public class Main {
  public static void main(String[] args) throws IOException {
    String fileName = "input.txt";

    // read file into stream, try-with-resource
    try (Stream<String> stream = Files.lines(Paths.get(fileName))) {
      // int total = stream
      //     .filter(Main::isRealRoom)
      //     .mapToInt(Main::getSectorId)
      //     .sum();
      // System.out.println(total);

      stream
        .map(Main::shift)
        .forEach(System.out::println);

    } catch (IOException e) {
      e.printStackTrace();
    }
  }

  public static String shift(String s) {
    int id = getSectorId(s);
    s = s.substring(0, s.length()-7);
    int rotBy = id % 26;
    StringBuilder sb = new StringBuilder(s.length());
    for (int i = 0; i < s.length(); i++) {
      char c = s.charAt(i);
      if (c == '-') {
        sb.append(' ');
        continue;
      }
      sb.append((char) (((c - 'a' + rotBy) % 26) + 'a'));
    }
    sb.append(" " + id);
    return sb.toString();
  }

  public static int getSectorId(String s) {
    return Integer.parseInt(s.substring(s.lastIndexOf("-")+1, s.indexOf("[")));
  }

  public static boolean isRealRoom(String s) {
    String a = s.substring(s.length()-6, s.length()-1);
    String b = topFiveLetters(countLetters(s));

    return a.equals(b);
  }

  // countLetters counts the number of a-z characters.
  public static int[] countLetters(String s) {
    int[] counts = new int[26];
    for (int i = 0; i < s.length(); i++) {
      char c = s.charAt(i);
      if (c < 'a' || c > 'z') {
        continue;
      }
      counts[c - 'a']++;
    }
    return counts;
  }

  public static String topFiveLetters(int[] counts) {
    PriorityQueue<LetterCount> pq = new PriorityQueue<>();
    for (int i = 0; i < counts.length; i++) {
      pq.add(new LetterCount((char) (i+'a'), counts[i]));
    }

    StringBuilder sb = new StringBuilder(5);
    for (int i = 0; i < 5; i++) {
      LetterCount x = pq.poll();
      sb.append(x.letter);
    }

    return sb.toString();
  }

  private static class LetterCount implements Comparable<LetterCount> {
    public char letter;
    public int count;

    public LetterCount(char letter, int count) {
      this.letter = letter;
      this.count = count;
    }

    @Override
    public int compareTo(LetterCount other) {
      int c = Integer.compare(this.count, other.count);
      if (c != 0) {
        // biggest first
        return -c;
      }
      // alphabetical
      return Character.compare(this.letter, other.letter);
    }
  }
}
