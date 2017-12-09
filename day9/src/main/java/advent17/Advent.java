package advent17;


public class Advent
{
    int nestedGroup = 0;
    int groups = 0;
    int score = 0;
    int characters = 0;

    public void parse(String line) {
        boolean inGarbage = false;
        boolean openingGarbage;
        for (int i = 0; i < line.length(); i++) {
            openingGarbage = false;
            if (line.charAt(i) == '!') {
                i++;
                continue;
            }
            if (line.charAt(i) == '<') {
                openingGarbage = !inGarbage;
                inGarbage = true;
            }
            if (line.charAt(i) == '>') {
                inGarbage = false;
            }
            if ((inGarbage) && (!openingGarbage)) {
                characters++;
            } else {
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
