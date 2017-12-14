package advent17;

import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThat;

public class DiskDefragTest {

    @Test
    public void count_used() {
        assertThat(new DiskDefrag().used("hwlqcszp")).isEqualTo(8304);
    }

    @Test
    public void region() {
        assertThat(new DiskDefrag().region("hwlqcszp")).isEqualTo(1018);
    }

}
