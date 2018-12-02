package advent17;

import org.junit.Test;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

import static org.assertj.core.api.Assertions.assertThat;

public class AdventTest {

    @Test
    public void example_step1() {
        Advent advent = new Advent();
        assertThat(advent.firewall("0: 3\n" +
                "1: 2\n" +
                "4: 4\n" +
                "6: 4\n")).isEqualTo(24);
    }

    @Test
    public void example_step2() {
        Advent advent = new Advent();
        assertThat(advent.firewall2("0: 3\n" +
                "1: 2\n" +
                "4: 4\n" +
                "6: 4\n")).isEqualTo(10);
    }


    @Test
    public void move_up() {
        assertThat(new Advent().findPosition(0,4)).isEqualTo(0);
        assertThat(new Advent().findPosition(1,4)).isEqualTo(1);
        assertThat(new Advent().findPosition(2,4)).isEqualTo(2);
        assertThat(new Advent().findPosition(3,4)).isEqualTo(3);
        assertThat(new Advent().findPosition(4,4)).isEqualTo(2);
        assertThat(new Advent().findPosition(5,4)).isEqualTo(1);
        assertThat(new Advent().findPosition(6,4)).isEqualTo(0);
        assertThat(new Advent().findPosition(7,4)).isEqualTo(1);
        assertThat(new Advent().findPosition(8,4)).isEqualTo(2);
        assertThat(new Advent().findPosition(9,4)).isEqualTo(3);
    }

    @Test
    public void Step1() throws IOException {
        Advent advent = new Advent();
        String input = new String(Files.readAllBytes(new File("./input.txt").toPath()));
        assertThat(advent.firewall(input)).isEqualTo(2604);
    }

    @Test
    public void Step2() throws IOException {
        Advent advent = new Advent();
        String input = new String(Files.readAllBytes(new File("./input.txt").toPath()));
        assertThat(advent.firewall2(input)).isEqualTo(2604);
    }
}
