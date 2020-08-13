
const screen = require('./bringdesk_screen/build/Release/screen.node');

const box = require('./widgets/basic/BoxWidget');
const vbox = require('./widgets/layout/VBoxLayoutWidget');

const calendar = require('./widgets/CalendarWidget');
const weather = require('./widgets/WeatherWidget');

class Application {

    constructor() {
        const options = {
        };
        this.scr1 = new screen.Screen(options);
        //
        this.vbox = new vbox.VBoxLayoutWidget({
            screen: this.scr1,
        });
        this.vbox.start();
        //
        const box1 = new box.BoxWidget({});
        const box2 = new box.BoxWidget({});
        //
        const box1_1 = new box.BoxWidget({
            color: '#C000C0',
        });
        const box2_1 = new box.BoxWidget({
            color: '#AA0000',
        });
        //
        const widget1 = new calendar.CalendarWidget({});
        widget1.start();
        const widget2 = new weather.WeatherWidget({});
        widget2.start();
        //
        box1.child = box1_1;
        box2.child = box2_1;
        //
        box1_1.child = widget1;
        box2_1.child = widget2;
        //
        this.vbox.children.push(box1);
        this.vbox.children.push(box2);
    }


    renderScreen() {
        const width = this.scr1.width;
        const height = this.scr1.height;
        /* Clear screen */
        this.scr1.Clear();
        /* Render money widget */
        this.vbox.render({
            screen: this.scr1,
            position: {
                'left': 0,
                'top': 0,
                'width': width,
                'height': height,
            },
            debug: true,
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