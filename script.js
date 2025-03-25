// Render Gantt Chart using Chart.js or any other Gantt chart library

const tasks = {{ tasks | tojson }}; // Data passed from Flask to the frontend

function renderGanttChart(tasks) {
    const ctx = document.getElementById('ganttChart').getContext('2d');
    
    const data = {
        labels: tasks.map(task => task[1]), // Task names
        datasets: [{
            label: 'Task Progress',
            data: tasks.map(task => task[4]), // Task progress
            backgroundColor: '#76c7c0',
            borderColor: '#4d9a9a',
            borderWidth: 1
        }]
    };

    const config = {
        type: 'bar',
        data: data,
        options: {
            indexAxis: 'y', // Horizontal Bar
            responsive: true,
            scales: {
                x: {
                    beginAtZero: true
                }
            }
        }
    };

    const chart = new Chart(ctx, config);
}

// Initial render of Gantt chart
renderGanttChart(tasks);

function addTask() {
    // Example of adding a task via API (AJAX)
    const taskData = {
        name: 'New Task',
        start_date: '2025-02-01',
        end_date: '2025-02-07',
        progress: 50
    };
    
    fetch('/add_task', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(taskData)
    })
    .then(response => response.json())
    .then(data => {
        if (data.status === 'success') {
            // Refresh Gantt chart with the new task
            window.location.reload();
        }
    });
}
