package advent17;

import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThat;

public class AdventTest {

    @Test
    public void is_zero() {
        assertThat(new Advent().foobar("foo")).isEqualTo(0);
    }
}
