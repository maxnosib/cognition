function CheckAuth(){
    var isLogin = localStorage.getItem("isLogin");
    var curent_time = new Date().getTime()
    if (isLogin < curent_time){
        window.location = '/auth.html';
    }
}

CheckAuth()







