document.addEventListener('DOMContentLoaded', function() {
    // Initialize Charts
    initializeCharts();

    // Populate table with mock data
    populateHistoryTable();

    // Handle time filter changes
    document.getElementById('timeFilter').addEventListener('change', function() {
        // Update charts and table based on selected time period
        updateData(this.value);
    });

    // Handle export button
    document.querySelector('.export-btn').addEventListener('click', exportHistory);
});

function initializeCharts() {
    // Recycling Trends Chart
    const trendsCtx = document.getElementById('recyclingTrends').getContext('2d');
    new Chart(trendsCtx, {
        type: 'line',
        data: {
            labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun'],
            datasets: [{
                label: 'Weight (kg)',
                data: [30, 45, 35, 50, 40, 45],
                borderColor: '#2ecc71',
                tension: 0.4,
                fill: true,
                backgroundColor: 'rgba(46, 204, 113, 0.1)'
            }]
        },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            plugins: {
                legend: {
                    display: false
                }
            },
            scales: {
                y: {
                    beginAtZero: true
                }
            }
        }
    });

    // Waste Composition Chart
    const compositionCtx = document.getElementById('wasteComposition').getContext('2d');
    new Chart(compositionCtx, {
        type: 'doughnut',
        data: {
            labels: ['Electronics', 'Batteries', 'Appliances', 'Computers', 'Phones'],
            datasets: [{
                data: [30, 20, 25, 15, 10],
                backgroundColor: [
                    '#2ecc71',
                    '#3498db',
                    '#e74c3c',
                    '#f1c40f',
                    '#9b59b6'
                ]
            }]
        },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            plugins: {
                legend: {
                    position: 'right'
                }
            }
        }
    });
}

function populateHistoryTable() {
    const tbody = document.getElementById('historyTableBody');
    const mockData = [
        {
            date: '2024-03-15',
            type: 'Electronics',
            quantity: '15 kg',
            status: 'completed',
            impact: '25 kg CO₂'
        },
        {
            date: '2024-03-10',
            type: 'Batteries',
            quantity: '5 kg',
            status: 'pending',
            impact: '8 kg CO₂'
        },
        // Add more mock data as needed
    ];

    tbody.innerHTML = mockData.map(item => `
        <tr>
            <td>${formatDate(item.date)}</td>
            <td>${item.type}</td>
            <td>${item.quantity}</td>
            <td>
                <span class="status-badge ${item.status}">
                    ${item.status.charAt(0).toUpperCase() + item.status.slice(1)}
                </span>
            </td>
            <td>${item.impact}</td>
            <td>
                <button class="action-btn view">
                    <i class="fas fa-eye"></i> View
                </button>
            </td>
        </tr>
    `).join('');
}

function formatDate(dateString) {
    const options = { year: 'numeric', month: 'short', day: 'numeric' };
    return new Date(dateString).toLocaleDateString('en-US', options);
}

function updateData(timeFilter) {
    // Update charts and table based on selected time period
    console.log('Updating data for:', timeFilter);
    // Implement actual data filtering logic here
}

function exportHistory() {
    // Implement export functionality
    console.log('Exporting history...');
    // You could generate CSV/PDF here
    alert('Export feature coming soon!');
} 