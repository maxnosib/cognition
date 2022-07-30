var host = window.location.origin


function get_note(){
    let req = new XMLHttpRequest();
    req.open("GET", host+"/note/"+localStorage.getItem("note"));

    req.setRequestHeader("Content-Type", "application/json");

    req.onload  = function() {
        if (req.status != 200){
            console.log("что-то пошло не так")
            return
        }
        outputNote(JSON.parse(req.response))
    };

    req.send();
}

get_note()


function outputNote(notes){
        // просмотр записки
        var newDiv = document.createElement("div");
        newDiv.innerHTML = '<div style="border-style:solid;border-color:red;padding:5px;">'+
        '<a href="#" onclick="goToOneNote(this);return false;">'+notes.ID+'</a>'+
        '<p>'+notes.Category+'</p>'+
        '<p>'+notes.Description+'</p>'+
        '<p>'+notes.Links+'</p>'+
        '<p>'+notes.Tags+'</p>'+
        '<p>'+notes.Sources+'</p>'+
        '<p>'+notes.IsFirst+'</p>'+
        '<p>'+notes.CreatedAt+'</p>'+
        '<p>'+notes.LastUpdatedAt+'</p>'+
        '<p><input type="submit" value="Редактировать" onclick="redact(this)"></p>'+
        '</div>';
        var my_div = document.getElementById("view_note");
        my_div.append(newDiv)

        var isFirst = "";
        if(notes.IsFirst == true){
            isFirst="checked"
        }
        // редактирование записки
        var newDiv = document.createElement("div");
        newDiv.innerHTML = '<div id="form_add_note">'+
        '<p>ID записки <span id="form_add_note_id">'+notes.ID+'</span></p>'+
        '<p>Выберите 1 категорию</p>'+
        '<select id="form_add_note_category" required>'+
            '<option>'+notes.Category+'</option>'+
            '<option>знания</option>'+
            '<option>мысли</option>'+
            '<option>архив</option>'+
            '<option>заметки</option>'+
        '</select>'+
        '<p>Является ли эта карточка первой в векторе</p>'+
        '<input type="checkbox" id="form_add_note_is_first" name="Первая карточка в векторе" '+isFirst+'>'+
        '<p>Введите описание</p>'+
        '<textarea id="form_add_note_description" placeholder="Описание" type="text" required>'+notes.Description+'</textarea>'+
        '<p>Введите ссылки (id) на другие заметки через запятую</p>'+
        '<input id="form_add_note_links" placeholder="ссылка1,ссылка2" type="text" value="'+notes.Links+'">'+
        '<p>Введите теги через запятую</p>'+
        '<input id="form_add_note_tags" placeholder="тег1, тег2" type="text" required value="'+notes.Tags+'">'+
        '<p>Введите ссылки на ресурсы</p>'+
        '<textarea id="form_add_note_sources" placeholder="Ресурсы" type="text">'+notes.Sources+'</textarea>'+
        '<p>Время создания '+notes.CreatedAt+'</p>'+
        '<p>Время последнего обновления '+notes.LastUpdatedAt+'</p>'+
        '<p><input type="submit" value="Обновить записку" onclick="updateNote()"></p>'+
        '</div>';
        var my_div = document.getElementById("update_note");
        my_div.append(newDiv)
}

function updateNote(){
    var id = document.getElementById("form_add_note_id").innerHTML;
    var category = document.getElementById("form_add_note_category").value;
    var isFirst = document.getElementById("form_add_note_is_first").checked;
    var description = document.getElementById("form_add_note_description").value;
    var links = document.getElementById("form_add_note_links").value;
    var tags = document.getElementById("form_add_note_tags").value;
    var sources = document.getElementById("form_add_note_sources").value;
    var resp = document.getElementById('response');


    let req = new XMLHttpRequest();
    req.open("PATCH", host+"/note");

    req.setRequestHeader("Content-Type", "application/json");

    req.onload  = function() {
        if (req.status != 200){
            resp.innerHTML = "Попробуйте позже";
            return
        }
        document.getElementById("view_note").hidden = false;
        document.getElementById("view_note").innerHTML = "";
        document.getElementById("update_note").hidden = true;
        document.getElementById("update_note").innerHTML = "";
        get_note()
    };

    let data = `{`+
    `"ID":"`+id+`",`+
    `"Category":"`+category+`",`+
    `"Description":"`+description+`",`+
    `"Links":["`+links.replace(",",'","')+`"],`+
    `"Tags":["`+tags.replace(",",'","')+`"],`+
    `"Sources":"`+sources+`",`+
    `"IsFirst":`+isFirst+``+
    `}`;

    req.send(data);
}

function redact(){
    document.getElementById("view_note").hidden = true;
    document.getElementById("update_note").hidden = false;
}

