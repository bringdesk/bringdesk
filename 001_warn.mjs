
import {BringDesk} from './lib';

class Application extends BringDesk {

    setup() {

        /* Регистрируем шрифт с именем "FreeSans" для написания обьявления */
        this.LoadFont({
            name: 'FreeSans',
            path: './res/font/FreeSans.ttf',
            size: 36,
        });

    }

    render() {

        /* Получаем разрешение поверхности экрана */
        const [ w, h ] = this.GetViewport();

        /* Очищаем изображение (смысла в этом без изменений нет, но тут сделано для примера) */
//        this.Clear();

        /* Закрашиваем красный фон */
        this.FillRect([0, 0, w, h], [255, 0, 0]);

        /* Выводим строки текста (сложный способ с вычислением координат текста) */
        const lines = [
            'ВНИМАНИЕ !',
            'Сохраняйте спокойствие.',
            'Всем срочно покинуть территорию Арены.',
            'Следуйте планам эвакуации !',
        ];
        let posX = 120;
        let posY = 300;
        lines.forEach((line) => {
            this.DrawText({
                name: 'FreeSans',
                text: line,
                color: [255, 255, 255],
                x: posX,
                y: posY,
            });
            posY += 80;
        });

        /* Обновим изображение на экране */
        this.RenderPresent();

    }

}

const app = new Application();
app.run();
