
import {BringDesk} from './lib';

import {WeatherWidget} from './widgets';

class Application extends BringDesk {

    setup() {

        /* Создание погодного виджета */
//        const weatherWidget = new WeatherWidget({
//        });
//        weatherWidget.start();

        /* Создание коробки для виджета */
//        const boxWidget = BoxWidget({
//            'background-color': '#bd11f9',
//        });
//        boxWidget.SetChildren(weatherWidget);

        /* Регистрируем только коробку */
//        this.RegisterWidget(boxWidget);

    }

}

const app = new Application();
app.run();
