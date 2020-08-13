
const https = require('https');

class MoneyWidget {

    constructor(options) {
        this._options = options;
        this.moneyUpdateInterval = 5 * 60 * 1000.0;
        this.money = null;
    }

    start() {
        setImmediate(() => {
            this.updateMoney();
        });
        setInterval(() => {
            this.updateMoney();
        }, this.moneyUpdateInterval);
    }

    stop() {
    }

    processMoney(rawData) {
        console.log(rawData);
        this.money = JSON.parse(rawData);
    }

    updateMoney() {
        const url = "https://www.cbr-xml-daily.ru/daily_json.js";
        const request = https.get(url, (res) => {
            res.setEncoding('utf8');
            let rawData = '';
            res.on('data', (chunk) => {
                rawData += chunk;
            });
            res.on('end', () => {
                this.processMoney(rawData);
            });
        });
        request.on('error', (err) => {
            console.log(err);
        });
    }

    searchExchange() {
        let result = null;
        try {
            const money = this.money;
            const valute = money.Valute;
            const USD = valute.USD;
            const value = USD.Value;
            result = value;
        } catch(err) {
            console.log(err);
        }
        return result;
    }

    renderInfo(options) {
        const screen = options.screen;
        const position = options.position;
        //
        screen.SelectFontFace("Helvetica")
        screen.SetFontSize(48.0)
        screen.SetColor(1.0, 1.0, 1.0, 1.0);
        //
        const value = this.searchExchange('USD');
        //
        if (value === null) {
            screen.MoveTo(position.left, position.top + 1 * 48);
            screen.DrawText('Update error.');
        } else {
            screen.MoveTo(position.left, position.top + 1 * 48);
            const line1 = `USD: ${value}`;
            screen.DrawText(line1);
        }
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
        /* Draw weather */
        this.renderInfo(options);
    }

}

module.exports = {
    MoneyWidget
}
