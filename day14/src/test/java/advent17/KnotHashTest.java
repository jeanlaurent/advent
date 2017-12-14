package advent17;

import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThat;

public class KnotHashTest {
    int[] sampleList = new int[] {0, 1, 2, 3, 4};
    int[] sampleLength = new int[] {3, 4, 1, 5};

    @Test
    public void first_turn() {
        KnotHash knotHash = new KnotHash(sampleList, sampleLength);
        knotHash.turn(0);
        assertThat(knotHash.list).containsExactly(2,1,0,3,4);
        assertThat(knotHash.index).isEqualTo(3);
        assertThat(knotHash.skipSize).isEqualTo(1);
    }

    @Test
    public void second_turn() {
        KnotHash knotHash = new KnotHash(new int[] {2,1,0,3,4}, sampleLength,3,1);
        knotHash.turn(1);
        assertThat(knotHash.list).containsExactly(4,3,0,1,2);
        assertThat(knotHash.index).isEqualTo(3);
        assertThat(knotHash.skipSize).isEqualTo(2);
    }

    @Test
    public void third_turn() {
        KnotHash knotHash = new KnotHash(new int[]{4, 3, 0, 1, 2}, sampleLength, 3, 2);
        knotHash.turn(2);
        assertThat(knotHash.list).containsExactly(4, 3, 0, 1, 2);
        assertThat(knotHash.index).isEqualTo(1);
        assertThat(knotHash.skipSize).isEqualTo(3);
    }

    @Test
    public void fourth_turn() {
        KnotHash knotHash = new KnotHash(new int[]{4, 3, 0, 1, 2}, sampleLength, 1, 3);
        knotHash.turn(3);
        assertThat(knotHash.list).containsExactly(3, 4, 2, 1, 0);
        assertThat(knotHash.index).isEqualTo(4);
        assertThat(knotHash.skipSize).isEqualTo(4);
    }

    @Test
    public void sample() {
        KnotHash knotHash = new KnotHash(sampleList, sampleLength);
        assertThat(knotHash.run()).isEqualTo(12);
        assertThat(knotHash.list).containsExactly(3, 4, 2, 1, 0);
        assertThat(knotHash.index).isEqualTo(4);
        assertThat(knotHash.skipSize).isEqualTo(4);
    }

    @Test
    public void part1() {
        int[] lengths = {183,0,31,146,254,240,223,150,2,206,161,1,255,232,199,88};
        KnotHash knotHash = new KnotHash(_256List(), lengths);
        assertThat(knotHash.run()).isEqualTo(15990);
    }

    @Test
    public void part2_emptystring() {
        assertThat((new KnotHash("")).part2()).isEqualTo("a2582a3a0e66e6e86e3812dcb672a272");
    }

    @Test
    public void part2_aoc() {
        assertThat((new KnotHash("AoC 2017")).part2()).isEqualTo("33efeb34ea91902bb2f59c9920caa6cd");
    }

    @Test
    public void part2_123() {
        assertThat((new KnotHash("1,2,3")).part2()).isEqualTo("3efbe78a8d82f29979031a4aa0b16a9d");
    }

    @Test
    public void part2_124() {
        assertThat((new KnotHash("1,2,4")).part2()).isEqualTo("63960835bcdc130f0b66d7ff4f6a5a8e");
    }

    @Test
    public void part2() {
        assertThat(new KnotHash("183,0,31,146,254,240,223,150,2,206,161,1,255,232,199,88").part2()).isEqualTo("90adb097dd55dea8305c900372258ac6");;
    }

    int[] _256List() {
        int[] list = new int[256];
        for (int i = 0; i < 256; i++) {
            list[i] = i;
        }
        return list;
    }

}
