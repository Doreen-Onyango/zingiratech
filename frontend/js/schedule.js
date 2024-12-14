document.addEventListener('DOMContentLoaded', function() {
    const scheduleForm = document.getElementById('scheduleForm');
    const pickupDateInput = document.getElementById('pickupDate');
    const wasteTypeSelect = document.getElementById('wasteType');
    
    // Set minimum date to tomorrow
    const tomorrow = new Date();
    tomorrow.setDate(tomorrow.getDate() + 1);
    pickupDateInput.min = tomorrow.toISOString().split('T')[0];

    // Set maximum date to 30 days from now
    const maxDate = new Date();
    maxDate.setDate(maxDate.getDate() + 30);
    pickupDateInput.max = maxDate.toISOString().split('T')[0];

    // Update pricing based on waste type
    wasteTypeSelect.addEventListener('change', function() {
        updatePricing(this.value);
    });

    // Form submission
    scheduleForm.addEventListener('submit', function(e) {
        e.preventDefault();
        
        // Collect form data
        const formData = {
            wasteType: wasteTypeSelect.value,
            quantity: document.getElementById('quantity').value,
            pickupDate: pickupDateInput.value,
            pickupTime: document.getElementById('pickupTime').value,
            address: document.getElementById('address').value,
            notes: document.getElementById('notes').value,
            manufacturer: {
                name: document.getElementById('manufacturerName').value,
                model: document.getElementById('modelNumber').value,
                year: document.getElementById('manufactureYear').value,
                condition: document.getElementById('condition').value
            }
        };

        // Validate form data
        if (validateFormData(formData)) {
            submitPickupRequest(formData);
        }
    });

    // Add price calculation
    const priceMap = {
        'electronics': 2,
        'batteries': 3,
        'appliances': 1.5,
        'computers': 2.5,
        'phones': 4,
        'other': 2
    };

    function updatePricing(wasteType) {
        const quantity = document.getElementById('quantity').value;
        const price = priceMap[wasteType] || 0;
        const total = price * quantity;
        
        const estimatedPrice = document.createElement('div');
        estimatedPrice.className = 'estimated-price';
        estimatedPrice.innerHTML = `
            <h4>Estimated Price</h4>
            <p>$${total.toFixed(2)}</p>
        `;

        const existingEstimate = document.querySelector('.estimated-price');
        if (existingEstimate) {
            existingEstimate.remove();
        }

        document.querySelector('.pricing-card').appendChild(estimatedPrice);
    }

    // Update price when quantity changes
    document.getElementById('quantity').addEventListener('input', function() {
        const wasteType = wasteTypeSelect.value;
        if (wasteType) {
            updatePricing(wasteType);
        }
    });

    // Enhanced form validation
    function validateFormData(data) {
        const errors = [];

        if (!data.wasteType) {
            errors.push('Please select a waste type');
        }

        if (!data.quantity || data.quantity <= 0) {
            errors.push('Please enter a valid quantity');
        }

        const selectedDate = new Date(data.pickupDate);
        const today = new Date();
        if (selectedDate <= today) {
            errors.push('Please select a future date');
        }

        if (!data.pickupTime) {
            errors.push('Please select a pickup time');
        }

        if (!data.address || data.address.trim().length < 10) {
            errors.push('Please enter a complete address (minimum 10 characters)');
        }

        // Manufacturer validations
        if (data.manufacturer.name && !data.manufacturer.model) {
            errors.push('Please provide the model number if manufacturer is specified');
        }

        if (data.manufacturer.model && !data.manufacturer.name) {
            errors.push('Please provide the manufacturer name');
        }

        if (data.manufacturer.year) {
            const year = parseInt(data.manufacturer.year);
            const currentYear = new Date().getFullYear();
            if (year < 1970 || year > currentYear) {
                errors.push('Please enter a valid manufacture year');
            }
        }

        if (errors.length > 0) {
            showError(errors.join('\n'));
            return false;
        }

        return true;
    }

    // Enhanced notifications
    function showError(message) {
        const notification = createNotification('error', message);
        document.body.appendChild(notification);
        setTimeout(() => notification.remove(), 5000);
    }

    function showSuccess(message) {
        const notification = createNotification('success', message);
        document.body.appendChild(notification);
        setTimeout(() => notification.remove(), 5000);
    }

    function createNotification(type, message) {
        const notification = document.createElement('div');
        notification.className = `notification ${type}`;
        notification.innerHTML = `
            <i class="fas ${type === 'error' ? 'fa-exclamation-circle' : 'fa-check-circle'}"></i>
            <p>${message}</p>
        `;
        return notification;
    }

    // Add loading animation to form submission
    function submitPickupRequest(data) {
        const submitBtn = scheduleForm.querySelector('.submit-btn');
        submitBtn.classList.add('loading');
        submitBtn.disabled = true;

        // Simulate API call
        setTimeout(() => {
            console.log('Pickup scheduled:', data);
            submitBtn.classList.remove('loading');
            showSuccess('Pickup scheduled successfully!');
            
            // Store in localStorage for history
            const pickups = JSON.parse(localStorage.getItem('pickups') || '[]');
            pickups.push({
                ...data,
                id: Date.now(),
                status: 'scheduled'
            });
            localStorage.setItem('pickups', JSON.stringify(pickups));

            // Redirect after success
            setTimeout(() => {
                window.location.href = 'dashboard-user.html#pickups';
            }, 2000);
        }, 1500);
    }

    // Add touch feedback for mobile
    const buttons = document.querySelectorAll('button');
    buttons.forEach(button => {
        button.addEventListener('touchstart', function() {
            this.style.opacity = '0.8';
        });
        button.addEventListener('touchend', function() {
            this.style.opacity = '1';
        });
    });

    // Mobile menu handling
    const mobileMenuToggle = document.querySelector('.mobile-menu-toggle');
    const sidebar = document.querySelector('.dashboard-sidebar');
    const overlay = document.querySelector('.sidebar-overlay');

    function toggleSidebar() {
        sidebar.classList.toggle('active');
        overlay.classList.toggle('active');
        document.body.classList.toggle('sidebar-open');
    }

    // Toggle sidebar when clicking menu button
    if (mobileMenuToggle) {
        mobileMenuToggle.addEventListener('click', function(e) {
            e.stopPropagation();
            toggleSidebar();
        });
    }

    // Close sidebar when clicking overlay
    if (overlay) {
        overlay.addEventListener('click', toggleSidebar);
    }

    // Close sidebar when clicking outside
    document.addEventListener('click', function(e) {
        if (window.innerWidth <= 768) {
            if (!sidebar.contains(e.target) && 
                !mobileMenuToggle.contains(e.target) && 
                sidebar.classList.contains('active')) {
                toggleSidebar();
            }
        }
    });

    // Close sidebar when clicking nav items on mobile
    const navItems = document.querySelectorAll('.nav-item');
    navItems.forEach(item => {
        item.addEventListener('click', function() {
            if (window.innerWidth <= 768 && sidebar.classList.contains('active')) {
                toggleSidebar();
            }
        });
    });

    // Handle window resize
    window.addEventListener('resize', function() {
        if (window.innerWidth > 768) {
            sidebar.classList.remove('active');
            overlay.classList.remove('active');
            document.body.classList.remove('sidebar-open');
        }
    });
}); 