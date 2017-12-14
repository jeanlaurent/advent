package advent17;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.List;

public class DiskDefrag {

    String pad = "0000";

    public int used(String prefix) {
        int used = 0;
        List<String> binaries = new ArrayList<>();
        for (int index = 0; index < 128; index++) {
            String knotHash = computeHash(prefix + "-" + index);
            for (int i = 0; i < knotHash.length(); i++) {
                String binary = new BigInteger("" + knotHash.charAt(i), 16).toString(2);
                for (int j = 0; j < binary.length(); j++) {
                    if (binary.charAt(j) == '1') {
                        used++;
                    }
                }


            }
        }

        return used;
    }

    public String computeHash(String stringtobehashed) {
        return new KnotHash(stringtobehashed).part2();
    }
}
