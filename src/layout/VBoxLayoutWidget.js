
class VBoxLayoutWidget {
    constructor(options) {
        this._options = options;
        this.children = [];
    }

    start() {
    }

    stop() {
    }

    render(options1) {
        const options = options1 || this._options;
        const screen = options.screen;
        const position = options.position;
        const count = this.children.length;
        const boxSize = position.height / count;
        console.log('boxSize = ', boxSize);
        for (const childrenIndex in this.children) {
            const children = this.children[childrenIndex];
            const newOptions = {};
            /* Set screen */
            newOptions['screen'] = screen;
            /* Set position */
            const position2 = {
                left: position.left,
                top: boxSize * childrenIndex,
                width: position.width,
                height: boxSize,
            };
            newOptions['position'] = position2;
            //
            console.log('newOptions = ', newOptions);
            children.render(newOptions);
        }
    }

}

module.exports = {
    VBoxLayoutWidget
}
