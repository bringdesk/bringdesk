
import {BringDesk} from './lib';

class Application extends BringDesk {

    renderCenterText(posY, msg) {
        const width = this.scr1.width;
        const height = this.scr1.height;
//      const window = this.scr1.GetTextResolution(msg);
        const window = {
            width: 1260,
            height: 75,
        };
        const posX = (width - window.width) / 2.0;
        this.scr1.MoveTo(posX, posY);
        this.scr1.DrawText(msg);
    }

    GetViewport() {
        return [ 0, 0 ];
    }

    render() {

        /* Get viewport */
        const [ w, h ] = this.GetViewport();

        /* Clear screen */
        this.Clear();

        /* Draw background */
//        this.scr1.SetColor(1.0, 0.0, 0.0, 1.0);
//        this.scr1.DrawRectangle(0, 0, width, height);

        /* Calculate text window */
//        const window = {}
//        window.top = 0 || 300;
//        const singleLinePadding = 25
//        const singleLineHeight = 75

        /* Draw message */
//        this.renderCenterText(window.top + 1 * singleLinePadding + 0 * singleLineHeight, 'ВНИМАНИЕ !')
//        this.renderCenterText(window.top + 2 * singleLinePadding + 1 * singleLineHeight, 'Сохраняйте спокойствие.')
//        this.renderCenterText(window.top + 3 * singleLinePadding + 2 * singleLineHeight, 'Всем срочно покинуть территорию Арены.')
//        this.renderCenterText(window.top + 4 * singleLinePadding + 3 * singleLineHeight, 'Следуйте планам эвакуации !')

        //Update screen
        this.RenderPresent();

    }

    setup() {
        /* Preload fonts */
        this.LoadFont({
            name: 'FreeSans',
            path: './res/font/FreeSans.ttf',
            size: 24,
        });
    }

}

const app = new Application();
app.run();
