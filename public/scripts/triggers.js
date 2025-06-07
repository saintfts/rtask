let task_input = document.getElementsByClassName('task-input')[0]


task_input.addEventListener('keydown', e => {
if (e.key == 'Enter') {
    task_input.children[0].value = task_input.children[0].value.trim()
    sendTask()
    task_input.children[0].value = ""
}
})


document.getElementById('task-list').addEventListener('click', (e) => {
    const delete_btn = e.target.closest('.delete-btn')
    if (delete_btn)
    {
        const task_div = delete_btn.closest('.task-item')
        const id = task_div.getAttribute('task-id')
        console.log(JSON.stringify({n:Number(id)}))
        fetch('/api/remove_task', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({n:Number(id)})
        })
        .then(_ => window.location.reload())
        console.log(1, Number(id))
    }
})
