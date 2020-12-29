
const dateFormat = require('dateformat');
dateFormat.i18n = {
    dayNames: [
        'Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat',
        'Воскресенье', 'Понедельник', 'Вторник', 'Среда', 'Четверг', 'Пятница', 'Суббота'
    ],
    monthNames: [
        'Янв', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Окт', 'Ноя', 'Дек',
        'Январь', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'Ноябрь', 'Декабрь'
    ],
    timeNames: [
        'a', 'p', 'am', 'pm', 'A', 'P', 'AM', 'PM'
    ]
};

class CalendarWidget {

    constructor(options) {
        this._options = options;
        this._active = true;
    }

    start() {
    }

    stop() {
    }

    searchDate() {
        const now = new Date();
        const baseFormat = 'dddd, mmmm d, yyyy';
        const result = dateFormat(now, baseFormat);
        return result;
    }

    searchDate2() {
        const date = new Date();
        const day = date.getDate();
        const month = date.getMonth() + 1;
        const year = date.getFullYear();
        //
        const hour = date.getHours();
        const min = date.getMinutes();
        const result = `${year}-${month}-${day}`;
        return result;
    }

    searchTime2() {
        const date = new Date();
        //
        const _hour = String(date.getHours());
        const _min = String(date.getMinutes());
        //
        const hour = _hour.padStart(2, '0');
        const min = _min.padStart(2, '0');
        //
        let result;
        if (this._active) {
            result = `${hour}:${min}`;
        } else {
            result = `${hour} ${min}`;
        }
        return result;
    }

    renderInfo(options) {
        const screen = options.screen;
        const position = options.position;
        const fontSize = options['font-size'] || 48.0;
        //
        screen.SelectFontFace("Courier")
        screen.SetFontSize(fontSize)
        screen.SetColor(0.0, 0.0, 0.0, 1.0);
        //
        const line1 = this.searchDate();
        screen.MoveTo(position.left, position.top + 1*fontSize);
        screen.DrawText(line1);
        //
        const line2 = this.searchTime2();
        screen.MoveTo(position.left, position.top + 2*fontSize);
        screen.DrawText(line2);
        //
        this._active = !this._active;
    }

    mergeOptions(mainOptions, currentOptions) {
        const options = {};
        /* Overlay current options */
        Object.keys(currentOptions).forEach((key) => {
            const value = currentOptions[key];
            options[key] = value;
        });
        /* Overlay main options */
        Object.keys(mainOptions).forEach((key) => {
            const value = mainOptions[key];
            options[key] = value;
        });
        /* Done */
        return options;
    }

    render(options1) {
        const options = this.mergeOptions(this._options, options1);
        const screen = options.screen;
        /* Draw value */
        this.renderInfo(options);
    }

}

export {
    CalendarWidget
}
