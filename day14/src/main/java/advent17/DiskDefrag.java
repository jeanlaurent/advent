package advent17;

import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.ToString;

import java.math.BigInteger;
import java.util.*;
import java.util.stream.Collectors;

public class DiskDefrag {

    String pad = "0000";

    public List<String> computeGrid(String prefix) {
        List<String> binaries = new ArrayList<>();
        for (int index = 0; index < 128; index++) {
            String knotHash = computeHash(prefix + "-" + index);
            for (int i = 0; i < knotHash.length(); i++) {
                String binary = new BigInteger("" + knotHash.charAt(i), 16).toString(2);
                binaries.add(pad.substring(binary.length()) + binary);
            }
        }
        return binaries;
    }

    public int used(String prefix) {
        List<String> strings = computeGrid(prefix);
        int used = 0;
        for(String binary: strings) {
            for (int j = 0; j < binary.length(); j++) {
                if (binary.charAt(j) == '1') {
                    used++;
                }
            }
        }
        return used;
    }

    public int region(String prefix) {
        List<String> strings = computeGrid(prefix);
        StringBuilder mapBuilder = new StringBuilder();
        for(String string: strings) {
            mapBuilder.append(string);
        }
        Set<Coord> visitedCoord = new HashSet<>();
        String map = mapBuilder.toString();
        int group = 0;
        for (int x = 0; x < 128; x++) {
            for (int y = 0; y < 128; y++) {
                Coord currentCoord = new Coord(x,y);
                if (!visitedCoord.contains(currentCoord)) {
                    if (getChar(map, currentCoord) == '1') {
                        group++;
                        markAsVisitedAdjacent(currentCoord,visitedCoord, map);
                    }
                }
            }
        }
        return group;
    }

    public void markAsVisitedAdjacent(Coord center, Set<Coord> visitedCoord, String map) {
        List<Coord> adjacentList = center.adjacent().stream().filter(c ->
                c.isValid() &&
                        !visitedCoord.contains(c) &&
                        getChar(map, c) == '1')
                .collect(Collectors.toList());
        adjacentList.forEach(c -> visitedCoord.add(c));
        adjacentList.forEach(c -> markAsVisitedAdjacent(c,visitedCoord, map));
    }

    public char getChar(String map, Coord coord) {
        return map.charAt(coord.x + coord.y * 128);
    }

    public String computeHash(String stringtobehashed) {
        return new KnotHash(stringtobehashed).part2();
    }

    @AllArgsConstructor
    @EqualsAndHashCode
    @ToString
    class Coord {
        int x;
        int y;

        boolean isValid() {
            return x >=0 && x < 128 && y >=0 && y < 128;
        }

        Coord north() {
            return new Coord(x,y-1);
        }
        Coord east() {
            return new Coord(x-1,y);
        }

        Coord west() {
            return new Coord(x+1,y);
        }

        Coord south() {
            return new Coord(x,y+1);
        }

        List<Coord> adjacent() {
            return Arrays.asList(north(), east(), south(), west());
        }

    }
}
