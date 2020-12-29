
import {BringDesk} from './lib';

class Application extends BringDesk {

//    constructor() {
//        const options = {
//        };
//        this.scr1 = new screen.Screen(options);
//        this.mw = new money.MoneyWidget({
//            screen: this.scr1,
//        });
//        this.mw.start();
//        this.box = new box.BoxWidget({
//            color: '#11bdf9',
//        });
//        this.box.child = this.mw;
//    }

//    renderScreen() {
//        const width = this.scr1.width;
//        const height = this.scr1.height;
//        this.scr1.Clear();
//        this.box.render({
//            screen: this.scr1,
//            position: {
//                left: 0,
//                top: 0,
//                width: width,
//                height: height,
//            }
//        });
//        this.scr1.Swap();
//    }

}

const app = new Application();
app.run();
