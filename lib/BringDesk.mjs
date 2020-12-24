
import SDL2 from 'napi-sdl2';

class BringDesk {

    constructor() {
        this.window = null;
        this.renderer = null;
        this.font = null;
    }

    LoadFont(options = {}) {
        const{name, path, size = 24} = options;
        this.font = SDL2.TTF_OpenFont(path, size);
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

    RenderPresent() {
        SDL2.SDL_RenderPresent( this.renderer );
    }

    render() {
        this.Clear();
        this.RenderPresent();
    }

    run() {

        SDL2.SDL_Init(SDL2.SDL_INIT_EVERYTHING);
        SDL2.TTF_Init();

        this.setup();

        const numVideoDisplays = SDL2.SDL_GetNumVideoDisplays();
        console.log(`Number video displays: ${numVideoDisplays}`);

        const displayIndex = 0;
        const [w, h] = SDL2.SDL_GetDesktopDisplayMode(displayIndex);

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

        while(!quit)
        {
            const event = {};
            const ret = SDL2.SDL_PollEvent(event);
            if (ret == 1) {
                //console.log(`Event: Type = ${event.type}`);
                if (event.type === 'QUIT')
                {
                    quit = true;
                    break;
                }
            }

            this.render();
            SDL2.SDL_Delay(100);

        }

        SDL2.TTF_Quit();
        SDL2.SDL_DestroyRenderer(this.renderer);
        SDL2.SDL_DestroyWindow(this.window);
        SDL2.SDL_Quit();

    }

}

export {
    BringDesk
}
