let newTaskOpen = false;
let taskDateChanged = false;
let allTasks = new Map()
let allTaskRows = new Map()
let activeTaskId = -1;
let lastActiveTaskId = -1;

function addTaskRow(taskId) {
    const sidebar = document.querySelector("#sidebar")
    let task = allTasks.get(taskId)

    const input = document.createElement('input')
    input.type = 'checkbox'
    input.name = task.title
    input.style.margin = 2
    input.style.alignSelf = 'right'
    input.style.fontStyle = 'bold'

    const title = document.createElement('label')
    title.textContent = task.title
    title.id = 'title'

    const date = document.createElement('label')
    date.textContent = task.dueDate
    date.style.marginLeft = 'auto'
    date.style.marginRight = '2px'
    date.style.fontStyle = 'italic'
    date.id = 'date'

    const row = document.createElement('div')
    row.className = 'task-row'
    row.id = task.taskId
    row.appendChild(input)
    row.appendChild(title)
    row.appendChild(date)

    row.querySelector('#task')

    row.addEventListener('click', function () {
        lastActiveTaskId = activeTaskId
        activeTaskId = parseInt(row.id)
        highlightActiveTaskRow()
        disableSaveAndCancel()
    })

    sidebar.appendChild(row)
    return row
}

function highlightActiveTaskRow() {
    if (lastActiveTaskId !== -1) {
        let lastActiveRow = allTaskRows.get(lastActiveTaskId)
        lastActiveRow.style.backgroundColor = 'var(--main-color)'
    }
    if (activeTaskId !== -1) {
        let activeRow = allTaskRows.get(activeTaskId)
        activeRow.style.backgroundColor = 'var(--highlight-color)'
    }
    showTaskDetails()
}

function showTaskDetails() {
    let activeTask = allTasks.get(activeTaskId)
    const placeholder = document.querySelector("#no-content");
    placeholder.style.display = 'none'

    const content = document.querySelector("#content");
    content.style.display = 'flex'
    content.style.flexDirection = 'column'

    const title = document.querySelector("#task-details-title");
    title.value = activeTask.title

    const desc = document.querySelector("#task-details-desc");
    desc.value = activeTask.description

    const date = document.querySelector("#task-details-date");
    date.value = activeTask.dueDate
}

function enableSaveAndCancel() {
    const title = document.querySelector("#task-details-title");
    const desc = document.querySelector("#task-details-desc");
    const date = document.querySelector("#task-details-date");

    const saveButton = document.querySelector("#task-details-save");
    saveButton.style.opacity = '100%';
    saveButton.disabled = false;

    let activeTask = allTasks.get(activeTaskId)

    saveButton.addEventListener('click', function () {
        activeTask.title = title.value
        activeTask.description = desc.value
        activeTask.dueDate = date.value
        allTasks.set(activeTaskId, activeTask)
        updateTaskRow(title.value, date.value)
        disableSaveAndCancel()
    })

    const cancelButton = document.querySelector("#task-details-cancel");
    cancelButton.style.opacity = '100%';
    cancelButton.disabled = false;

    cancelButton.addEventListener('click', function () {
        title.value = activeTask.title
        desc.value = activeTask.description
        date.value = activeTask.dueDate
        disableSaveAndCancel()
    })
}

function updateTaskRow(title, date) {
    let activeTaskRow = allTaskRows.get(activeTaskId)
    activeTaskRow.querySelector('#title').textContent = title
    activeTaskRow.querySelector('#date').textContent = date
    allTaskRows.set(activeTaskId, activeTaskRow)
}

function disableSaveAndCancel() {
    const saveButton = document.querySelector("#task-details-save");
    saveButton.style.opacity = '50%';
    saveButton.disabled = true;

    const cancelButton = document.querySelector("#task-details-cancel");
    cancelButton.style.opacity = '50%';
    cancelButton.disabled = true;
}

function fetchCurrentTasks() {
    const response = fetch("/list")
        .then(response => response.json())
        .then(data => {
            console.log(data)
            if (data != null) {
                for (let i = 0; i < data.length; i++) {
                    if (activeTaskId === -1) {
                        activeTaskId = data[i].taskId;
                    }
                    allTasks.set(data[i].taskId, data[i])
                    allTaskRows.set(data[i].taskId, addTaskRow(data[i].taskId))
                }
                highlightActiveTaskRow()
            }
        })
}

function createNewTask(title, desc, dueDate) {
    const params = new URLSearchParams();
    params.append("title", title);
    params.append("description", desc);
    params.append("dueDate", dueDate);

    let url = "/create?";
    url += params.toString();

    const response = fetch(url)
        .then(response => response.json())
        .then(data => {
            console.log(data)
            allTasks.set(data.taskId, data)
            allTaskRows.set(data.taskId, addTaskRow(data))
        })
}

function addNewTaskModal() {
    const title = document.querySelector('#new-task-title')
    const details = document.querySelector('#new-task-details')
    const date = document.querySelector('#new-task-date')
    const desc = document.querySelector('#new-task-desc')
    const save = document.querySelector('#new-task-save')
    const cancel = document.querySelector('#new-task-cancel')

    title.addEventListener('click', function () {
        showNewTaskModal(details, date);
        newTaskOpen = true;
    })

    date.addEventListener('click', function () {
        taskDateChanged = true;
    })

    save.addEventListener('click', function () {
        console.log("Creating new task with desc: " + desc.value)
        createNewTask(title.value, desc.value, date.value)
        clearNewTaskModal(title, date, desc);
    })

    cancel.addEventListener('click', function () {
        clearNewTaskModal(title, date, desc);
    })

    window.onclick = function(event) {
        if (event.target.id !== 'new-task-details'
            && event.target.id !== 'new-task-title'
            && event.target.id !== 'new-task-desc'
            && event.target.id !== 'new-task-footer'
            && event.target.id !== 'new-task-date'
            && newTaskOpen) {
            hideNewTaskModal();
        }
    }

}

function showNewTaskModal(details, date) {
    details.style.display = 'flex'
    details.style.flexDirection = 'column'
    if (!taskDateChanged) {
        date.valueAsDate = new Date()
    }
}

function hideNewTaskModal() {
    const details = document.querySelector('#new-task-details')
    details.style.display = 'none'
}

function clearNewTaskModal(title, date, desc) {
    title.value = "";
    date.valueAsDate = new Date();
    taskDateChanged = false;
    desc.value = "";
}

function addListener() {
    console.log("Started adding listeners")
    fetchCurrentTasks()
    addNewTaskModal()
}

document.addEventListener("DOMContentLoaded", addListener);

