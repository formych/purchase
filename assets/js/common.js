var $download = document.getElementById("download")

// $download.click()
$download.onclick = function(){
    window.open("/excel")
};
$.ajax({
    url: '/v1/list',
    type: 'GET',
    cache: true,
    dataType: 'json',
    crossDomain: true,
    success: function (data) {
        let html = "";
        data.data.forEach(element => {
            html += "<tr>"  
            + "<td>" + element.User + "</td>"
            + "<td>" + element.Company + "</td>"
            + "<td>" + element.Tel + "</td>"
            + "<td>" + element.PurchaseNum + "</td>"
            + "<td>" + element.PurchaseTime + "</td>"
            + "</tr>";
        });
        pageNum = data.result_total / data.page_total + 1;
        pages = "<ul class='pagination'>"
        for(i = 1; i <= pageNum; i++) {
            pages += "<li><a href=''>" + i +  "</a></li>";
        }
        pages += "</ul>"
        // a(href='results?#{query}&p=#{i}') #{i + 1}
        $('.panel-default tbody').append(html)
        $('.panel-default ').append(pages)

    //   $('#inputTitle').val(data.title)
    //   $('#inputDirector').val(data.directors[0].name)
    //   $('#inputCountry').val(data.countries[0])
    //   $('#inputPoster').val(data.images.large)
    //   $('#inputYear').val(data.year)
    //   
    }
})
