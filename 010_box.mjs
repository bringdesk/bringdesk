
import {BringDesk} from './lib';

import {BoxWidget} from './widgets';

class Application extends BringDesk {

    setup() {

        /* Создание виджета коробки */
        const boxWidget = new BoxWidget({
            'background-color': '#bd11f9',
        });

        /* Регистрируем коробку */
        this.RegisterWidget(boxWidget);

    }

}

const app = new Application();
app.run();
