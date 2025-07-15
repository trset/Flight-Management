// service-worker.js
self.addEventListener('install', e => {
  console.log('Service Worker Installed');
});

self.addEventListener('fetch', e => {
  // No caching logic for now (optional)
});
