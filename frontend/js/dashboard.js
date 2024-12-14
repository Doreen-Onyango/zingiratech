document.addEventListener('DOMContentLoaded', function() {
    // Mobile menu toggle
    const menuToggle = document.querySelector('.menu-toggle');
    const sidebar = document.querySelector('.dashboard-sidebar');
    
    if (menuToggle) {
        menuToggle.addEventListener('click', () => {
            sidebar.classList.toggle('active');
        });
    }

    // Close sidebar when clicking outside on mobile
    document.addEventListener('click', (e) => {
        if (window.innerWidth <= 768) {
            if (!sidebar.contains(e.target) && !menuToggle.contains(e.target)) {
                sidebar.classList.remove('active');
            }
        }
    });

    // User menu dropdown
    const userMenu = document.querySelector('.user-menu');
    if (userMenu) {
        userMenu.addEventListener('click', (e) => {
            userMenu.classList.toggle('active');
        });
    }

    // Quick action button - New Pickup
    const quickActionBtn = document.querySelector('.quick-action-btn');
    if (quickActionBtn) {
        quickActionBtn.addEventListener('click', () => {
            // Navigate to schedule pickup page
            window.location.href = '#schedule';
        });
    }

    // View buttons functionality
    const viewButtons = document.querySelectorAll('.btn-outline');
    viewButtons.forEach(button => {
        button.addEventListener('click', function(e) {
            const pickupId = this.closest('.pickup-item')
                               .querySelector('.pickup-header h4')
                               .textContent.split('#')[1];
            // Navigate to pickup details
            window.location.href = `#pickups/${pickupId}`;
        });
    });

    // Cancel buttons functionality
    const cancelButtons = document.querySelectorAll('.btn-danger');
    cancelButtons.forEach(button => {
        button.addEventListener('click', function(e) {
            if (confirm('Are you sure you want to cancel this pickup?')) {
                // Handle pickup cancellation
                const pickupItem = this.closest('.pickup-item');
                pickupItem.style.opacity = '0.5';
                pickupItem.style.pointerEvents = 'none';
                // Here you would typically make an API call to cancel the pickup
            }
        });
    });

    // Select recycler functionality
    const selectButtons = document.querySelectorAll('.recycler-item .btn-primary');
    selectButtons.forEach(button => {
        button.addEventListener('click', function(e) {
            const recyclerName = this.closest('.recycler-item')
                                   .querySelector('h4')
                                   .textContent;
            // Navigate to schedule pickup with pre-selected recycler
            window.location.href = `#schedule?recycler=${encodeURIComponent(recyclerName)}`;
        });
    });

    // Stats counter animation
    function animateValue(element, start, end, duration) {
        let startTimestamp = null;
        const step = (timestamp) => {
            if (!startTimestamp) startTimestamp = timestamp;
            const progress = Math.min((timestamp - startTimestamp) / duration, 1);
            const currentValue = Math.floor(progress * (end - start) + start);
            element.textContent = currentValue;
            if (progress < 1) {
                window.requestAnimationFrame(step);
            }
        };
        window.requestAnimationFrame(step);
    }

    // Animate stats when they come into view
    const statsObserver = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                const statNumbers = entry.target.querySelectorAll('.stat-number');
                statNumbers.forEach(stat => {
                    const finalValue = parseInt(stat.textContent);
                    animateValue(stat, 0, finalValue, 2000);
                });
                statsObserver.unobserve(entry.target);
            }
        });
    }, { threshold: 0.5 });

    const statsGrid = document.querySelector('.stats-grid');
    if (statsGrid) {
        statsObserver.observe(statsGrid);
    }

    // Handle window resize
    let resizeTimer;
    window.addEventListener('resize', () => {
        clearTimeout(resizeTimer);
        resizeTimer = setTimeout(() => {
            if (window.innerWidth > 768) {
                sidebar.classList.remove('active');
            }
        }, 250);
    });

    // Error handling function
    function handleError(error, message = 'An error occurred') {
        console.error(error);
        // You can implement a toast notification system here
        alert(message);
    }

    // Add active class to current nav item based on URL
    function updateActiveNavItem() {
        const currentPath = window.location.hash || '#overview';
        const navItems = document.querySelectorAll('.sidebar-nav a');
        navItems.forEach(item => {
            if (item.getAttribute('href') === currentPath) {
                item.parentElement.classList.add('active');
            } else {
                item.parentElement.classList.remove('active');
            }
        });
    }

    // Listen for hash changes
    window.addEventListener('hashchange', updateActiveNavItem);
    updateActiveNavItem();

    // Filter functionality
    const filterButtons = document.querySelectorAll('.filter-tabs button');
    const pickupCards = document.querySelectorAll('.pickup-card');

    function updateFilterCounts() {
        const counts = {
            all: pickupCards.length,
            scheduled: 0,
            inProgress: 0,
            completed: 0
        };

        pickupCards.forEach(card => {
            const progress = parseInt(card.querySelector('.status').textContent);
            if (progress === 100) {
                counts.completed++;
            } else if (progress === 0) {
                counts.scheduled++;
            } else {
                counts.inProgress++;
            }
        });

        // Update button counts
        filterButtons.forEach(button => {
            const type = button.textContent.split(' ')[0].toLowerCase();
            const count = counts[type === 'all' ? 'all' : 
                         type === 'scheduled' ? 'scheduled' : 
                         type === 'in' ? 'inProgress' : 'completed'];
            const span = button.querySelector('span');
            if (span) {
                span.textContent = count;
            }
        });
    }

    function filterPickups(status) {
        pickupCards.forEach(card => {
            const progress = parseInt(card.querySelector('.status').textContent);
            
            if (status === 'all') {
                card.style.display = 'block';
            } else if (status === 'scheduled' && progress === 0) {
                card.style.display = 'block';
            } else if (status === 'in progress' && progress > 0 && progress < 100) {
                card.style.display = 'block';
            } else if (status === 'completed' && progress === 100) {
                card.style.display = 'block';
            } else {
                card.style.display = 'none';
            }
        });
    }

    filterButtons.forEach(button => {
        button.addEventListener('click', () => {
            // Remove active class from all buttons
            filterButtons.forEach(btn => btn.classList.remove('active'));
            // Add active class to clicked button
            button.classList.add('active');
            
            // Get filter status from button text
            const status = button.textContent.split(' ')[0].toLowerCase();
            filterPickups(status);
        });
    });

    // Initialize counts
    updateFilterCounts();

    // Mobile menu toggle functionality
    const mobileMenuToggle = document.querySelector('.mobile-menu-toggle');
    const main = document.querySelector('.dashboard-main');

    if (mobileMenuToggle) {
        mobileMenuToggle.addEventListener('click', () => {
            sidebar.classList.toggle('active');
            document.body.classList.toggle('sidebar-open');
        });

        // Close sidebar when clicking outside
        document.addEventListener('click', (e) => {
            if (window.innerWidth <= 768) {
                if (!sidebar.contains(e.target) && 
                    !mobileMenuToggle.contains(e.target) && 
                    sidebar.classList.contains('active')) {
                    sidebar.classList.remove('active');
                    document.body.classList.remove('sidebar-open');
                }
            }
        });

        // Handle window resize
        window.addEventListener('resize', () => {
            if (window.innerWidth > 768) {
                sidebar.classList.remove('active');
                document.body.classList.remove('sidebar-open');
            }
        });
    }

    // Add touch swipe functionality for mobile
    let touchStartX = 0;
    let touchEndX = 0;

    document.addEventListener('touchstart', e => {
        touchStartX = e.changedTouches[0].screenX;
    }, false);

    document.addEventListener('touchend', e => {
        touchEndX = e.changedTouches[0].screenX;
        handleSwipe();
    }, false);

    function handleSwipe() {
        const swipeThreshold = 50;
        const swipeDistance = touchEndX - touchStartX;

        if (Math.abs(swipeDistance) > swipeThreshold) {
            if (swipeDistance > 0 && touchStartX < 50) {
                // Swipe right from left edge
                sidebar.classList.add('active');
                document.body.classList.add('sidebar-open');
            } else if (swipeDistance < 0 && sidebar.classList.contains('active')) {
                // Swipe left
                sidebar.classList.remove('active');
                document.body.classList.remove('sidebar-open');
            }
        }
    }

    // Handle scroll behavior
    let lastScrollTop = 0;
    window.addEventListener('scroll', () => {
        const st = window.pageYOffset || document.documentElement.scrollTop;
        if (st > lastScrollTop && st > 100) {
            // Scrolling down
            document.querySelector('.mobile-menu-toggle').style.transform = 'translateY(-100%)';
        } else {
            // Scrolling up
            document.querySelector('.mobile-menu-toggle').style.transform = 'translateY(0)';
        }
        lastScrollTop = st <= 0 ? 0 : st;
    }, false);

    document.querySelector('.schedule-btn').addEventListener('click', function() {
        window.location.href = 'schedule-pickup.html';
    });
}); 