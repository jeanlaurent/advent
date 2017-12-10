package advent17;

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
    
    public int run() {
        for (int i = 0; i < lengths.length; i++) {
            turn(i);
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

}
