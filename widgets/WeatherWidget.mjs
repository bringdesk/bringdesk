
import https from 'https';

class WeatherWidget {

    constructor(options) {
        this.options = options;
        this.weatherUpdateInterval = 5 * 1000;
        this.weather = {};
        this.images = {};
    }

    setup() {
        this.LoadFont({
            name: 'WeatherFont',
            path: 'Helvetica',
            size: 48,
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

    render(parent) {

        if (this.weather) {

            const {
                temp = 0,
                name = 'Неизвестно',
                description = 'Идет загрузка данных ...',
            } = this.weather;

            /* Рисуем небольшое изображение для прогноза погоды (иконку) */
            //screen.DrawImage(, position.left + 16, position.top + 16);

            /* Рисуем сам прогноз погоды */
            parent.DrawText({
                font: 'WeatherFont',
                text: `${name} ... ${temp} ... ${description}`,
                x: 100,
                y: 100,
                color: [255,255,255],
            });

        }

    }

}

export {
    WeatherWidget
}
