var host = window.location.origin

function add_tag(){
    var tag = document.getElementById("form_add_tag_tag").value;
    var description = document.getElementById("form_add_tag_description").value;
    var resp = document.getElementById('response');

    let req = new XMLHttpRequest();
    req.open("POST", host+"/tag");

    req.setRequestHeader("Content-Type", "application/json");

    req.onload  = function() {
        if (req.status != 200){
            resp.innerHTML = "Попробуйте позже";
            return
        }
        clearForm()
    };

    let data = `{`+
        `"Tag":"`+tag+`",`+
        `"Description":"`+description+`"`+
    `}`;

    req.send(data);
}

function clearForm(){
    document.getElementById("form_add_tag_tag").value="";
    document.getElementById("form_add_tag_description").value="";
}
