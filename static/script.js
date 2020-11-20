var regform = $("#reg-form");
regform.submit(function (e) {
    e.preventDefault();
    var dataArray = $(this).serializeArray(),
        dataObj = {};

    $(dataArray).each(function (i, field) {
        dataObj[field.name+""] = field.value;
    });

    var cleanData = JSON.stringify(dataObj)

    $.ajax({
        type: regform.attr('method'),
        url: regform.attr('action'),
        data: cleanData,
        success: function(data){
            alert("your api key is : "+data["msg"])
        },
        error: function(req, status, err){
            alert(JSON.parse(req.responseText)["msg"])
        }
    })


})