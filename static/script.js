const APIKEY = '801e17cd-9bb7-4ee0-b671-886be1225e55'

var regform = $("#reg-form");
regform.submit(function(e){
    e.preventDefault();

    $.ajax({
        type: regform.attr('method'),
        url: regform.attr('action'),
        data: regform.serialize(),
        success: function(data){
            console.log(data)
            console.log('success')
        },
        error: function(data){
            console.log(data)
            console.log('error')
        }
    })
})