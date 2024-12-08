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
  $('#closeModal').on('click', function() {
    $('#modal').addClass('hidden').removeClass('modal');
  });

  income();
});
