import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import { VitePWA } from 'vite-plugin-pwa'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    svelte(), 
    VitePWA({
      //includeAssets: ['./assets/favicon.ico', './assets/apple-touch-icon.png', './assets/masked-icon.svg'],
      manifest: {
        name: "ONECard Dashboard",
        short_name: "OCD",
        description: "OneCARD data easily accesible for UVIC students",
        theme_color: "#ffffff",
        
        icons: [
          {
            "src": "/android-chrome-192x192.png",
            "sizes": "192x192",
            "type": "image/png"
          },
          {
            "src": "/android-chrome-512x512.png",
            "sizes": "512x512",
            "type": "image/png"
          }
        ]
      }
    })
  ]
})
