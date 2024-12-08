import $ from 'jquery';

export default () => {
  $('#deleteIncome').on('click', function () {
    $('#modal').removeClass('hidden').addClass('modal');
  });
};
