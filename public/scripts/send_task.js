function sendTask(){
    const message = document.querySelector('.task-input').querySelector('input').value
    if (message.length == 0) {
        console.log("Tried to send an empty string")
        return
    }
    fetch("api/send_task", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({text: message})
    })
}
