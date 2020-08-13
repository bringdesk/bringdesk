
class BringDesk {

    renderScreen() {
    }

    run() {
        const renderInterval = 1000 / 30;
        setInterval(() => {
            this.renderScreen();
        }, renderInterval);
    }

}

export {
    BringDesk
}
