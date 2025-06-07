
//var data_div = document.getElementById("data")
//console.log(data_div.dataset.tasks_len)

var task_id = tasks_arr.length
let tasks_div = document.getElementById("task-list")
function load_task(tasks) {
    console.log(tasks)
    tasks.forEach(task => {
        tasks_div.innerHTML += `
                <div task-id="` + ++task_id + `" class="task-item">
                    <div class="priority-indicator priority-high"></div>
                    <input type="checkbox" class="task-checkbox">
                    <div class="task-text">` + task.text + `</div>
                    <div class="task-date">Today</div>
                    <div class="task-actions">
                        <button class="task-btn edit-btn"><i class="fas fa-edit"></i></button>
                        <button type="submit" class="task-btn delete-btn"><i class="fas fa-trash"></i></button>
                    </div>
                </div>`
        tasks_arr.push(task)
    })
        //tasks_div.innerHTML += "<div = ><h1>" + task.id + ") "+ (task.text) + "</h1></div>";
}

async function get_new_tasks(){
    let body
    if (tasks_arr.length == 0) {
        body = {}
    } else {
        body = {id: tasks_arr.length}
    }
    fetch('/api/get_new_tasks',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
    })
    .then(response => response.json())
    .then(new_tasks => {
            //console.log(tasks_arr[tasks_arr.length - 1])
            if (new_tasks != undefined && Object.keys(new_tasks).length != 0) {
                console.log(new_tasks)
                load_task(new_tasks)
            }
        })
}

async function load_continuously() { // load new tasks and compare with most recent saved. If last id < last fetched_id, then update
    for (;;){
        get_new_tasks()
        await new Promise(r => setTimeout(r, 1500))
    }
}
load_continuously()
