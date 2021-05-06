
import parse from 'parse-color';

class BoxWidget {

    constructor(options) {
        this.options = options;
    }

    setup(parent) {
        const {content} = this.options;
        if (content) {
            content.setup(parent);
        }
    }

    render(parent) {

        const {
            rect = [ 0, 0, 0, 0 ],
            background,
//            border,
            padding = [0, 0, 0, 0],
            content,
        } = this.options;

//        /* Если указаны край коробки */
//        if (border) {
//            const newBorder = parse(border);
//            if (newBorder) {
//                parent.FillRect(rect, newBorder);
//            }
//        }

        /* Если указа фон коробки */
        if (background) {
            const newBackground = parse(background);
            if (newBackground) {
                parent.FillRect(rect, newBackground);
            }
        }

        /* Рисуем содержимое коробки */
        if (content) {
            content.render(parent);
        }

    }

}

export {
    BoxWidget
}
