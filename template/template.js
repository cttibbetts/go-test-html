window.onload = function() {
    var headers = document.querySelectorAll('.test-section-header');

    Array.from(headers).forEach(function (link) {
        link.addEventListener('click', function (_) {
            link.parentNode.classList.toggle('hidden-section');
        });
    });

    var localRunText = document.querySelector('.local-run');

    localRunText.addEventListener('click', function(event) {
        this.select();

        try {
            var successful = document.execCommand('copy');
            if (successful) {
                var copySuccess = document.querySelector('.copy-success');
                copySuccess.classList.add('mdl-snackbar--active');
                setTimeout(function() {
                    copySuccess.classList.remove('mdl-snackbar--active');
                }, 1500);
            }
        } catch (err) {
            console.log('Oops, unable to copy');
        }
    });
}
