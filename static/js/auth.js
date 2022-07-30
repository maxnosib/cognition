var host = window.location.origin
var sutki_in_milisecond = 86400000

function login(){
    var nik = document.getElementById("form_login_nik").value;
    var pwd = document.getElementById("form_login_pwd").value;
    var resp = document.getElementById('response');

    if (nik == "" || pwd == ""){
        resp.innerHTML = "Заполните поля nik и pwd";
        return
    }

    let req = new XMLHttpRequest();
    req.open("POST", host+"/auth");

    req.setRequestHeader("Content-Type", "application/json");

    req.onload  = function() {
        if (req.status != 200){
            resp.innerHTML = "Попробуйте зарегистрироваться";
            document.getElementById("form_login").hidden = true;
            document.getElementById("form_register").hidden = false;
            return
        }

        var curent_time = new Date().getTime()
        localStorage.setItem("isLogin", curent_time+sutki_in_milisecond);
        resp.innerHTML = "Вы успешно залогинились";
        window.location = '/';
     };

    let data = `{
        "nik":"`+nik+`",`+
        `"pwd":"`+pwd+`"`+
    `}`;

    req.send(data);
}


function register(){
    var nik = document.getElementById("form_register_nik").value;
    var pwd = document.getElementById("form_register_pwd").value;
    var resp = document.getElementById('response');

    if (nik == "" || pwd == ""){
        resp.innerHTML = "Заполните поля nik и pwd";
        return
    }

    let req = new XMLHttpRequest();
    req.open("POST", host+"/register");

    req.setRequestHeader("Content-Type", "application/json");

    req.onload  = function() {
        if (req.status != 200){
            resp.innerHTML = "Попробуйте еще раз";
            return
        }

        var curent_time = new Date().getTime()
        localStorage.setItem("isLogin", curent_time+sutki_in_milisecond);
        resp.innerHTML = "Вы успешно зарегистрировались";
        window.location = '/';
     };

    let data = `{
        "nik":"`+nik+`",`+
        `"pwd":"`+pwd+`"`+
    `}`;

    req.send(data);
}