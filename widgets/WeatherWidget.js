
const https = require('https');

class WeatherWidget {

    constructor(options) {
        this._options = options;
        this.weatherUpdateInterval = 5 * 60 * 1000.0;
        this.weather = null;
    }

    processWeather(rawData) {
        console.log(rawData);
        this.weather = JSON.parse(rawData);
    }

    updateWeather() {
        const url = "https://api.openweathermap.org/data/2.5/weather?id=498817&lang=ru&units=metric&appid=62a65b29ff2592740c31b0022281cc33";
        const request = https.get(url, (res) => {
            res.setEncoding('utf8');
            let rawData = '';
            res.on('data', (chunk) => {
                rawData += chunk;
            });
            res.on('end', () => {
                this.processWeather(rawData);
            });
        });
        request.on('error', (err) => {
            console.log(err);
        });
    }

    start() {
        setImmediate(() => {
            this.updateWeather();
        });
        setInterval(() => {
            this.updateWeather();
        }, this.weatherUpdateInterval);
    }

    stop() {
    }

    searchWeather() {
        let result = null;
        try {
            const weathers = this.weather.weather;
            const weather = weathers[0];
            const description = weather.description;
            const main = this.weather.main;
            const name = this.weather.name;
            result = {
                temp: main.temp,
                name: name,
                description: description,
            };
        } catch(err) {
            console.log(err);
        }
        return result;
    }

    renderImage(options) {
        const screen = options.screen;
        const position = options.position;
        //
        screen.DrawImage("./res/weather/04n.png", position.left + 16, position.top + 16);
    }

    renderError(options) {
    }

    renderInfo(options) {
        const screen = options.screen;
        const position = options.position;
        //
        screen.SelectFontFace("Helvetica")
        screen.SetFontSize(48.0)
        screen.SetColor(1.0, 1.0, 1.0, 1.0);
        //
        const weather = this.searchWeather();
        //
        if (weather === null) {
            screen.MoveTo(position.left + 64, position.top + 1 * 48);
            screen.DrawText("Update error.");
        } else {
            screen.MoveTo(position.left + 64, position.top + 1 * 48);
            const line1 = String(weather.name);
            screen.DrawText(line1);
            //
            screen.MoveTo(position.left + 64, position.top + 2 * 48);
            const line2 = `${weather.temp} C`;
            screen.DrawText(line2);
            //
            screen.MoveTo(position.left + 64, position.top + 3 * 48);
            const line3 = String(weather.description);
            screen.DrawText(line3);
        }
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
        /* Draw information */
        this.renderImage(options);
        this.renderInfo(options);
    }

}

module.exports = {
    WeatherWidget
}
