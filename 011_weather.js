
const screen = require('./bringdesk_screen/build/Release/screen.node');

const weather = require('./widgets/WeatherWidget');
const box = require('./widgets/basic/BoxWidget');

class Application {

    constructor() {
        const options = {
        };
        this.scr1 = new screen.Screen(options);
        this.ww  = new weather.WeatherWidget({
            screen: this.scr1,
        });
        this.ww.start();
        this.box = new box.BoxWidget({
            color: '#bd11f9',

        });
        this.box.child = this.ww;
    }

    renderScreen() {
        const width = this.scr1.width;
        const height = this.scr1.height;

        /* Clear screen */
        this.scr1.Clear();

        /* Draw weather */
        this.box.render({
            screen: this.scr1,
            position: {
                left: 0,
                top: 0,
                width: width,
                height: height,
            },
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
