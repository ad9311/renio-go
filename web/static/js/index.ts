import htmx from 'htmx.org';
import $ from 'jquery';
import income from '@/income';

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

  $('#toggle-sidebar').on('click', function () {
    $('#sidebar-overlay').removeClass('hidden').addClass('sidebar-overlay');
    $('#sidebar-container').removeClass('-right-full').addClass('right-0');
  });

  $('#sidebar-overlay').on('click', function (event) {
    if (!$(event.target).closest('#sidebar-container').length) {
      $('#sidebar-overlay').addClass('hidden').removeClass('sidebar-overlay');
      $('#sidebar-container').addClass('-right-full').removeClass('right-0');
    }
  });

  income();
});
