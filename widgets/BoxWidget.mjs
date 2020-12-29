
import parse from 'parse-color';

class BoxWidget {

    constructor(options) {
        this.options = options;
    }

    setup() {
    }

    render(parent) {

        const {
            rect = [ 0, 0, 0, 0 ],
            color = [ 255, 0, 0 ],
        } = this.options;

        /* Если указа фон коробки */
//      const newColor = parse(color);
        if (color) {
            parent.FillRect(rect, color);
        }

        /* Если указаны край коробки */
        // TODO - ...

    }

}

export {
    BoxWidget
}
