
import https from 'https';

class WeatherWidget {

    constructor(options) {
        this.options = options;
        this.weatherUpdateInterval = 5 * 1000;
        this.weather = {};
        this.images = {};
    }

    startUpdate() {
        setImmediate(() => {
            this.updateWeather();
        });
        this.weatherTimer = setInterval(() => {
            this.updateWeather();
        }, this.weatherUpdateInterval);
    }
    stopUpdate() {
    }

    setup(parent) {
        /* Загружаем шрифт для использования в компоненте прогноза погоды */
        parent.LoadFont({
            name: 'WeatherFont',
            path: './res/font/FreeSans.ttf',
            size: 48,
        });
        /* Загружаем изображения */
        this.images = {
            Cloud: parent.LoadImage('./res/weather/04n.png'),
        };
        /* Стартуем рутину обновления данных о прогнозе погоды */
        this.startUpdate();
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
        const address = new URL('https://api.openweathermap.org/data/2.5/weather');
        address.searchParams = {
            id: 498817,
            lang: 'ru',
            units: 'metric',
            appid: '62a65b29ff2592740c31b0022281cc33'
        };
        console.log(`Weather server address is ${address}`);
        const request = https.get(address, (res) => {
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

    render(parent) {

        const {
            temp = 0,
            name = 'Неизвестно',
            description = 'Идет загрузка данных ...',
        } = this.weather;

        /* Рисуем небольшое изображение для прогноза погоды (иконку) */
        //screen.DrawImage(, position.left + 16, position.top + 16);

        /* Рисуем сам прогноз погоды */
        parent.DrawText({
            fontName: 'WeatherFont',
            text: `${name} ... ${temp} ... ${description}`,
            x: 100,
            y: 100,
            color: [255,255,255],
        });

    }

}

export {
    WeatherWidget
}
