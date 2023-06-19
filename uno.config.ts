import {
  defineConfig,
  presetAttributify,
  presetUno,
  presetWebFonts,
} from 'unocss'

export default defineConfig({
  cli: {
    entry: {
      patterns: ['./src/templates/**/*.html'],
      outFile: './src/static/style.css',
    },
  },
  presets: [
    presetUno(),
    presetAttributify(),
    presetWebFonts({
      provider: 'google',
      fonts: {
        sans: 'Roboto',
      },
    }),
  ],

  shortcuts: {
    card: 'block rounded-lg bg-main px-4 py-2 shadow-lg',
    chip: 'bg-background shadow-lg rounded-lg pa-1',
    button:
      'rounded-lg px-4 py-2 text-center shadow-lg bg-main hover:opacity-80 transition text-md font-medium cursor-pointer',
    link: 'rounded-lg border px-2 py-1 bg-accent border-primary',
    input: 'rounded-lg shadow p-2 bg-main',
  },
  theme: {
    colors: {
      background: '#fef6e4',
      main: '#f3d2c1',
      primary: '#0d0d0d',
      accent: '#dfc4ae',
    },
  },
})
