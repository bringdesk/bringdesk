
import parse from 'parse-color';

class BoxWidget {

    constructor(options) {
        this._options = options;
        this.child = null;
    }

    start() {
    }

    stop() {
    }

    mergeOptions(mainOptions, currentOptions) {
        const options = {};
        /* Overlay current options */
        Object.keys(currentOptions).forEach((key) => {
            const value = currentOptions[key];
            options[key] = value;
        });
        /* Overlay main options */
        Object.keys(mainOptions).forEach((key) => {
            const value = mainOptions[key];
            options[key] = value;
        });
        /* Done */
        return options;
    }

    render(options1) {
        const options = this.mergeOptions(this._options, options1);
        const screen = options.screen;
        const position = options.position;
        const color = options.color;
        /* Debug message */
        console.log('Box: ', options);
        /* Card position */
        const left = position.left;
        const top = position.top;
        const width = position.width;
        const height = position.height;
        /* Card box */
        if (color) {
            console.log("Color:", color);
            /* Debug message */
            const newColor = parse(color);
            const rgba = newColor.rgba;
            /* Render */
            screen.SetColor(rgba[0] / 255.0, rgba[1] / 255.0, rgba[2] / 255.0, 1.0);
            screen.DrawRectangle(left, top, width, height);
        }
        /* Children */
        if (this.child) {
            const newOptions = {
                screen: screen,
                position: {
                    left: left + 15,
                    top: top + 15,
                    width: width - 2*15,
                    height: height - 2*15,
                },
            };
            console.log(newOptions);
            this.child.render(newOptions);
        }
    }

}

export {
    BoxWidget
}
