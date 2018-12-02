package advent17;

import java.util.*;

public class Advent
{
    private final Map<Integer, Integer> layers;
    private int points = 0;
    private int maxLayer;
    private boolean collision = false;

    public Advent() {
        layers = new HashMap<>();
    }

    void readInput(String input) {
        String[] lines = input.split("\n");
        for (String line : lines) {
            String[] strings = line.split(":");
            Integer depth = Integer.valueOf(strings[0].trim());
            Integer range = Integer.valueOf(strings[1].trim());
            layers.put(depth, range);
            maxLayer = depth;
//            System.out.println("maxLayer is now " +maxLayer);
        }
//        System.out.println(layers);
    }
    int firewall(String input) {
        readInput(input);
        for (int i = 0; i <= maxLayer; i++) {
            turn(0, i);
        }
        return points;
    }

    int firewall2(String input) {
        readInput(input);
        int delay = 0;
        while(true) {
            if ((delay % 1000) == 0) {
                System.out.println(delay);
            }
//            System.out.println("=============== trying with delay ---> " + delay);
            for (int i = 0; i <= maxLayer; i++) {
                turn(delay, i);
            }
            if (!collision && (points == 0)) {
                return delay;
            }
            points = 0;
            collision = false;
            delay++;
        }
    }

    void turn(int delay, int layerNumber) {
        int turnNumber = delay + layerNumber;
//        System.out.println("Turn " + turnNumber);
        if (!layers.containsKey(layerNumber)) {
//            System.out.println("no layer");
//            System.out.println();
            return;
        }
        int layerRange = layers.get(layerNumber);
//        System.out.println("Layer is #" + layerNumber + ", range " + layerRange);
        int scannerPosition = findPosition(turnNumber, layerRange);
//        System.out.println("Current scanner position " + scannerPosition);
        if (scannerPosition == 0) {
            collision = true;
//            System.out.println("Collision : " + layerRange * layerNumber);
            points += layerRange * layerNumber;
        }
//        System.out.println();
    }


    int findPosition(int turn, int max) {
        boolean down = true;
        int index = 0;
        for (int i = 0; i < turn ; i++) {
            if ((!down && (index == 0)) || ((index == max -1) && down)) {
                down = !down;
            }
            if (down) {
                index++;
            } else {
                index--;
            }
        }
        return index;
    }
}
