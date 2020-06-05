var div = document.getElementById("result");

var newp = document.createElement("p");

var domain = window.location.host;
var protocol = window.location.protocol;
console.log(protocol)
console.log(domain)

var host_url = protocol + "//" + domain

function send_request() {

    var xhr = null; //先设置xhr为空，为了轮询时再次调用函数对xhr重用，引发错误

    xhr = new XMLHttpRequest();
    var url = document.getElementById("address").value;

    xhr.open('POST', '/url?url=' + url, true); //第三个参数一定要设置为true，异步不阻塞，不会影响到后面JS的执行。

    xhr.send();

    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4 && xhr.status == 200) {
            var obj = JSON.parse(xhr.responseText)
            console.log(obj)
            if (obj['code'] == 1) {
                console.log("处理完成!")
                short_url = host_url + "/" + obj['data']
                new_link = "<input id='result_text' value=" + short_url + " > "
                newp.innerHTML = new_link;
                div.appendChild(newp);
            } else {
                newp.innerHTML = obj["msg"];
                div.appendChild(newp);
            }
        }
    };
}

function clear_ip() {
    console.log("clear ok!")
    document.getElementById("address").value = ""
}

function copy_result_text() {
    res_context = document.getElementById("result_text")
    res_context.select()
    document.execCommand("copy")
}