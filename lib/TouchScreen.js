
const EventEmitter = require('events');

const InputEvent = require('input-event');

const AXIS = {
    AXIS_X    : 53,
    AXIS_Y    : 54,
};

class TouchScreen extends EventEmitter {
    constructor(input) {
        super();
        input.on('data', (ev, part) => {
            this.parse(ev);
        });
        this.x = null;
        this.y = null;
        this.scaleFactorX = 1.0;
        this.scaleFactorY = 1.0;
    }

    SetScaleFactor(scaleFactorX, scaleFactorY) {
        this.scaleFactorX = scaleFactorX;
        this.scaleFactorY = scaleFactorY;
    }

    /* type, code, value */
    parse(ev) {

        //console.log(ev);

        if (InputEvent.EVENT_TYPES.EV_ABS == ev.type) {
            if (AXIS.AXIS_X == ev.code) {
                this.x = ev.value;
            }
            if (AXIS.AXIS_Y == ev.code) {
                this.y = ev.value;
            }
            if (57 == ev.code) {
                const event = {
                    rawX: this.x,
                    rawY: this.y,
                    x: this.x * this.scaleFactorX,
                    y: this.y * this.scaleFactorY,
                    scaleFactorX: this.scaleFactorX,
                    scaleFactorY: this.scaleFactorY,
                };
                this.emit('touch', event);
            }
        }
    }
}

module.exports = exports = TouchScreen;
