import { defineConfig, presetAttributify, presetUno } from 'unocss'

export default defineConfig({
  cli: {
    entry: {
      patterns: ['./templates/**/*.tmpl'],
      outFile: './static/style.css',
    },
    // CliEntryItem | CliEntryItem[]
  },
  presets: [presetUno(), presetAttributify()],
  // ...
})
