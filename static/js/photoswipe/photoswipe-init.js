import PhotoSwipeLightbox from '/js/photoswipe/photoswipe-lightbox.esm.js';

const lightbox = new PhotoSwipeLightbox({
  gallery: '.pswp-gallery',
  children: 'a',
  pswpModule: () => import('/js/photoswipe/photoswipe.esm.js')
});

lightbox.init();
