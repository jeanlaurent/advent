package advent17;

import org.junit.Test;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

import static org.assertj.core.api.Assertions.assertThat;

public class AdventTest {

    @Test
    public void zero_contained() {
        Advent advent = new Advent();
        assertThat(advent.findNumberOfProgramLinkedTo0("0 <-> 2")).isEqualTo(2);
    }

    @Test
    public void zero_contained_on_right() {
        Advent advent = new Advent();
        assertThat(advent.findNumberOfProgramLinkedTo0("2 <-> 0")).isEqualTo(2);
    }

    @Test
    public void no_zero_contained() {
        Advent advent = new Advent();
        assertThat(advent.findNumberOfProgramLinkedTo0("3 <-> 2")).isEqualTo(0);
    }

    @Test
    public void Threelines() {
        Advent advent = new Advent();
        assertThat(advent.findNumberOfProgramLinkedTo0("0 <-> 2\n1 <-> 1\n2 <-> 0, 3, 4")).isEqualTo(4);
    }

    @Test
    public void Fourlines() {
        Advent advent = new Advent();
        assertThat(advent.findNumberOfProgramLinkedTo0("0 <-> 2\n1 <-> 1\n2 <-> 0, 3, 4\n3 <-> 2, 4")).isEqualTo(4);
    }

    @Test
    public void Fivelines() {
        Advent advent = new Advent();
        assertThat(advent.findNumberOfProgramLinkedTo0("0 <-> 2\n1 <-> 1\n2 <-> 0, 3, 4\n3 <-> 2, 4\n4 <-> 2, 3, 6")).isEqualTo(5);
    }

    @Test
    public void Sixlines() {
        Advent advent = new Advent();
        assertThat(advent.findNumberOfProgramLinkedTo0("0 <-> 2\n1 <-> 1\n2 <-> 0, 3, 4\n3 <-> 2, 4\n4 <-> 2, 3, 6\n5 <-> 6")).isEqualTo(5);
    }

    @Test
    public void Sevenlines() {
        Advent advent = new Advent();
        assertThat(advent.findNumberOfProgramLinkedTo0("0 <-> 2\n1 <-> 1\n2 <-> 0, 3, 4\n3 <-> 2, 4\n4 <-> 2, 3, 6\n5 <-> 6\n6 <-> 4, 5")).isEqualTo(6);
    }

    @Test
    public void SevenlinesPart2() {
        Advent advent = new Advent();
        assertThat(advent.findNumberOfGroups("0 <-> 2\n1 <-> 1\n2 <-> 0, 3, 4\n3 <-> 2, 4\n4 <-> 2, 3, 6\n5 <-> 6\n6 <-> 4, 5")).isEqualTo(2);
    }

    @Test
    public void AddAfterlines() {
        Advent advent = new Advent();
        assertThat(advent.findNumberOfProgramLinkedTo0("0 <-> 1, 4\n2 <-> 3\n3 <-> 3\n4 <-> 2")).isEqualTo(5);
    }

    @Test
    public void Step1() throws IOException {
        Advent advent = new Advent();
        String input = new String(Files.readAllBytes(new File("./input.txt").toPath()));
        assertThat(advent.findNumberOfProgramLinkedTo0(input)).isEqualTo(283);
    }


    @Test
    public void Step2() throws IOException {
        Advent advent = new Advent();
        String input = new String(Files.readAllBytes(new File("./input.txt").toPath()));
        assertThat(advent.findNumberOfGroups(input)).isEqualTo(195);
    }
}
