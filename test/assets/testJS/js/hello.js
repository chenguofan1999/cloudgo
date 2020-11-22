$(document).ready(function() {
    $.ajax({
        url: "/js"
    }).then(function(data) {
       $('.greeting-id').append(data.id);
       $('.greeting-content').append(data.content);
    });
});