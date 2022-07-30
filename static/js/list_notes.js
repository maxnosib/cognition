var host = window.location.origin


function get_list_notes(){
    let req = new XMLHttpRequest();
    req.open("GET", host+"/notes");

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

get_list_notes()


function outputNotes(notes){
    for(var i in notes) {
        var newDiv = document.createElement("div");
        newDiv.innerHTML = '<div style="border-style:solid;border-color:red;padding:5px;">'+
        '<a href="#" onclick="goToOneNote(this);return false;">'+notes[i].ID+'</a>'+
        '<p>'+notes[i].Category+'</p>'+
        '<p>'+notes[i].Description+'</p>'+
        '<p>'+notes[i].Links+'</p>'+
        '<p>'+notes[i].Tags+'</p>'+
        '<p>'+notes[i].Sources+'</p>'+
        '<p>'+notes[i].IsFirst+'</p>'+
        '<p>'+notes[i].CreatedAt+'</p>'+
        '<p>'+notes[i].LastUpdatedAt+'</p>'+
        '</div>';
    
        var my_div = document.getElementById("list_note");
        my_div.append(newDiv)

    }
}

function goToOneNote(elem){
    localStorage.setItem("note", elem.innerHTML);
    window.location = '/one_notes.html';
}
