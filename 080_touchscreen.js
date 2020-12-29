#!/usr/bin/env -S node

const InputEvent = require('input-event');
const input = new InputEvent('/dev/input/event3');
const TouchScreen = require('./TouchScreen.js');
const touchscreen = new TouchScreen(input);
const screen = require('./bringdesk_screen/build/Release/screen.node');

const STATE_CALIBRATION = 1;
const STATE_PRODUCTION = 2;

class Application extends BringDesk {

    constructor() {
        const options = {
        };
        this.scr1 = new screen.Screen(options);
        this.mode = 1;
        this.state = STATE_CALIBRATION;
        /* Get screen size */
        const width = this.scr1.width;
        const height = this.scr1.height;
        //
        this.crossX = width / 2.0;
        this.crossY = height / 2.0;
    }

    renderCross() {
        const posX = this.crossX;
        const posY = this.crossY;
        const width = 120;
        const height = 120;
        //
        if (this.state == STATE_CALIBRATION) {
            this.scr1.SetColor(1.0, 0.0, 0.0, 1.0);
            this.scr1.DrawRectangle(posX - (width / 2), posY - (height / 2), width, height);
        } else if (this.state == STATE_PRODUCTION) {
            this.scr1.SetColor(0.0, 1.0, 0.0, 1.0);
            this.scr1.DrawRectangle(posX - (width / 2), posY - (height / 2), width, height);
        }
    }

    renderScreen() {



        /* Clear screen */
        this.scr1.Clear();
        this.renderCross();

        /* Swap image */
        this.scr1.Swap();

    }

    processCalibration(event) {
        /* Get screen size */
        const width = this.scr1.width;
        const height = this.scr1.height;
        //
        const scaleFactorX = (1.0 * event.x) / (width / 2.0);
        const scaleFactorY = (1.0 * event.y) / (height / 2.0);
        //
        console.log('Scale factor x: ', scaleFactorX);
        console.log('Scale factor y: ', scaleFactorY);
        //
        if ((scaleFactorX != 0) && (scaleFactorY != 0)) {
            touchscreen.SetScaleFactor(1.0 / scaleFactorX, 1.0 / scaleFactorY);
            //
            this.state = STATE_PRODUCTION;
        }
    }

    run() {
        setInterval(() => {
            this.renderScreen();
        }, 1000.0);
        touchscreen.on('touch', (event) => {


            console.log(event);

            if (this.state == STATE_CALIBRATION) {
                this.processCalibration(event);
            } else if (this.state == STATE_PRODUCTION) {
                this.crossX = event.x;
                this.crossY = event.y;
            }
        });
    }

}

const app = new Application();
app.run();

