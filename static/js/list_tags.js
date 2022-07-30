var host = window.location.origin


function get_list_tags(){
    let req = new XMLHttpRequest();
    req.open("GET", host+"/tags");

    req.setRequestHeader("Content-Type", "application/json");

    req.onload  = function() {
        if (req.status != 200){
            console.log("что-то пошло не так")
            return
        }
        outputNotes(JSON.parse(req.response))
    };

    req.send();
}

get_list_tags()

function outputNotes(tags){
    for(var i in tags) {
        var newDiv = document.createElement("div");
        newDiv.innerHTML = '<div style="border-style:solid;border-color:red;padding:5px;">'+
        '<p>'+tags[i].Tag+'</p>'+
        '<p>'+tags[i].Description+'</p>'+
        '<input id="'+tags[i].Tag+'" type="submit" value="Удалить тег" onclick="del_tag(this)">'+
        '</div>';
    
        var my_div = document.getElementById("list_tags");
        my_div.append(newDiv)

    }
}

function del_tag(elem){
    let req = new XMLHttpRequest();
    req.open("DELETE", host+"/tag/"+elem.id);

    req.setRequestHeader("Content-Type", "application/json");

    req.onload  = function() {
        if (req.status != 200){
            console.log("что-то пошло не так")
            return
        }

        document.getElementById("list_tags").innerHTML = "";
        get_list_tags()
    };

    req.send();
}

