function add_task(name) {
    $("#root ul").append("<li> <span>" + name + " </span> <button id = 'del_task'>Удалить</button> </li>");
    $("ul li:nth-last-child(1) button").click(function() {$(this).parent().remove()});
}

$(document).ready(function() {
    $("#root").append("<ul></ul>");
    add_task("Сделать задание #3 по web-программированию");
    $("#root").append("<input id = 'add_task_input'> <button id = 'add_task'>add_task</button>");
    $('#add_task').click(function(){add_task($("#add_task_input").val())});
});