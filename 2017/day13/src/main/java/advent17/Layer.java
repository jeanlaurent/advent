package advent17;

import lombok.ToString;

@ToString
public class Layer {
    int range;
    int position;
    boolean down;

    public Layer(Integer range) {
        this.range = range;
        this.position = 0;
        this.down = true;
    }

    public int move() {
        if ((!down && (position == 0)) || ((position == range -1) && down)) {
            down = !down;
        }
        if (down) {
            position++;
        } else {
            position--;
        }
        return position;
    }

    public int cost(int index) {
        return range * index;
    }
}
