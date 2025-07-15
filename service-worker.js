// service-worker.js
self.addEventListener('install', event => {
  console.log('✅ Service Worker Installed');
  self.skipWaiting(); // Activate worker immediately
});

self.addEventListener('activate', event => {
  console.log('✅ Service Worker Activated');
  return self.clients.claim();
});

self.addEventListener('fetch', event => {
  // Always go to network first
  event.respondWith(
    fetch(event.request).catch(() => {
      return caches.match(event.request);
    })
  );
});
