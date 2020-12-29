
import {BringDesk} from './lib';

class Application extends BringDesk {

    constructor() {
        super();
        this.active = true;
    }

    setup() {

        /* Регистрируем шрифт с именем "FreeSans" для написания обьявления */
        this.LoadFont({
            name: 'FreeSans',
            path: './res/font/FreeSans.ttf',
            size: 96,
        });

        /* Добавим мигание */
        setInterval(() => {
            this.active = !this.active;
        }, 250);

    }

    render() {

        const [ w, h ] = this.GetViewport();
        this.FillRect([0, 0, w, h], [0, 0, 0]);

        this.DrawText({
            name: 'FreeSans',
            text: 'На все Кухни',
            color: [255, 255, 0],
            x: 200,
            y: 300,
        });

        if (this.active) {

            this.DrawText({
                name: 'FreeSans',
                text: 'скидка 5%',
                color: [255, 0, 0],
                x: 200,
                y: 600,
            });

        }

        this.DrawText({
            name: 'FreeSans',
            text: '2 этаж',
            color: [0, 255, 0],
            x: 200,
            y: 900,
        });

        /* Обновим изображение на экране */
        this.RenderPresent();

    }

};

const app = new Application();
app.run();
