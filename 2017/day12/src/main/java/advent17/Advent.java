package advent17;

import java.util.*;

public class Advent
{
    private final HashMap<Integer, List<Integer>> visited;
    private final List<Integer> found;

    public Advent() {
        found = new ArrayList<>();
        visited = new HashMap<>();
    }

    void add(int number) {
        if (!found.contains(number)) {
            found.add(number);
            if (visited.containsKey(number)) {
                for (Integer visitedNumber: visited.get(number)) {
                    add(visitedNumber);
                }
            }
        }
    }

    public int findNumberOfProgramLinkedTo0(String input) {
        readInput(input);
        return findNumberOfProgramLinked(0);
    }

    public void readInput(String input) {
        String[] lines = input.split("\n");
        for (String line : lines) {
            String[] strings = line.split("<->");
            int left = Integer.valueOf(strings[0].trim());
            List<Integer> rights = new ArrayList<>();
            Arrays.stream(strings[1].split(",")).forEach(s -> {
                rights.add(Integer.valueOf(s.trim()));
            });
            visited.put(left, rights);
        }
    }

    public int findNumberOfProgramLinked(int linkNumber) {
        for( Integer left: visited.keySet()) {
            List<Integer> rights = visited.get(left);
            if (left == linkNumber)  {
                add(left);
                for (int right : rights) {
                    add(right);
                }
            } else {
                if (found.contains(left)) {
                    for (int right : rights) {
                        add(right);
                    }
                }
            } {
                for (int right : rights) {
                    if (right == linkNumber) {
                        add(left);
                        add(right);
                    }
                }
            }
        }
        return found.size();
    }

    public int findNumberOfGroups(String input) {
        readInput(input);
        findNumberOfProgramLinked(0);
        Set<Integer> keys = visited.keySet();
        int groupCount = 1;
        found.forEach(keys::remove);
        found.clear();

        while(!keys.isEmpty()) {
            Optional<Integer> first = keys.stream().sorted().findFirst();
            if (first.isPresent()) {
                findNumberOfProgramLinked(first.get());
                found.forEach(keys::remove);
                found.clear();
                groupCount++;
            }
        }
        return groupCount;
    }
}
