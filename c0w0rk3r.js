#!/usr/bin/env node

const puppeteer = require('puppeteer-core');

const url = 'https://meet.jit.si/C0MFYC0W0RK1NGCLUB';
const pages = [
  {count: 1, url: 'chrome://webrtc-internals'},
  {count: 0, url: url + '#config.startWithVideoMuted=true&config.startWithAudioMuted=true'},
  {count: 2, url: url}
];

(async (pages) => {
  const browser = await puppeteer.launch({
    executablePath: '/Applications/Google Chrome.app/Contents/MacOS/Google Chrome',
    devtools: true,
    defaultViewport: null,
    args: [
      "--enable-logging=stderr --vmodule=*/webrtc/*=3",
      "--mute-audio",
      "--use-fake-device-for-media-stream",
      "--use-fake-ui-for-media-stream",
      "--use-file-for-fake-video-capture=/Users/cat/Projects/C0W0RK/tet.mjpeg"
    ]
  });

  async function goto(url, count) {
    for (var idx = 0; idx < count; idx++) {
      const page = await browser.newPage();
      await page.goto(url);
    }
  }

  pages.forEach(async p => {
    await goto(p.url, p.count)
  });
  
})(pages);