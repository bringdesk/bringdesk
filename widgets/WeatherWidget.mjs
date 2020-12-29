
import https from 'https';

class WeatherWidget {

    constructor(options) {
        this.options = options;
        this.weatherUpdateInterval = 5 * 1000;
        this.weather = null;
        this.images = {};
    }

    setup() {
        this.LoadFont({
            name: 'WeatherFont',
            path: '...',
            size: 24,
        });
        this.images = {
            Cloud: this.parent.LoadImage('./res/weather/04n.png'),
        };
    }

    processResponse(rawData) {
        try {
            const data = JSON.parse(rawData);
            this.weather = this.processWeather(data);
        } catch (err) {
            console.log(`Error process response: err = ${err}`);
        }
    }

    processWeather(data) {
        const result = {};
        try {
            const weathers = data.weather;
            const weather = weathers[0];
            const description = weather.description;
            const main = data.main;
            const name = data.name;
            result.temp = main.temp;
            result.name = name;
            result.description = description;
        } catch(err) {
            console.log(`Error process weather data: err = ${err}`);
        }
        return result;
    }

    updateWeather() {
        const {
            host = 'api.openweathermap.org',
            path = '/data/2.5/weather',
            query = {
                id: 498817,
                lang: 'ru',
                units: 'metric',
                appid: '62a65b29ff2592740c31b0022281cc33'
            }
        } = this.options;
        const url = '...';
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
        this.weatherTimer = setInterval(() => {
            this.updateWeather();
        }, this.weatherUpdateInterval);
    }

    stop() {
        // this.weatherTimer
    }


    renderImage(options) {
        screen.DrawImage(, position.left + 16, position.top + 16);
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

    render(parent) {
        this.renderImage(options);
        this.renderInfo(options);
    }

}

export {
    WeatherWidget
}
