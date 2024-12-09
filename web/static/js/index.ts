import htmx from 'htmx.org';
import $ from 'jquery';

/* eslint-disable @typescript-eslint/no-explicit-any */
declare global {
  interface Window {
    htmx: any;
  }
}
/* eslint-enable @typescript-eslint/no-explicit-any */

window.htmx = htmx;

$(() => {
  $('#closeModal').on('click', function () {
    $('#modal').addClass('hidden').removeClass('modal-overlay');
  });

  function toggleSidebar() {
    $('#sidebar-overlay').toggleClass('sidebar-overlay');
    $('#sidebar-container').toggleClass('translate-x-full');
  }

  $('#toggle-sidebar, .close-sidebar').on('click', function () {
    toggleSidebar();
  });

  $('#sidebar-overlay').on('click', function (event) {
    if (!$(event.target).closest('#sidebar-container').length) {
      toggleSidebar();
    }
  });

  $('#deleteEntry').on('click', function () {
    $('#modal').removeClass('hidden').addClass('modal-overlay');
  });
});
