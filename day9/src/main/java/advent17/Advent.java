package advent17;


public class Advent
{
    int nestedGroup;
    int groups;
    int score;

    public Advent() {
        this.groups = 0;
        this.score = 0;
        this.nestedGroup = 0;
    }

    public void parse(String line) {
        boolean inGarbage = false;
        for (int i = 0; i < line.length(); i++) {
            if (line.charAt(i) == '!') {
                i++;
                continue;
            }
            if (line.charAt(i) == '<') {
                inGarbage = true;
            }
            if (line.charAt(i) == '>') {
                inGarbage = false;
            }
            if (!inGarbage) {
                if (line.charAt(i) == '{') {
                    this.nestedGroup++;
                    score += nestedGroup;
                }
                if (line.charAt(i) == '}') {
                    groups++;
                    nestedGroup--;
                }
            }
        }
    }
}
