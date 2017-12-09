package advent17;

import org.junit.Test;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

import static org.assertj.core.api.Assertions.assertThat;

public class AdventTest {

    @Test
    public void find_one_group() {
        Advent advent = new Advent();
        advent.parse("{}");
        assertThat(advent.groups).isEqualTo(1);
        assertThat(advent.score).isEqualTo(1);
        assertThat(advent.characters).isEqualTo(0);
    }

    @Test
    public void find_two_nested_group() {
        Advent advent = new Advent();
        advent.parse("{{{}}}");
        assertThat(advent.groups).isEqualTo(3);
        assertThat(advent.score).isEqualTo(6);
        assertThat(advent.characters).isEqualTo(0);
    }

    @Test
    public void find_two_group_in_an_array() {
        Advent advent = new Advent();
        advent.parse("{{},{}}");
        assertThat(advent.groups).isEqualTo(3);
        assertThat(advent.score).isEqualTo(5);
        assertThat(advent.characters).isEqualTo(0);
    }

    @Test
    public void find_six_groups() {
        Advent advent = new Advent();
        advent.parse("{{{},{},{{}}}}");
        assertThat(advent.groups).isEqualTo(6);
        assertThat(advent.score).isEqualTo(16);
        assertThat(advent.characters).isEqualTo(0);
    }

    @Test
    public void find_one_group_with_big_garbage() {
        Advent advent = new Advent();
        advent.parse("{<{},{},{{}}>}");
        assertThat(advent.groups).isEqualTo(1);
        assertThat(advent.characters).isEqualTo(10);
    }

    @Test
    public void find_one_group_with_lots_of_tiny_garbage() {
        Advent advent = new Advent();
        advent.parse("{<a>,<a>,<a>,<a>}");
        assertThat(advent.groups).isEqualTo(1);
        assertThat(advent.score).isEqualTo(1);
        assertThat(advent.characters).isEqualTo(4);
    }

    @Test
    public void find_five_group_with_lots_of_tiny_garbage() {
        Advent advent = new Advent();
        advent.parse("{{<a>},{<a>},{<a>},{<a>}}");
        assertThat(advent.groups).isEqualTo(5);
        assertThat(advent.score).isEqualTo(9);
    }

    @Test
    public void find_2_groups_with_skips() {
        Advent advent = new Advent();
        advent.parse("{{<!>},{<!>},{<!>},{<a>}}");
        assertThat(advent.groups).isEqualTo(2);
    }

    @Test
    public void find_score_with_lots_of_escape_chars() {
        Advent advent = new Advent();
        advent.parse("{{<!!>},{<!!>},{<!!>},{<!!>}}");
        assertThat(advent.score).isEqualTo(9);
    }

    @Test
    public void find_score_with_weird_string() {
        Advent advent = new Advent();
        advent.parse("{{<a!>},{<a!>},{<a!>},{<ab>}}");
        assertThat(advent.score).isEqualTo(3);
    }

    @Test
    public void find_first_solution() throws IOException {
        Advent advent = new Advent();
        String input = new String(Files.readAllBytes(new File("./input.txt").toPath()));
        advent.parse(input);
        assertThat(advent.score).isEqualTo(14190);
        assertThat(advent.characters).isEqualTo(7053);
    }

    @Test
    public void count_garbage() {
        Advent advent = new Advent();
        advent.parse("<{!>}>");
        assertThat(advent.characters).isEqualTo(2);
    }

    @Test
    public void count_garbage_with_skips() {
        Advent advent = new Advent();
        advent.parse("<!!>");
        assertThat(advent.characters).isEqualTo(0);
    }

    @Test
    public void count_garbage_with_more_skips() {
        Advent advent = new Advent();
        advent.parse("<!!!>>");
        assertThat(advent.characters).isEqualTo(0);
    }

    @Test
    public void count_garbage_with_lots_of_opening() {
        Advent advent = new Advent();
        advent.parse("<<<<>");
        assertThat(advent.characters).isEqualTo(3);
    }

    @Test
    public void count_garbage_with_weird_string() {
        Advent advent = new Advent();
        advent.parse("<{o\"i!a,<{i<a>");
        assertThat(advent.characters).isEqualTo(10);
    }


}
