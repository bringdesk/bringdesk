
const https = require('https');

const screen = require('./bringdesk_screen/build/Release/screen.node');

const calendar = require('./widgets/CalendarWidget');
const box = require('./widgets/basic/BoxWidget');

class Application {

    constructor() {
        const options = {
        };
        this.scr1 = new screen.Screen(options);
        this.cw = new calendar.CalendarWidget({
            screen: this.scr1,
            'font-size': 96,
        });
        this.cw.start();
        this.box = new box.BoxWidget({
            color: '#f9bd11',

        });
        this.box.child = this.cw;
    }


    renderScreen() {
        const width = this.scr1.width;
        const height = this.scr1.height;
        /* Clear screen */
        this.scr1.Clear();
        /* Render money widget */
        this.box.render({
            screen: this.scr1,
            position: {
                left: 0,
                top: 0,
                width: width,
                height: height,
            }
        });
        this.scr1.Swap();
    }

    run() {
        setInterval(() => {
            this.renderScreen();
        }, 1000.0);

    }

}

const app = new Application();
app.run();