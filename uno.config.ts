import { defineConfig, presetAttributify, presetUno } from 'unocss'

export default defineConfig({
  cli: {
    entry: {
      patterns: ['./src/templates/**/*.html'],
      outFile: './src/static/style.css',
    },
    // CliEntryItem | CliEntryItem[]
  },
  presets: [presetUno(), presetAttributify()],
  // ...
})
