var host = window.location.origin

function add_note(){
    var category = document.getElementById("form_add_note_category").value;
    var isFirst = document.getElementById("form_add_note_is_first").checked;
    var description = document.getElementById("form_add_note_description").value;
    var links = document.getElementById("form_add_note_links").value;
    var tags = document.getElementById("form_add_note_tags").value;
    var sources = document.getElementById("form_add_note_sources").value;
    var resp = document.getElementById('response');

    let req = new XMLHttpRequest();
    req.open("POST", host+"/note");

    req.setRequestHeader("Content-Type", "application/json");

    req.onload  = function() {
        if (req.status != 200){
            resp.innerHTML = "Попробуйте позже";
            return
        }
        clearForm();
    }

    links=links.replace(" ","");
    tags=tags.replace(" ","");

    if (links.length != 0){
        links = `["`+links.replace(",",'","')+`"]`;
    }else {
        links = "[]";
    }

    let data = `{`+
        `"Category":"`+category+`",`+
        `"Description":"`+description+`",`+
        `"Links":`+links+`,`+
        `"Tags":["`+tags.replace(",",'","')+`"],`+
        `"Sources":"`+sources+`",`+
        `"IsFirst":`+isFirst+``+
    `}`;

    req.send(data);
}

function clearForm(){
    document.getElementById("form_add_note_category").value="";
    document.getElementById("form_add_note_is_first").checked=false;
    document.getElementById("form_add_note_description").value="";
    document.getElementById("form_add_note_links").value="";
    document.getElementById("form_add_note_tags").value="";
    document.getElementById("form_add_note_sources").value="";
}



function get_list_categories(){
    let req = new XMLHttpRequest();
    req.open("GET", host+"/categories");

    req.setRequestHeader("Content-Type", "application/json");

    req.onload  = function() {
        if (req.status != 200){
            console.log("что-то пошло не так")
            return
        }
        createHTML(JSON.parse(req.response))
    };

    req.send();
}

get_list_categories()

function createHTML(notes){
    for(var i in notes) {
        var newDiv = document.createElement("option");
        newDiv.innerHTML = '<option>'+notes[i]+'</option>';
        var my_div = document.getElementById("form_add_note_category");
        my_div.append(newDiv);
    }
    document.getElementById("form_add_note_category").value="";
}
