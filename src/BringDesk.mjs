
import SDL2 from '@vit1251/napi_sdl2';

class BringDesk {

    constructor() {
        this.window = null;
        this.renderer = null;
        this.fonts = {};
        this.widgets = [];
        this.width = 0;
        this.height = 0;
    }

    LoadFont(options = {}) {
        const{name, path, size = 24} = options;
        this.fonts[name] = SDL2.TTF_OpenFont(path, size);
    }
    LoadImage(options = {}) {
    }

    /**
     * Register widget
     */
    RegisterWidget(widget) {
        widget.setup(this);
        this.widgets.push(widget);
    }

    setup() {
    }

    /**
     * Clear screen
     */
    Clear() {
        SDL2.SDL_SetRenderDrawColor( this.renderer, [0xFF, 0xFF, 0xFF, 0xFF] );
        SDL2.SDL_RenderClear( this.renderer );
    }

    FillRect(rect, color) {
        SDL2.SDL_SetRenderDrawColor(this.renderer, color);
        SDL2.SDL_RenderFillRect(this.renderer, rect);
    }

    DrawText(options = {}) {

        const {
            fontName,
            text = 'Hello, world.',
            color = [128, 128, 128],
            x = 0,
            y = 0,
        } = options;

        const {
            font = this.fonts[fontName],
        } = options;

        if (font) {

            const msgSurface = SDL2.TTF_RenderUTF8_Solid(font, text, color);
            const newTexture = SDL2.SDL_CreateTextureFromSurface(this.renderer, msgSurface);
            const [msgWidth, msgHeight] = SDL2.SDL_QueryTexture(newTexture);
            const newRect = [x, y, msgWidth, msgHeight];
            SDL2.SDL_RenderCopy(this.renderer, newTexture, null, newRect);
            SDL2.SDL_FreeSurface(msgSurface);
            SDL2.SDL_DestroyTexture(newTexture);

        } else {
            console.warn(`No font in DrawText.`);
        }

    }

    RenderPresent() {
        SDL2.SDL_RenderPresent( this.renderer );
    }

    render() {
        this.Clear();
        this.widgets.forEach((widget) => {
            widget.render(this);
        });
        this.RenderPresent();
    }

    GetViewport() {
        return [ this.width, this.height ];
    }

    run() {

        SDL2.SDL_Init(SDL2.SDL_INIT_EVERYTHING);
        SDL2.TTF_Init();

        this.setup();

        const numVideoDisplays = SDL2.SDL_GetNumVideoDisplays();
        console.log(`Number video displays: ${numVideoDisplays}`);

        const displayIndex = 0;
        const [w, h] = SDL2.SDL_GetDesktopDisplayMode(displayIndex);
        this.width = w;
        this.height = h;

        console.log(`Display resolution: ${w} x ${h}`);

        const [screen_width, screen_height] = [w, h];

        const sdl_window = SDL2.SDL_CreateWindow("SDL Sample", 0, 0, screen_width, screen_height, SDL2.SDL_WINDOW_FULLSCREEN);
        const sdl_renderer = SDL2.SDL_CreateRenderer( sdl_window, -1, SDL2.SDL_RENDERER_ACCELERATED );

        //Initialize renderer color

        SDL2.SDL_SetRenderDrawColor(sdl_renderer, [0xFF, 0xFF, 0xFF, 0xFF]);

        let quit = false;

        this.window = sdl_window;
        this.renderer = sdl_renderer;

        this.render();

        const cursor = SDL2.SDL_CreateSystemCursor(SDL2.SDL_SYSTEM_CURSOR_ARROW);
        SDL2.SDL_SetCursor(cursor);
        SDL2.SDL_ShowCursor(1);

        this.mainTimer = setInterval(() => {

            const event = {};
            const ret = SDL2.SDL_PollEvent(event);
            if (ret == 1) {
                //console.log(`Event: Type = ${event.type}`);
                if (event.type === 'QUIT')
                {
                    this.Stop();
                }
            }

            this.render();

        }, 100);

    }

    Stop() {

        /* Stop */
        cancelInterval(this.mainTimer);

        /* Release resources */
        SDL2.TTF_Quit();
        SDL2.SDL_DestroyRenderer(this.renderer);
        SDL2.SDL_DestroyWindow(this.window);
        SDL2.SDL_Quit();

    }

}

export {
    BringDesk
}
