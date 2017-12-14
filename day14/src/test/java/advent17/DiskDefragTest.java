package advent17;

import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThat;

public class DiskDefragTest {

    @Test
    public void defrag() {
        assertThat(new DiskDefrag().used("hwlqcszp")).isEqualTo(0);
    }

}
