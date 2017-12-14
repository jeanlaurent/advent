package advent17;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Advent
{
    int index = 0;
    int skipSize = 0;
    int[] list;
    int[] lengths;

    public Advent(int[] list, int[] lengths) {
        this.list = list;
        this.lengths = lengths;
    }

    public Advent(int[] list, int[] lengths, int index, int skipSize) {
        this.list = list;
        this.lengths = lengths;
        this.index = index;
        this.skipSize = skipSize;
    }

    public Advent(int[] list, String input) {
        this.list = list;
        List<Integer> numbers = new ArrayList<>();
        for (int i = 0; i < input.length(); i++) {
            numbers.add((int)input.charAt(i));
        }
        numbers.addAll(Arrays.asList(17, 31, 73, 47, 23));
        lengths = toIntArray(numbers);
        for (int i = 0; i < lengths.length; i++) {
            System.out.print(lengths[i] + " ");
        }
        System.out.println();
    }

    // java we love you... not
    int[] toIntArray(List<Integer> list){
        int[] ret = new int[list.size()];
        for(int i = 0;i < ret.length;i++)
            ret[i] = list.get(i);
        return ret;
    }

    public int run() {
        for (int indexLength = 0; indexLength < lengths.length; indexLength++) {
            turn(indexLength);
        }
        return list[0] * list[1];
    }

    public void turn(int indexLength) {
        int reverseLength = lengths[indexLength];
        for (int i = 0; i < reverseLength / 2; i++) {
            int indexA = (index + i) % list.length;
            int indexB = (index + reverseLength - i - 1 ) % list.length;

            int swap = list[indexA];
            list[indexA] = list[indexB];
            list[indexB] = swap;
        }
        index = (index + reverseLength + skipSize) % list.length;
        skipSize++;
    }

    public String part2() {
        for (int i = 0; i < 64; i++) {
            for (int indexLength = 0; indexLength < lengths.length; indexLength++) {
                turn(indexLength);
            }
        }
        int[] blocks = new int[16];
        for (int block = 0; block < 16; block++) {
            blocks[block] = list[block*16];
            for (int i = 1; i < 16; i++) {
                blocks[block] ^= list[block*16 + i];
            }
        }
        StringBuilder hash = new StringBuilder(new String());
        for (int i = 0; i <blocks.length; i++) {
            String hex = Integer.toHexString(blocks[i]);
            if (hex.length() == 1) {
                hex = "0" + hex;
            }
            hash.append(hex);
        }
        System.out.println(Arrays.toString(blocks));
        System.out.println(hash);
        return hash.toString();
    }

}
