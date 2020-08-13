
import BringDesk from './lib/BringDesk';

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

    renderScreen() {
        /* Get screen size */
        const width = this.scr1.width;
        const height = this.scr1.height;
        /* Clear screen */
        this.scr1.Clear();
        /* Draw red background */
        this.scr1.SetColor(1.0, 0.0, 0.0, 1.0);
        this.scr1.DrawRectangle(0, 0, width, height);
        /* Set canvas parameters */
        this.scr1.SetFontFace("Terminus")
        //this.scr1.SetFontFace("Helvetica")
        this.scr1.SetFontSize(64.0)
        this.scr1.SetColor(1.0, 1.0, 1.0, 1.0)
        /* Calculate text window */
        const window = {}
        window.top = 0 || 300; // TODO - calculate height of text minus srceen resolution and etc...
        const singleLinePadding = 25
        const singleLineHeight = 75
        /* Draw message */
        this.renderCenterText(window.top + 1 * singleLinePadding + 0 * singleLineHeight, 'ВНИМАНИЕ !')
        this.renderCenterText(window.top + 2 * singleLinePadding + 1 * singleLineHeight, 'Сохраняйте спокойствие.')
        this.renderCenterText(window.top + 3 * singleLinePadding + 2 * singleLineHeight, 'Всем срочно покинуть территорию Арены.')
        this.renderCenterText(window.top + 4 * singleLinePadding + 3 * singleLineHeight, 'Следуйте планам эвакуации !')
        /* Swap image */
        this.scr1.Swap();
    }

}

const app = new Application();
app.run();
