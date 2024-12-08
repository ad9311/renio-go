import htmx from 'htmx.org'
import $ from 'jquery';
import income from './income'

window.htmx = htmx

$(() => {
  $('#closeModal').on('click', function() {
    $('#modal').addClass('hidden').removeClass('modal');
  });

  income();
})
