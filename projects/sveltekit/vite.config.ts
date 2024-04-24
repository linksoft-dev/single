import {sveltekit} from '@sveltejs/kit/vite';
import {defineConfig} from 'vite';
import commonjs from 'vite-plugin-commonjs'

export default defineConfig({
    plugins: [sveltekit(), commonjs(/* options */{
        filter(id) {
            // `node_modules` is exclude by default, so we need to include it explicitly
            // https://github.com/vite-plugin/vite-plugin-commonjs/blob/v0.7.0/src/index.ts#L125-L127
            if (id.includes('node_modules/comps')) {
                return true
            }
        },
        dynamic: {
            loose: true,
        }
    }),]
});
