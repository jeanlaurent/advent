package advent17;

import org.junit.Test;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

import static org.assertj.core.api.Assertions.assertThat;

public class AdventTest {
    int[] sampleList = new int[] {0, 1, 2, 3, 4};
    int[] sampleLength = new int[] {3, 4, 1, 5};

    @Test
    public void first_turn() {
        Advent advent = new Advent(sampleList, sampleLength);
        advent.turn(0);
        assertThat(advent.list).containsExactly(2,1,0,3,4);
        assertThat(advent.index).isEqualTo(3);
        assertThat(advent.skipSize).isEqualTo(1);
    }

    @Test
    public void second_turn() {
        Advent advent = new Advent(new int[] {2,1,0,3,4}, sampleLength,3,1);
        advent.turn(1);
        assertThat(advent.list).containsExactly(4,3,0,1,2);
        assertThat(advent.index).isEqualTo(3);
        assertThat(advent.skipSize).isEqualTo(2);
    }

    @Test
    public void third_turn() {
        Advent advent = new Advent(new int[]{4, 3, 0, 1, 2}, sampleLength, 3, 2);
        advent.turn(2);
        assertThat(advent.list).containsExactly(4, 3, 0, 1, 2);
        assertThat(advent.index).isEqualTo(1);
        assertThat(advent.skipSize).isEqualTo(3);
    }

    @Test
    public void fourth_turn() {
        Advent advent = new Advent(new int[]{4, 3, 0, 1, 2}, sampleLength, 1, 3);
        advent.turn(3);
        assertThat(advent.list).containsExactly(3, 4, 2, 1, 0);
        assertThat(advent.index).isEqualTo(4);
        assertThat(advent.skipSize).isEqualTo(4);
    }

    @Test
    public void sample() {
        Advent advent = new Advent(sampleList, sampleLength);
        assertThat(advent.run()).isEqualTo(12);
        assertThat(advent.list).containsExactly(3, 4, 2, 1, 0);
        assertThat(advent.index).isEqualTo(4);
        assertThat(advent.skipSize).isEqualTo(4);
    }

    @Test
    public void part1() {
        int[] list = new int[256];
        for (int i = 0; i < 256; i++) {
            list[i] = i;
        }
        int[] lengths = {183,0,31,146,254,240,223,150,2,206,161,1,255,232,199,88};
        Advent advent = new Advent(list, lengths);
        assertThat(advent.run()).isEqualTo(15990);
    }

}
