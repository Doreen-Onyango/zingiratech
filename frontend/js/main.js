// Sticky Header
window.addEventListener('scroll', () => {
    const header = document.querySelector('.header');
    header.classList.toggle('sticky', window.scrollY > 0);
});

// Scroll to Top
const scrollBtn = document.getElementById('scroll-to-top');
window.addEventListener('scroll', () => {
    if (window.scrollY > 300) {
        scrollBtn.classList.add('show');
    } else {
        scrollBtn.classList.remove('show');
    }
});

scrollBtn.addEventListener('click', () => {
    window.scrollTo({
        top: 0,
        behavior: 'smooth'
    });
});

// Mobile Menu Toggle
const mobileMenu = document.querySelector('.mobile-menu');
const navLinks = document.querySelector('.nav-links');
const navRight = document.querySelector('.nav-right');
const body = document.body;

// Clone nav-right elements for mobile menu
const navRightClone = navRight.cloneNode(true);
navLinks.appendChild(navRightClone);

mobileMenu.addEventListener('click', () => {
    mobileMenu.classList.toggle('active');
    navLinks.classList.toggle('active');
    body.style.overflow = navLinks.classList.contains('active') ? 'hidden' : '';
});

// Close mobile menu when clicking outside
document.addEventListener('click', (e) => {
    if (!navLinks.contains(e.target) && !mobileMenu.contains(e.target)) {
        mobileMenu.classList.remove('active');
        navLinks.classList.remove('active');
        body.style.overflow = '';
    }
});

// Handle dropdown menus on mobile
const dropdowns = document.querySelectorAll('.dropdown');
dropdowns.forEach(dropdown => {
    const link = dropdown.querySelector('a');
    const menu = dropdown.querySelector('.dropdown-menu');
    
    link.addEventListener('click', (e) => {
        if (window.innerWidth <= 768) {
            e.preventDefault();
            menu.style.display = menu.style.display === 'block' ? 'none' : 'block';
        }
    });
});

// Resize handling
window.addEventListener('resize', () => {
    if (window.innerWidth > 768) {
        mobileMenu.classList.remove('active');
        navLinks.classList.remove('active');
        body.style.overflow = '';
        
        // Reset dropdown menus
        document.querySelectorAll('.dropdown-menu').forEach(menu => {
            menu.style.display = '';
        });
    }
});

// Enhanced Scroll Animations
const scrollElements = document.querySelectorAll('.scroll-animation');

const elementInView = (el, percentageScroll = 100) => {
    const elementTop = el.getBoundingClientRect().top;
    return (
        elementTop <= 
        ((window.innerHeight || document.documentElement.clientHeight) * (percentageScroll/100))
    );
};

const displayScrollElement = (element) => {
    element.classList.add('scrolled');
};

const hideScrollElement = (element) => {
    element.classList.remove('scrolled');
};

const handleScrollAnimation = () => {
    scrollElements.forEach((el) => {
        if (elementInView(el, 100)) {
            displayScrollElement(el);
        } else {
            hideScrollElement(el);
        }
    })
}

// Stats Counter Animation
const stats = document.querySelectorAll('.stat-number');
const statsSection = document.querySelector('.stats');
let animated = false;

const animateStats = () => {
    stats.forEach(stat => {
        const target = parseInt(stat.getAttribute('data-target'));
        let current = 0;
        const increment = target / 100;
        
        const updateCount = () => {
            if(current < target) {
                current += increment;
                stat.textContent = Math.round(current);
                setTimeout(updateCount, 10);
            } else {
                stat.textContent = target;
            }
        }
        updateCount();
    });
}

// Intersection Observer for Stats
const observerOptions = {
    threshold: 0.7,
};

const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
        if(entry.isIntersecting && !animated) {
            animateStats();
            animated = true;
        }
    });
}, observerOptions);

if(statsSection) {
    observer.observe(statsSection);
}

// Smooth Scroll for Navigation Links
document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        e.preventDefault();
        document.querySelector(this.getAttribute('href')).scrollIntoView({
            behavior: 'smooth'
        });
    });
});

// Enhanced Mobile Menu Animation
const mobileMenu = document.querySelector('.mobile-menu');
mobileMenu.addEventListener('click', () => {
    mobileMenu.classList.toggle('active');
    document.querySelector('.nav-links').classList.toggle('active');
    
    // Animate hamburger to X
    const spans = mobileMenu.querySelectorAll('span');
    spans.forEach(span => span.classList.toggle('active'));
});

// Add loading animation
window.addEventListener('load', () => {
    document.body.classList.add('loaded');
});

// Add fade-in animations
const fadeElements = document.querySelectorAll('.fade-in');
const observerOptions = {
    threshold: 0.3,
    rootMargin: '0px 0px -50px 0px'
};

const fadeObserver = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
        if (entry.isIntersecting) {
            entry.target.classList.add('visible');
        }
    });
}, observerOptions);

fadeElements.forEach(element => {
    fadeObserver.observe(element);
});

// Hero Slider
const slides = document.querySelectorAll('.slide');
const dots = document.querySelector('.slide-dots');
const prevBtn = document.querySelector('.prev-slide');
const nextBtn = document.querySelector('.next-slide');
let currentSlide = 0;

// Create dots for each slide
slides.forEach((_, index) => {
    const dot = document.createElement('div');
    dot.classList.add('dot');
    if (index === 0) dot.classList.add('active');
    dot.addEventListener('click', () => goToSlide(index));
    dots.appendChild(dot);
});

// Function to update slides
function updateSlides() {
    slides.forEach((slide, index) => {
        if (index === currentSlide) {
            slide.style.transform = 'translateX(0)';
            slide.style.opacity = '1';
        } else if (index < currentSlide) {
            slide.style.transform = 'translateX(-100%)';
            slide.style.opacity = '0';
        } else {
            slide.style.transform = 'translateX(100%)';
            slide.style.opacity = '0';
        }
    });

    // Update dots
    document.querySelectorAll('.dot').forEach((dot, index) => {
        dot.classList.toggle('active', index === currentSlide);
    });
}

// Next slide function
function nextSlide() {
    currentSlide = (currentSlide + 1) % slides.length;
    updateSlides();
}

// Previous slide function
function prevSlide() {
    currentSlide = (currentSlide - 1 + slides.length) % slides.length;
    updateSlides();
}

// Go to specific slide
function goToSlide(index) {
    currentSlide = index;
    updateSlides();
}

// Event listeners for controls
prevBtn.addEventListener('click', () => {
    prevSlide();
    resetTimer();
});

nextBtn.addEventListener('click', () => {
    nextSlide();
    resetTimer();
});

// Auto slide functionality
let slideInterval = null;

function startTimer() {
    slideInterval = setInterval(nextSlide, 3000); // Change slide every 3 seconds
}

function resetTimer() {
    if (slideInterval) {
        clearInterval(slideInterval);
        startTimer();
    }
}

// Start auto-sliding
startTimer();

// Pause on hover
const heroSection = document.querySelector('.hero');
heroSection.addEventListener('mouseenter', () => {
    clearInterval(slideInterval);
    slideInterval = null;
});

heroSection.addEventListener('mouseleave', () => {
    startTimer();
});

// Initialize first slide
updateSlides(); 