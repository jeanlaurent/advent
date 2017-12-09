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
        for (int index = 0; index < line.length(); index++) {
            openingGarbage = false;
            if (line.charAt(index) == '!') {
                index++;
                continue;
            }
            if (line.charAt(index) == '<') {
                openingGarbage = !inGarbage;
                inGarbage = true;
            }
            if (line.charAt(index) == '>') {
                inGarbage = false;
            }
            if ((inGarbage) && (!openingGarbage)) {
                characters++;
            } else {
                if (line.charAt(index) == '{') {
                    this.nestedGroup++;
                    score += nestedGroup;
                }
                if (line.charAt(index) == '}') {
                    groups++;
                    nestedGroup--;
                }
            }
        }
    }
}
