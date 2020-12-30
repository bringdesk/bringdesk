
import {BringDesk} from './lib';

import {BoxWidget} from './widgets';

class Application extends BringDesk {

    setup() {

        /* Создание виджета коробки */
        const boxWidget = new BoxWidget({
            background: '#bd11f9',
            rect: [0, 0, 100, 100],
        });

        /* Регистрируем коробку */
        this.RegisterWidget(boxWidget);

    }

}

const app = new Application();
app.run();
