
const screen = require('./bringdesk_screen/build/Release/screen.node');

class Application {

    constructor() {
        const options = {
        };
        this.scr1 = new screen.Screen(options);
        this._active = true;
    }

    renderContent(options) {
        const screen = options.screen;
        //
        screen.SetFontSize(96);
        //
        screen.SetColor(1.0, 1.0, 0.0, 1.0);
        this.scr1.MoveTo(200, 300);
        this.scr1.DrawText('На все Кухни');
        //
        if (this._active) {
            screen.SetColor(1.0, 0.0, 0.0, 1.0)
            this.scr1.MoveTo(200, 600);
            this.scr1.DrawText('скидка 5%');
        }
        this._active = !this._active;
        //
        screen.SetColor(0.0, 1.0, 0.0, 1.0)
        this.scr1.MoveTo(200, 900);
        this.scr1.DrawText('2 этаж');
    }

    renderScreen() {
        /* Get screen size */
        const width = this.scr1.width;
        const height = this.scr1.height;
        /* Clear screen */
        this.scr1.Clear();
        /* Draw red background */
        this.scr1.SetColor(0.0, 0.0, 0.5, 1.0);
        this.scr1.DrawRectangle(0, 0, width, height);
        /* Render content */
        this.renderContent({
            screen: this.scr1,
        });
        /* Swap image */
        this.scr1.Swap();
    }

    run() {
        setInterval(() => {
            this.renderScreen();
        }, 250.0);
    }

};

const app = new Application();
app.run();
