
const process = require('process');

const screen = require('./bringdesk_screen/build/Release/screen.node');

const fs = require('fs');
const YAML = require('yaml');

var parse = require('parse-color');

const calendar = require('./widgets/CalendarWidget');
const money = require('./widgets/MoneyWidget');
const weather = require('./widgets/WeatherWidget');

class Application extends BringDesk {

    constructor() {
        const options = {
        };
        this.scr1 = new screen.Screen(options);
        //
        this.exitTimeout = 10 * 60.0 * 1000.0;
        //
        this.widgets = [];
        /* Create Calendar widget */
        this.widgetImplementations = {
            'CalendarWidget': calendar.CalendarWidget,
            'MoneyWidget': money.MoneyWidget,
            'WeatherWidget': weather.WeatherWidget,
        };
    }

    processConfig() {
        const settings = fs.readFileSync('./settings.yml', 'utf8')
        const params = YAML.parse(settings);
        console.log(params);
        const widgets = params.widgets;
        for (const widgetName in widgets) {
            console.log(widgetName);
            const widgetParams = widgets[widgetName];
            console.log(widgetParams);
            const widgetImplementation = widgetParams.implementation;
            console.log(widgetImplementation);
            const widgetType = this.widgetImplementations[widgetImplementation];
            console.log(widgetType);
            const widgetColor = parse(widgetParams.color);
            console.log(widgetColor);
            const widgetPosition = widgetParams.position;
            console.log(widgetPosition);
            const colorRed   = widgetColor.rgba[0] / 255.0;
            const colorGreen = widgetColor.rgba[1] / 255.0;
            const colorBlue  = widgetColor.rgba[2] / 255.0;
            const w = new widgetType({
                screen: this.scr1,
                color: [colorRed, colorGreen, colorBlue, 1.0],
                position: widgetPosition,
            });
            w.start();
            this.widgets.push(w);
        }
    }

    renderScreen() {
        const width = 1024;
        const height = 768;
        /* Clear screen */
        this.scr1.Clear();
        /* Render money widget */
        for (const widgetIndex in this.widgets) {
            const widget = this.widgets[widgetIndex];
            widget.render({});
        }
        this.scr1.Swap();
    }

    run() {
        this.processConfig();
        setInterval(() => {
            this.renderScreen();
        }, 1000.0);
        setTimeout(() => {
            process.exit();
        }, this.exitTimeout);
    }

}

const app = new Application();
app.run();
