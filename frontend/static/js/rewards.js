document.addEventListener('DOMContentLoaded', function() {
    // Initialize filter buttons
    const filterButtons = document.querySelectorAll('.filter-btn');
    const rewardCards = document.querySelectorAll('.reward-card');

    // Filter functionality
    filterButtons.forEach(button => {
        button.addEventListener('click', () => {
            // Remove active class from all buttons
            filterButtons.forEach(btn => btn.classList.remove('active'));
            // Add active class to clicked button
            button.classList.add('active');

            const filter = button.dataset.filter;
            
            // Filter reward cards
            rewardCards.forEach(card => {
                if (filter === 'all' || card.dataset.category === filter) {
                    card.style.display = 'block';
                } else {
                    card.style.display = 'none';
                }
            });
        });
    });

    // Redeem button functionality
    const redeemButtons = document.querySelectorAll('.redeem-btn');
    redeemButtons.forEach(button => {
        button.addEventListener('click', function(e) {
            const card = this.closest('.reward-card');
            const pointsRequired = parseInt(card.querySelector('.points-required').textContent);
            const currentPoints = parseInt(document.querySelector('.points-value').textContent);

            if (currentPoints >= pointsRequired) {
                if (confirm('Are you sure you want to redeem this reward?')) {
                    // Simulate redemption
                    showNotification('success', 'Reward redeemed successfully!');
                    // Update points
                    document.querySelector('.points-value').textContent = 
                        (currentPoints - pointsRequired).toLocaleString();
                    // Add to redemption history
                    addToHistory(card.querySelector('h3').textContent, pointsRequired);
                }
            } else {
                showNotification('error', 'Not enough points to redeem this reward');
            }
        });
    });

    // Populate initial redemption history
    const mockHistoryData = [
        {
            date: '2024-03-15',
            reward: 'Shopping Voucher',
            points: 500,
            status: 'completed',
            details: 'Eco-Store Gift Card'
        },
        {
            date: '2024-03-10',
            reward: 'Premium Pickup Service',
            points: 1000,
            status: 'active',
            details: 'Valid until April 10, 2024'
        },
        {
            date: '2024-02-28',
            reward: 'Reusable Eco Bag Set',
            points: 200,
            status: 'shipped',
            details: 'Tracking: KE123456789'
        },
        {
            date: '2024-02-15',
            reward: 'Tree Planting Initiative',
            points: 300,
            status: 'completed',
            details: 'Tree planted in Karura Forest'
        },
        {
            date: '2024-02-01',
            reward: 'Solar Power Bank',
            points: 800,
            status: 'processing',
            details: 'Preparing for shipment'
        }
    ];

    // Populate history table with mock data
    function populateHistoryTable() {
        const tbody = document.getElementById('redemptionHistory');
        tbody.innerHTML = mockHistoryData.map(item => `
            <tr>
                <td>${formatDate(item.date)}</td>
                <td>
                    <div class="reward-info-cell">
                        <span class="reward-name">${item.reward}</span>
                        <span class="reward-details">${item.details}</span>
                    </div>
                </td>
                <td>
                    <span class="points-cell">
                        <i class="fas fa-star"></i>
                        ${item.points}
                    </span>
                </td>
                <td>
                    <span class="status-badge ${item.status}">
                        ${capitalizeFirst(item.status)}
                    </span>
                </td>
                <td>
                    <div class="action-buttons">
                        <button class="action-btn view" onclick="viewRewardDetails('${item.reward}')">
                            <i class="fas fa-eye"></i> View
                        </button>
                        ${item.status === 'shipped' ? `
                            <button class="action-btn track" onclick="trackShipment('${item.details.split(': ')[1]}')">
                                <i class="fas fa-truck"></i> Track
                            </button>
                        ` : ''}
                        ${item.status === 'completed' ? `
                            <button class="action-btn review" onclick="addReview('${item.reward}')">
                                <i class="fas fa-star"></i> Review
                            </button>
                        ` : ''}
                    </div>
                </td>
            </tr>
        `).join('');
    }

    // Helper functions
    function formatDate(dateString) {
        const options = { year: 'numeric', month: 'short', day: 'numeric' };
        return new Date(dateString).toLocaleDateString('en-US', options);
    }

    function capitalizeFirst(string) {
        return string.charAt(0).toUpperCase() + string.slice(1);
    }

    // Initialize history table
    populateHistoryTable();

    // Add to existing addToHistory function
    function addToHistory(rewardName, points) {
        const tbody = document.getElementById('redemptionHistory');
        const date = new Date().toLocaleDateString();
        
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${formatDate(date)}</td>
            <td>
                <div class="reward-info-cell">
                    <span class="reward-name">${rewardName}</span>
                    <span class="reward-details">Newly redeemed reward</span>
                </div>
            </td>
            <td>
                <span class="points-cell">
                    <i class="fas fa-star"></i>
                    ${points}
                </span>
            </td>
            <td>
                <span class="status-badge processing">
                    Processing
                </span>
            </td>
            <td>
                <div class="action-buttons">
                    <button class="action-btn view" onclick="viewRewardDetails('${rewardName}')">
                        <i class="fas fa-eye"></i> View
                    </button>
                </div>
            </td>
        `;
        
        tbody.insertBefore(tr, tbody.firstChild);
    }

    // Add these styles to your rewards.css
    const additionalStyles = `
        /* Redemption History Table Styles */
        .reward-info-cell {
            display: flex;
            flex-direction: column;
            gap: 0.25rem;
        }

        .reward-name {
            font-weight: 500;
            color: var(--text-primary);
        }

        .reward-details {
            font-size: 0.85rem;
            color: var(--text-secondary);
        }

        .points-cell {
            display: flex;
            align-items: center;
            gap: 0.5rem;
            color: var(--primary-color);
            font-weight: 500;
        }

        .points-cell i {
            color: #ffd700;
            font-size: 0.85rem;
        }

        .status-badge {
            padding: 0.35rem 0.75rem;
            border-radius: 20px;
            font-size: 0.85rem;
            font-weight: 500;
        }

        .status-badge.completed {
            background: rgba(46, 204, 113, 0.1);
            color: #27ae60;
        }

        .status-badge.processing {
            background: rgba(52, 152, 219, 0.1);
            color: #3498db;
        }

        .status-badge.active {
            background: rgba(155, 89, 182, 0.1);
            color: #9b59b6;
        }

        .status-badge.shipped {
            background: rgba(241, 196, 15, 0.1);
            color: #f1c40f;
        }

        .action-buttons {
            display: flex;
            gap: 0.5rem;
        }

        .action-btn {
            padding: 0.35rem 0.75rem;
            border: none;
            border-radius: 6px;
            cursor: pointer;
            font-size: 0.85rem;
            display: flex;
            align-items: center;
            gap: 0.35rem;
            transition: all 0.3s ease;
        }

        .action-btn.view {
            background: #f8f9fa;
            color: var(--text-primary);
        }

        .action-btn.track {
            background: rgba(52, 152, 219, 0.1);
            color: #3498db;
        }

        .action-btn.review {
            background: rgba(241, 196, 15, 0.1);
            color: #f1c40f;
        }

        .action-btn:hover {
            transform: translateY(-2px);
        }
    `;

    // Inject additional styles
    const styleSheet = document.createElement("style");
    styleSheet.textContent = additionalStyles;
    document.head.appendChild(styleSheet);

    // Notification system
    function showNotification(type, message) {
        const notification = document.createElement('div');
        notification.className = `notification ${type}`;
        notification.innerHTML = `
            <i class="fas ${type === 'success' ? 'fa-check-circle' : 'fa-exclamation-circle'}"></i>
            <p>${message}</p>
        `;
        
        document.body.appendChild(notification);
        
        setTimeout(() => {
            notification.remove();
        }, 3000);
    }

    // Mobile menu handling
    const mobileMenuToggle = document.querySelector('.mobile-menu-toggle');
    const sidebar = document.querySelector('.dashboard-sidebar');
    const overlay = document.querySelector('.sidebar-overlay');

    if (mobileMenuToggle) {
        mobileMenuToggle.addEventListener('click', function() {
            sidebar.classList.toggle('active');
            overlay.classList.toggle('active');
            document.body.classList.toggle('sidebar-open');
        });
    }

    if (overlay) {
        overlay.addEventListener('click', function() {
            sidebar.classList.remove('active');
            overlay.classList.remove('active');
            document.body.classList.remove('sidebar-open');
        });
    }
});

// Global functions for button actions
window.viewRewardDetails = function(rewardName) {
    alert(`Viewing details for ${rewardName}`);
    // Implement modal or page navigation for details
};

window.trackShipment = function(trackingNumber) {
    alert(`Tracking shipment: ${trackingNumber}`);
    // Implement tracking modal or redirect to tracking page
};

window.addReview = function(rewardName) {
    alert(`Add review for ${rewardName}`);
    // Implement review modal or form
}; 