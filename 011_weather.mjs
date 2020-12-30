
import {BringDesk} from './lib';

import {BoxWidget} from './widgets';
import {WeatherWidget} from './widgets';

class Application extends BringDesk {

    setup() {

        /* Создание погодного виджета */
        const weatherWidget = new WeatherWidget({
        });

        /* Создание коробки для виджета */
        const boxWidget = new BoxWidget({
            background: '#bd11f9',
            rect: [100, 100, 200, 50],
            content: weatherWidget,
            padding: [5, 5, 5, 5],
        });

        /* Регистрируем только коробку */
        this.RegisterWidget(boxWidget);

    }

}

const app = new Application();
app.run();
